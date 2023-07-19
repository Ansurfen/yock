// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util/container"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	_ yocki.Ycho = (*tlog)(nil)
	_ tea.Model  = (*tlog)(nil)
	_ io.Writer  = (*progressWriter)(nil)

	textInfo  = lipgloss.NewStyle().Foreground(lipgloss.Color("#006bbb")).Render("INFO")
	textWarn  = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff7b52")).Render("WARN")
	textError = lipgloss.NewStyle().Foreground(lipgloss.Color("#d251a6")).Render("ERROR")
	textFatal = lipgloss.NewStyle().Foreground(lipgloss.Color("#d251a6")).Render("FATAL")
	textDebug = lipgloss.NewStyle().Foreground(lipgloss.Color("#8866e9")).Render("DEBUG")
)

const (
	padding  = 2
	maxWidth = 80
)

type stackFrame struct {
	name string
	line int
}

func getTopCaller(skip int) stackFrame {
	pc, _, _, _ := runtime.Caller(skip)
	file, line := runtime.FuncForPC(pc).FileLine(pc)
	str := strings.Split(file, "/")
	name := str[len(str)-2] + "/" + str[len(str)-1]
	return stackFrame{
		name: name,
		line: line,
	}
}

type (
	logMsg      string
	progressMsg struct {
		id    int
		ratio float64
	}
)

type progressWriter struct {
	onProgress func(float64)
	total      int
	downloaded int
	progress   progress.Model
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	pw.downloaded += len(p)
	if pw.total > 0 && pw.onProgress != nil {
		pw.onProgress(float64(pw.downloaded) / float64(pw.total))
	}
	return len(p), nil
}

type tlog struct {
	prog *tea.Program
	pws  container.DoubleLinkedList[*progressWriter]
	loop bool
	opt  YchoOpt
}

func (t *tlog) Eventloop() {
	t.loop = true
	_, err := t.prog.Run()
	if err != nil {
		panic(err)
	}
	t.loop = false
}

func (t *tlog) Init() tea.Cmd {
	return nil
}

func (t *tlog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return t, tea.Quit
		}
	case tea.WindowSizeMsg:
		for pw := t.pws.Front(); pw != nil; pw = pw.Next() {
			if v := pw.Value(); v != nil {
				v.progress.Width = msg.Width - padding*2 - 4
				if v.progress.Width > maxWidth {
					v.progress.Width = maxWidth
				}
			}
		}
		return t, nil
	case progressMsg:
		if t.pws.Find(msg.id) == nil {
			break
		}
		if msg.ratio >= 1.0 {
			t.pws.Remove(t.pws.Find(msg.id))
		}
		return t, t.pws.Find(msg.id).Value().progress.SetPercent(float64(msg.ratio))
	case progress.FrameMsg:
		cmds := []tea.Cmd{}
		for pw := t.pws.Front(); pw != nil; pw = pw.Next() {
			if v := pw.Value(); v != nil {
				progressModel, cmd := v.progress.Update(msg)
				v.progress = progressModel.(progress.Model)
				cmds = append(cmds, cmd)
			}
		}
		return t, tea.Batch(cmds...)
	case logMsg:
		return t, tea.Printf(string(msg))
	}
	return t, nil
}

func (t *tlog) View() string {
	if t.pws.Size() == 0 {
		return ""
	}
	pad := strings.Repeat(" ", padding)
	s := ""
	for pw := t.pws.Front(); pw != nil; pw = pw.Next() {
		if v := pw.Value(); v != nil {
			s += "\n" + pad + pw.Value().progress.View() + pad + "\n"
		}
	}
	return s
}

func (t *tlog) Progress(total int64, r io.Reader) io.Writer {
	idx := t.pws.Size()
	pw := &progressWriter{
		total: int(total),
		onProgress: func(f float64) {
			t.prog.Send(progressMsg{ratio: f, id: idx})
		},
		progress: progress.New(progress.WithDefaultGradient()),
	}
	t.pws.PushBack(pw)
	return pw
}

func (t *tlog) Info(msg string) {
	t.logger(textInfo, msg)
}

func (t *tlog) Infof(msg string, a ...any) {
	t.logger(textInfo, msg, a...)
}

func (t *tlog) Debug(msg string) {
	t.logger(textDebug, msg)
}

func (t *tlog) Debugf(msg string, a ...any) {
	t.logger(textDebug, msg, a...)
}

func (t *tlog) Fatal(msg string) {
	t.logger(textFatal, msg)
}

func (t *tlog) Fatalf(msg string, a ...any) {
	t.logger(textFatal, msg, a...)
	os.Exit(1)
}

func (t *tlog) Warn(msg string) {
	t.logger(textWarn, msg)
}

func (t *tlog) Warnf(msg string, a ...any) {
	t.logger(textWarn, msg, a...)
}

func (t *tlog) Error(msg string) {
	t.logger(textError, msg)
}

func (t *tlog) Errorf(msg string, a ...any) {
	t.logger(textError, msg, a...)
}

func (t *tlog) Print(msg string) {
	if t.loop {
		t.prog.Send(logMsg(msg))
	} else {
		fmt.Println(msg)
	}
}

func (t *tlog) logger(level, msg string, a ...any) {
	if t.opt.Caller {
		fr := getTopCaller(4)
		msg = fmt.Sprintf("%s %s %s:%d %s",
			time.Now().Format(t.opt.TimeFormat),
			level, fr.name, fr.line, fmt.Sprintf(msg, a...))
	} else {
		msg = fmt.Sprintf("%s %s %s",
			time.Now().Format(t.opt.TimeFormat),
			level, fmt.Sprintf(msg, a...))
	}
	if t.loop {
		t.prog.Send(logMsg(msg))
	} else {
		fmt.Println(msg)
	}
}

func NewTLog(conf YchoOpt) (*tlog, error) {
	t := &tlog{
		pws:  container.VectorOf[*progressWriter](),
		loop: false,
		opt:  conf,
	}
	t.prog = tea.NewProgram(t)
	return t, nil
}
