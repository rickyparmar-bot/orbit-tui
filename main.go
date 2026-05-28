package main

// imports
import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	Name                    string
	Primary, Accent, Warn   lipgloss.Color
	Base, Subtle, Bg, TabBg lipgloss.Color
}

var themes = []Theme{
	// Catppuccin Theme, who doesn't love catppuccin. hehehe
	{"Purple", "#CBA6F7", "#A6E3A1", "#F38BA8", "#CDD6F4", "#6C7086", "#1E1E2E", "#313244"},
	{"Green", "#A6E3A1", "#CBA6F7", "#F38BA8", "#CDD6F4", "#6C7086", "#1E1E2E", "#313244"},
	{"Blue", "#89B4FA", "#A6E3A1", "#F38BA8", "#CDD6F4", "#6C7086", "#1E1E2E", "#313244"},
}

type ProcessEntry struct {
	PID, Name, CPU, Memory, Status string
}

type PortEntry struct {
	Port, PID, Name string
}

type ConnectionEntry struct {
	Local, Remote, State, Process, Protocol string
}

type Model struct {
	ScreenWidth, ScreenHeight int
	ActiveTab                 int
	TabNames                  []string
	Spinner                   spinner.Model
	CPUProgress               progress.Model
	MemProgress               progress.Model
	DiskProgress              progress.Model
	ThemeIndex                int
	StatusMsg                 string
	ShowingHelp               bool

	CPUUsage     float64
	MemUsage     float64
	DiskUsage    float64
	MemTotal     string
	MemUsed      string
	MemAvail     string
	SwapTotal    string
	SwapUsed     string
	DiskTotal    string
	DiskUsed     string
	DiskRead     string
	DiskWrite    string
	DiskReadSpd  string
	DiskWriteSpd string
	SystemUptime string
	NetSent      string
	NetRecv      string
	NetUpSpd     string
	NetDnSpd     string
	BattPercent  int
	BattState    string
	BattTime     string
	GPUUsage     float64
	GPUName      string
	GPUMemUsed   string
	GPUMemTotal  string
	PublicIP     string
	CPUTemp      float64
	GPUTemp      float64
	CPUWarn      bool
	MemWarn      bool
	DiskWarn     bool

	PrevDiskReadBytes  uint64
	PrevDiskWriteBytes uint64
	PrevNetSentBytes   uint64
	PrevNetRecvBytes   uint64
	PrevSampleTime     int64

	ProcessList []ProcessEntry
	PortList    []PortEntry
	ConnList    []ConnectionEntry

	ProcTable table.Model
	PortTable table.Model
	ConnTable table.Model
}

func headerStyle(th Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(th.Base).
		Background(th.Primary).
		Bold(true)
}

func cardStyle(th Theme, w int) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.Primary).
		Padding(1, 2).
		Width(w - 4)
}

func tabStyle(th Theme, active bool) lipgloss.Style {
	if active {
		return lipgloss.NewStyle().
			Foreground(th.Primary).
			Background(th.TabBg).
			Padding(0, 2).
			Bold(true).
			BorderBottom(true).
			BorderForeground(th.Primary)
	}
	return lipgloss.NewStyle().
		Foreground(th.Subtle).
		Background(th.TabBg).
		Padding(0, 2)
}

func labelStyle(th Theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Accent).Bold(true)
}

func mutedStyle(th Theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Subtle)
}

func warnStyle(th Theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Warn).Bold(true)
}

func valueStyle(th Theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Base)
}

func sectionStyle(th Theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Base).Bold(true)
}

func warnLabelStyle(warn bool, label string, th Theme) string {
	if warn {
		return warnStyle(th).Render(label)
	}
	return labelStyle(th).Render(label)
}

func main() {
	p := tea.NewProgram(createModel(themes[0]), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
q