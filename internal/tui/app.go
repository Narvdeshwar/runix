package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Tab   key.Binding
	Enter key.Binding
	Quit  key.Binding
}

var keys = keyMap{
	Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
	Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
	Tab:   key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "switch tab")),
	Enter: key.NewBinding(key.WithKeys("enter", " "), key.WithHelp("enter", "launch")),
	Quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q/ctrl+c", "quit")),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Tab, k.Enter, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

type activePane int

const (
	paneWorkspaces activePane = iota
	paneNetwork
	paneSettings
)

type model struct {
	cursor     int
	workspaces []string
	activeTab  int
	tabs       []string
	quitting   bool
	width      int
	height     int
	help       help.Model
}

func InitialModel() model {
	return model{
		workspaces: []string{"Ubuntu-Dev-01 (Isolated)", "Python-Sandbox (Isolated)", "NodeJS-AI-Env (Isolated)", "Frontend-React (Isolated)"},
		tabs:       []string{"Workspaces", "Network", "Settings"},
		help:       help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, keys.Up):
			if m.activeTab == 0 && m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, keys.Down):
			if m.activeTab == 0 && m.cursor < len(m.workspaces)-1 {
				m.cursor++
			}

		case key.Matches(msg, keys.Tab):
			m.activeTab = (m.activeTab + 1) % len(m.tabs)

		case key.Matches(msg, keys.Enter):
			if m.activeTab == 0 {
				// Launch workspace
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "\n  Exiting Command Center...\n\n"
	}

	// 1. Header
	header := HeaderStyle.Render(" ✨ MYTERM STUDIO ")

	// 2. Tabs
	var renderedTabs []string
	for i, t := range m.tabs {
		if m.activeTab == i {
			renderedTabs = append(renderedTabs, ActiveTabStyle.Render(t))
		} else {
			renderedTabs = append(renderedTabs, InactiveTabStyle.Render(t))
		}
	}
	tabsRow := TabsStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...))

	// 3. Main Content Area
	var content string

	switch m.activeTab {
	case 0: // Workspaces
		// Left Pane: Workspaces List
		var listBuilder strings.Builder
		listBuilder.WriteString(lipgloss.NewStyle().Foreground(Secondary).Bold(true).Render("Available Environments") + "\n\n")

		for i, w := range m.workspaces {
			val := w
			if m.cursor == i {
				val = ListSelectedStyle.Render(fmt.Sprintf("▶ %s", w))
			} else {
				val = ListItemStyle.Render(w)
			}
			listBuilder.WriteString(val + "\n")
		}

		// Fill remaining lines
		remainingLines := 10 - len(m.workspaces)
		for i := 0; i < remainingLines; i++ {
			listBuilder.WriteString("\n")
		}

		leftPane := PaneStyle.Width(40).Height(15).Render(listBuilder.String())

		// Right Pane: Details & Metrics
		var rightBuilder strings.Builder
		rightBuilder.WriteString(MetricsTitleStyle.Render("Workspace Details") + "\n\n")

		selectedWorkspace := m.workspaces[m.cursor]
		rightBuilder.WriteString(fmt.Sprintf("%s %s\n", MetricsLabelStyle.Render("Name:"), lipgloss.NewStyle().Bold(true).Render(selectedWorkspace)))
		rightBuilder.WriteString(fmt.Sprintf("%s %s\n", MetricsLabelStyle.Render("Status:"), MetricsValueStyle.Render("Ready")))
		rightBuilder.WriteString(fmt.Sprintf("%s %s\n\n", MetricsLabelStyle.Render("Uptime:"), lipgloss.NewStyle().Render("0h 0m")))

		rightBuilder.WriteString(MetricsTitleStyle.Render("System Health") + "\n\n")
		rightBuilder.WriteString(fmt.Sprintf("%s %s\n", MetricsLabelStyle.Render("CPU Usage:"), lipgloss.NewStyle().Foreground(Green).Render("12%")))
		rightBuilder.WriteString(fmt.Sprintf("%s %s\n", MetricsLabelStyle.Render("Mem Usage:"), lipgloss.NewStyle().Foreground(Accent).Render("2.4GB / 16GB")))
		rightBuilder.WriteString(fmt.Sprintf("%s %s\n", MetricsLabelStyle.Render("Net I/O:  "), lipgloss.NewStyle().Render("15KB/s ↑ 45KB/s ↓")))

		for i := 0; i < 2; i++ {
			rightBuilder.WriteString("\n")
		}

		rightPane := InactivePaneStyle.Width(40).Height(15).Render(rightBuilder.String())

		// Combine left and right panes
		content = lipgloss.JoinHorizontal(lipgloss.Top, leftPane, lipgloss.NewStyle().Width(2).Render(""), rightPane)

	case 1: // Network
		content = PaneStyle.Width(86).Height(15).Render(MetricsTitleStyle.Render("Network Activity") + "\n\nNo active network connections isolated.")
	case 2: // Settings
		content = PaneStyle.Width(86).Height(15).Render(MetricsTitleStyle.Render("Settings") + "\n\nAdjust global preferences here.")
	}

	// 4. Status Bar & Help
	statusBar := StatusBarStyle.Width(86).Render(" STATUS: SYSTEM NOMINAL ")
	helpMenu := HelpStyle.Render(m.help.View(keys))

	// Assemble All
	view := lipgloss.JoinVertical(lipgloss.Left,
		header,
		tabsRow,
		content,
		statusBar,
		helpMenu,
	)

	return AppStyle.Render(view)
}
