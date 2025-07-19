package ui

// HelpContent provides detailed help information
type HelpContent struct {
	view    ViewType
	content string
}

// GetHelpContent returns help content for the current view
func (m Model) GetHelpContent() string {
	switch m.currentView {
	case TodayView:
		return `Today View Help:
- j/k: Navigate up/down in todo list
- ←/→: Switch between tabs
- x: Toggle todo completion
- i: Add new todo for today
- e: Edit selected todo
- d: Delete selected todo
- c: Jump to calendar view
- Ctrl+F/B: Next/previous page (10+ todos)
- q: Quit application`

	case UpcomingView:
		return `Upcoming View Help:
- j/k: Navigate up/down in todo list
- ←/→: Switch between tabs  
- x: Toggle todo completion
- i: Add new todo for selected date
- e: Edit selected todo
- d: Delete selected todo
- c: Jump to calendar view
- Ctrl+F/B: Next/previous page (10+ todos)
- q: Quit application`

	case CalendarView:
		return `Calendar View Help:
- h/j/k/l: Navigate calendar dates
- ←/→: Switch between tabs
- n/p: Next/previous month
- t: Jump to today's date
- Enter: View todos for selected date
- i: Add todo for selected date
- >/<: Next/previous month (alternative)
- q: Quit application`

	case GeneralView:
		return `General View Help:
- j/k: Navigate up/down in todo list
- ←/→: Switch between tabs
- x: Toggle todo completion
- i: Add new general todo
- e: Edit selected todo  
- d: Delete selected todo
- c: Jump to calendar view
- Ctrl+F/B: Next/previous page (10+ todos)
- q: Quit application`

	default:
		return "No help available for this view."
	}
}

// GetInputHelp returns help for input mode
func GetInputHelp() string {
	return `Input Mode Help:
- Tab: Switch between title and description
- Enter/Ctrl+S: Save todo
- Esc: Cancel and return to list
- Ctrl+A: Select all text in current field
- Ctrl+C: Quit application

Validation Rules:
- Title: Required, max 100 characters
- Description: Optional, max 500 characters
- Only printable characters allowed`
}
