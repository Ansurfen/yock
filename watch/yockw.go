package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())

	go p.Run()
	time.Sleep(100 * time.Second)
}

type errMsg error

type model struct {
	textarea textarea.Model
	err      error
	results  []result
	show     bool
	curPage  int
	pages    []Page
}

func initialModel() model {
	ti := textarea.New()
	ti.Placeholder = "Once upon a time..."
	ti.Focus()
	ti.SetHeight(1)
	ti.KeyMap.InsertNewline.SetEnabled(false)
	ti.ShowLineNumbers = false
	return model{
		textarea: ti,
		err:      nil,
		results:  make([]result, 5),
		show:     true,
		curPage:  0,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, runPretendProcess)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case processFinishedMsg:
		d := time.Duration(msg)
		res := result{emoji: randomEmoji(), duration: d}
		m.results = append(m.results, res)
		return m, runPretendProcess
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyLeft:
		case tea.KeyRight:
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			m.curPage = -1
			return m, tea.Quit
		case tea.KeyCtrlA:
			m.show = !m.show
		case tea.KeyEnter:
			if len(m.textarea.Value()) != 0 {
				m.results = append(m.results, result{emoji: m.textarea.Value(), duration: 10})
			}
			m.textarea.Reset()
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.curPage == -1 {
		return ""
	}
	s := ""
	for _, res := range m.results {
		if res.duration == 0 {
			s += "........................\n"
		} else {
			s += fmt.Sprintf("%s Job finished in %s\n", res.emoji, res.duration)
		}
	}
	if m.show {
		s += fmt.Sprintf(
			"\n\n%s\n\n%s",
			m.textarea.View(),
			"(ctrl+c to quit)",
		) + "\n\n"
	}
	return m.pages[m.curPage].View()
}

type result struct {
	duration time.Duration
	emoji    string
}

// processFinishedMsg is sent when a pretend process completes.
type processFinishedMsg time.Duration

// pretendProcess simulates a long-running process.
func runPretendProcess() tea.Msg {
	pause := time.Duration(rand.Int63n(899)+100) * time.Millisecond // nolint:gosec
	time.Sleep(pause)
	return processFinishedMsg(pause)
}

func randomEmoji() string {
	emojis := []rune("🍦🧋🍡🤠👾😭🦊🐯🦆🥨🎏🍔🍒🍥🎮📦🦁🐶🐸🍕🥐🧲🚒🥇🏆🌽")
	return string(emojis[rand.Intn(len(emojis))]) // nolint:gosec
}
