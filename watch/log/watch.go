// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ansurfen/yock/util"
)

var DefaultLoggerWatch *LoggerWatch

const LoggerFormat = "(.*)(\033\\[\\d+m)+(INFO|DEBUG|FATAL|WARN|PANIC|ERROR)(\033\\[0m)+(.*)"

type LoggerWatch struct {
	loggerEntries map[string][]*loggerEntry
}

func New() *LoggerWatch {
	return &LoggerWatch{
		loggerEntries: make(map[string][]*loggerEntry),
	}
}

func (lw *LoggerWatch) Find(file, time, level, caller, msg string) (ret []*loggerEntry) {
	for name, entries := range lw.loggerEntries {
		if file != "*" && !strings.Contains(name, file) {
			continue
		}
		for _, entry := range entries {
			if time != "*" && !strings.Contains(entry.Time, time) {
				continue
			}
			if level != "*" && !strings.Contains(entry.Level, level) {
				continue
			}
			if caller != "*" && !strings.Contains(entry.Caller, caller) {
				continue
			}
			if msg != "*" && !strings.Contains(entry.Msg, msg) {
				continue
			}
			if !entry.isTrim {
				entry.isTrim = true
				entry.trim()
			}
			ret = append(ret, entry)
		}
	}
	return
}

func (lw *LoggerWatch) Parse(path string) error {
	re := regexp.MustCompile(LoggerFormat)
	fmt.Println(path)
	if filepath.Ext(path) == ".log" {
		lw.loggerEntries[path] = []*loggerEntry{}
		util.ReadLineFromFile(path, func(s string) string {
			s = strings.TrimSpace(s)
			if len(s) == 0 {
				return ""
			}
			if re.MatchString(s) {
				res := re.FindStringSubmatch(s)
				lw.loggerEntries[path] = append(lw.loggerEntries[path], &loggerEntry{
					Time:  strings.TrimSpace(res[1]),
					Level: strings.TrimSpace(res[3]),
					Msg:   strings.TrimLeft(res[5], " "),
				})
			} else {
				n := len(lw.loggerEntries[path]) - 1
				lw.loggerEntries[path][n].Msg += "\n" + s
			}
			return ""
		})
	}
	return nil
}

type loggerEntry struct {
	Time   string `json:"time"`
	Level  string `json:"level"`
	Caller string `json:"caller"`
	Msg    string `json:"msg"`
	isTrim bool   `json:"-"`
}

var extract = regexp.MustCompile(`(.*:\d+)`)

func (entry *loggerEntry) trim() {
	if extract.MatchString(entry.Msg) {
		if loc := extract.FindStringIndex(entry.Msg); len(loc) > 1 {
			entry.Caller = strings.TrimSpace(entry.Msg[:loc[1]])
			entry.Msg = strings.TrimSpace(entry.Msg[loc[1]:])
		}
	} else {
		entry.Msg = strings.TrimSpace(entry.Msg)
	}
}
