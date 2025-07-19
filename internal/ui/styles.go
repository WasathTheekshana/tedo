package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Color palette
	primaryColor   = lipgloss.Color("86")  // Cyan
	secondaryColor = lipgloss.Color("212") // Pink
	accentColor    = lipgloss.Color("57")  // Blue
	successColor   = lipgloss.Color("42")  // Green
	warningColor   = lipgloss.Color("214") // Orange
	errorColor     = lipgloss.Color("196") // Red
	mutedColor     = lipgloss.Color("243") // Gray

	// Base styles
	baseStyle = lipgloss.NewStyle().
			Padding(1, 2)

	// Header styles
	headerStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(0, 1)

	activeTabStyle = lipgloss.NewStyle().
			Background(primaryColor).
			Foreground(lipgloss.Color("0")).
			Bold(true).
			Padding(0, 1)

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				Padding(0, 1)

	// List styles
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true)

	completedItemStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				Strikethrough(true)

	normalItemStyle = lipgloss.NewStyle()

	// Accent style for dates with todos
	accentStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)

	// Footer styles
	footerStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Padding(1, 1)

	// Error styles
	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	// Muted style for help text
	mutedStyle = lipgloss.NewStyle().
			Foreground(mutedColor)
)

// getViewName returns the display name for a view type
func getViewName(view ViewType) string {
	switch view {
	case TodayView:
		return "Today"
	case UpcomingView:
		return "Upcoming"
	case CalendarView:
		return "Calendar"
	case GeneralView:
		return "General"
	default:
		return "Unknown"
	}
}
