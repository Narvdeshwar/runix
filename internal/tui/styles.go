package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	Primary   = lipgloss.AdaptiveColor{Light: "#7D56F4", Dark: "#7D56F4"}
	Secondary = lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"}
	Accent    = lipgloss.AdaptiveColor{Light: "#00f2ff", Dark: "#00f2ff"}
	Green     = lipgloss.AdaptiveColor{Light: "#22c55e", Dark: "#22c55e"}
	Gray      = lipgloss.AdaptiveColor{Light: "#626262", Dark: "#626262"}
	LightGray = lipgloss.AdaptiveColor{Light: "#c1c1c1", Dark: "#c1c1c1"}
	White     = lipgloss.Color("#ffffff")
	Black     = lipgloss.Color("#000000")

	// Global Layout Styles
	AppStyle = lipgloss.NewStyle().Padding(1, 2)

	// Header Styles
	HeaderStyle = lipgloss.NewStyle().
			Foreground(Black).
			Background(Primary).
			Bold(true).
			Padding(0, 2).
			MarginBottom(1)

	// Tab Styles
	ActiveTabStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true).
			BorderForeground(Primary).
			Padding(0, 1).
			Foreground(Primary).
			Bold(true)

	InactiveTabStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true).
			BorderForeground(Gray).
			Padding(0, 1).
			Foreground(LightGray)

	TabsStyle = lipgloss.NewStyle().MarginBottom(1)

	// Layout panes
	PaneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Primary).
			Padding(1, 2)

	InactivePaneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Gray).
			Padding(1, 2)

	// List Styles
	ListItemStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(LightGray)

	ListSelectedStyle = lipgloss.NewStyle().
			PaddingLeft(0).
			Foreground(Accent).
			Bold(true)

	// Metrics Styles
	MetricsTitleStyle = lipgloss.NewStyle().
			Foreground(Secondary).
			Bold(true).
			MarginBottom(1)

	MetricsLabelStyle = lipgloss.NewStyle().
			Foreground(LightGray)

	MetricsValueStyle = lipgloss.NewStyle().
			Foreground(Green).
			Bold(true)

	// Footer Styles
	StatusBarStyle = lipgloss.NewStyle().
			Foreground(White).
			Background(Secondary).
			Padding(0, 1).
			MarginTop(1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(Gray).
			MarginTop(1)
)
