// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	yocki "github.com/ansurfen/yock/interface"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func LoadTea(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("tea")
	lib.SetField(map[string]any{
		"NewProgram": tea.NewProgram,
		"NewModel":   newTeaModel,
		"Quit":       tea.Quit,
		"KeyMsg":     tea.KeyMsg{},
		"NewStyle":   NewStyle,
		"NewSpinner": spinner.New,
	})
}

type Style struct {
	lipgloss.Style
}

func NewStyle() *Style {
	return &Style{Style: lipgloss.NewStyle()}
}

func (s *Style) Foreground(c string) *Style {
	s.Style = s.Style.Foreground(lipgloss.Color(c))
	return s
}

var _ tea.Model = (*TeaModel)(nil)

func newTeaModel() *TeaModel {
	return &TeaModel{}
}

type TeaModel struct {
	InitCallback   func() tea.Cmd
	UpdateCallback func(tea.Msg) (tea.Model, tea.Cmd)
	ViewCallback   func() string
}

func (m *TeaModel) Init() tea.Cmd {
	return m.InitCallback()
}

func (m *TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.UpdateCallback(msg)
}

func (m *TeaModel) View() string {
	return m.ViewCallback()
}
