package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
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
				if r < len(m.ProcTable.Rows())-1 {
					m.ProcTable.SetCursor(r + 1)
				}
			} else if m.ActiveTab == 2 {
				r := m.ProcTable.Cursor()
				if r < len(m.PortTable.Rows())-1 {
					m.PortTable.SetCursor(r + 1)
				}
			}
		case "enter":
			if m.ActiveTab == 1 {
				row := m.ProcTable.SelectedRow()
				if row != nil && len(row) > 0 {
					pid, err := strconv.Atoi(row[0])
					if err == nil && pid > 0 {
						if err := killProcess(int32(pid)); err != nil {
							m.StatusMsg = fmt.Sprintf("Failed to kill %d: %v", pid, err)
						} else {
							m.StatusMsg = fmt.Sprintf("Killed PID %d", pid)
						}
					}
				}
			} else if m,ActiveTab == 2 {
				row := m.PortTable.SelectedRow()
				if row != nil && len(row) > 1 {
					pid, err := strconv.Atoi(row[1])
					if err == nil && pid > 0 {
						if err := killProcess(int32(pid)); err != nil {
							m.StatusMsg = fmt.Sprintf("Failed to kill %d: %v", pid, err)
						} else {
							m.StatusMsg = fmt.Sprintf("Killed PID %d", pid)
						}
					}
				}
			}
		case "shift+right":
			m.ThemeIndex = (m.ThemeIndex + 1) % len(themes)
			th = themes[m.ThemeIndex]
			applyTableStyles(m, th)
			m.StatusMsg = fmt.Sprintf("Theme: %s", th.Name)
		case "shift+left":
			m.ThemeIndex = (m.ThemeIndex - 1 + len(themes)) % len(themes)
			th = themes[m.ThemeIndex]
			applyTableStyles(m, th)
			m.StatusMsg = fmt.Sprintf("Theme: %s", th.Name)
		}

	case tea.WindowSizeMsg:
		m.ScreenWidth = msg.Width
		m.ScreenHeight = msg.Height
		m.CPUProgress.Width = (m.ScreenWidth /3) - 10
		m.MemProgress.Width = (m.ScreenWidth /3) - 10
		m.DiskProgress.Width = (m.ScreenWidth / 3) - 10

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		cmds = append(cmds, cmd)

	case tickMsg:
		m.collectSystemStats()
		cmds = append(cmds, tickCommand())
	}

	if m.ActiveTab == 1 {
		var cmd tea.Cmd
		m.ProcTable, cmd = m.ProcTable.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.ActiveTab == 2 {
		var cmd tea.Cmd
		m.PortTable, cmd = m.PortTable.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func applyTableStyles(m *Model, th Theme) {
	for _, tbl := range []*table.Model{&m.ProcTable, &m,&m.PortTable, &m.ConnTable} {
		
	}
    	
}
