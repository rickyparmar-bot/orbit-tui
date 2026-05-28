package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// imports

type tickMsg time.Time

func (m *Model) Init() tea.Cmd {
	return tea.Batch(m.Spinner.Tick, tickCommand())
}

func tickCommand() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	th := themes[m.ThemeIndex]

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "?":
			m.ShowingHelp = !m.ShowingHelp
		case "right", "tab":
			m.ActiveTab = (m.ActiveTab + 1) % len(m.TabNames)
			m.StatusMsg = ""
		case "left", "shift+tab":
			m.ActiveTab = (m.ActiveTab - 1 + len(m.TabNames)) % len(m.TabNames)
			m.StatusMsg = ""
		case "up":
			if m.ActiveTab == 1 {
				r := m.ProcTable.Cursor()
				if r > 0 {
					m.ProcTable.SetCursor(r - 1)
				}
			} else if m.ActiveTab == 2 {
				r := m.PortTable.Cursor()
				if r > 0 {
					m.PortTable.SetCursor(r - 1)
				}
			}
		case "down":
			if m.ActiveTab == 1 {
				r := m.ProcTable.Cursor()
			}
		}
	}
}
