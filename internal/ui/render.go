package ui

import (
	"fmt"
	"strings"
)

// renderHeader renders the top navigation bar
func (m Model) renderHeader() string {
	var tabs []string

	// Update to include all four views
	views := []ViewType{TodayView, UpcomingView, CalendarView, GeneralView}

	for _, view := range views {
		name := getViewName(view)
		if view == m.currentView {
			tabs = append(tabs, activeTabStyle.Render(name))
		} else {
			tabs = append(tabs, inactiveTabStyle.Render(name))
		}
	}

	header := headerStyle.Render("📋 Todo CLI") + " " + strings.Join(tabs, " ")
	return header + "\n"
}

// renderFooter renders the bottom help bar
func (m Model) renderFooter() string {
	// Different help text based on input mode
	if m.inputState.mode != NavigationMode {
		help := []string{
			"tab: switch field",
			"enter: save",
			"esc: cancel",
		}
		return footerStyle.Render(strings.Join(help, " • "))
	}

	// Different help for calendar view
	if m.currentView == CalendarView {
		help := []string{
			"h/j/k/l: navigate dates",
			"n/p: month",
			"t: today",
			"enter: view date",
			"i: add",
			"←/→: switch tabs",
			"q: quit",
		}
		return footerStyle.Render(strings.Join(help, " • "))
	}

	// Help for Today, Upcoming, and General views
	help := []string{
		"j/k: navigate",
		"←/→: switch tabs",
		"x: toggle",
		"d: delete",
		"e: edit",
		"i: add",
		"c: calendar",
		"q: quit",
	}

	return footerStyle.Render(strings.Join(help, " • "))
}

// renderTodayView renders the today's todos view
func (m Model) renderTodayView() string {
	// If in input mode, show the input form
	if m.inputState.mode != NavigationMode {
		return m.renderInputForm()
	}

	paginatedTodos, currentPage, totalPages := m.getPaginatedTodos()

	if len(m.todayTodos) == 0 {
		return baseStyle.Render(
			fmt.Sprintf("📅 %s\n\nNo todos for today!\n\nPress 'i' to add a new todo.", m.selectedDate),
		)
	}

	var items []string

	// Header with pagination info
	header := fmt.Sprintf("📅 %s", m.selectedDate)
	if totalPages > 1 {
		header += fmt.Sprintf(" (Page %d/%d - %d total)", currentPage+1, totalPages, len(m.todayTodos))
	} else {
		header += fmt.Sprintf(" (%d todos)", len(m.todayTodos))
	}
	items = append(items, header+"\n")

	for i, todo := range paginatedTodos {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		checkbox := "☐"
		style := normalItemStyle
		if todo.Completed {
			checkbox = "✓"
			style = completedItemStyle
		}

		if i == m.cursor {
			style = selectedItemStyle
		}

		// Show absolute index
		absoluteIndex := currentPage*TodosPerPage + i + 1
		line := fmt.Sprintf("%s %s %d. %s", cursor, checkbox, absoluteIndex, todo.Title)
		if todo.Description != "" {
			line += fmt.Sprintf("\n      %s", todo.Description)
		}

		items = append(items, style.Render(line))
	}

	// Add pagination help if needed
	if totalPages > 1 {
		items = append(items, "")
		items = append(items, mutedStyle.Render("Navigation: j/k=item, Ctrl+f/b=page"))
	}

	return baseStyle.Render(strings.Join(items, "\n"))
}

func (m Model) renderUpcomingView() string {
	// If in input mode, show the input form
	if m.inputState.mode != NavigationMode {
		return m.renderInputForm()
	}

	paginatedTodos, currentPage, totalPages := m.getPaginatedTodos()

	if len(m.upcomingTodos) == 0 {
		return baseStyle.Render("📅 Upcoming Todos\n\nNo upcoming todos!\n\nPress 'i' to add a new todo or 'c' for calendar.")
	}

	var items []string

	// Header with pagination info
	header := "📅 Upcoming Todos"
	if totalPages > 1 {
		header += fmt.Sprintf(" (Page %d/%d - %d total)", currentPage+1, totalPages, len(m.upcomingTodos))
	} else {
		header += fmt.Sprintf(" (%d todos)", len(m.upcomingTodos))
	}
	items = append(items, header+"\n")

	for i, todo := range paginatedTodos {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		checkbox := "☐"
		style := normalItemStyle
		if todo.Completed {
			checkbox = "✓"
			style = completedItemStyle
		}

		if i == m.cursor {
			style = selectedItemStyle
		}

		// Show date and absolute index
		absoluteIndex := currentPage*TodosPerPage + i + 1
		dateStr := ""
		if todo.Date != nil {
			dateStr = fmt.Sprintf(" (%s)", *todo.Date)
		}
		line := fmt.Sprintf("%s %s %d. %s%s", cursor, checkbox, absoluteIndex, todo.Title, dateStr)
		if todo.Description != "" {
			line += fmt.Sprintf("\n      %s", todo.Description)
		}

		items = append(items, style.Render(line))
	}

	// Add pagination help if needed
	if totalPages > 1 {
		items = append(items, "")
		items = append(items, mutedStyle.Render("Navigation: j/k=item, Ctrl+f/b=page"))
	}

	return baseStyle.Render(strings.Join(items, "\n"))
}

// renderCalendarView renders the calendar view (placeholder for now)
func (m Model) renderCalendarView() string {
	// If in input mode, show the input form
	if m.inputState.mode != NavigationMode {
		return m.renderInputForm()
	}

	calendar := m.renderCalendar()

	// Add help text
	help := []string{
		"",
		mutedStyle.Render("Navigation: h/j/k/l=move, n/p=month, t=today, enter=view date, i=add todo"),
	}

	return baseStyle.Render(calendar + strings.Join(help, "\n"))
}

// renderGeneralView renders the general todos view
func (m Model) renderGeneralView() string {
	// If in input mode, show the input form
	if m.inputState.mode != NavigationMode {
		return m.renderInputForm()
	}

	paginatedTodos, currentPage, totalPages := m.getPaginatedTodos()

	if len(m.generalTodos) == 0 {
		return baseStyle.Render("📝 General Todos\n\nNo general todos!\n\nPress 'i' to add a new todo.")
	}

	var items []string

	// Header with pagination info
	header := "📝 General Todos"
	if totalPages > 1 {
		header += fmt.Sprintf(" (Page %d/%d - %d total)", currentPage+1, totalPages, len(m.generalTodos))
	} else {
		header += fmt.Sprintf(" (%d todos)", len(m.generalTodos))
	}
	items = append(items, header+"\n")

	for i, todo := range paginatedTodos {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		checkbox := "☐"
		style := normalItemStyle
		if todo.Completed {
			checkbox = "✓"
			style = completedItemStyle
		}

		if i == m.cursor {
			style = selectedItemStyle
		}

		// Show absolute index
		absoluteIndex := currentPage*TodosPerPage + i + 1
		line := fmt.Sprintf("%s %s %d. %s", cursor, checkbox, absoluteIndex, todo.Title)
		if todo.Description != "" {
			line += fmt.Sprintf("\n      %s", todo.Description)
		}

		items = append(items, style.Render(line))
	}

	// Add pagination help if needed
	if totalPages > 1 {
		items = append(items, "")
		items = append(items, mutedStyle.Render("Navigation: j/k=item, Ctrl+f/b=page"))
	}

	return baseStyle.Render(strings.Join(items, "\n"))
}

// renderInputForm renders the input form for adding/editing todos
func (m Model) renderInputForm() string {
	var title string
	if m.inputState.mode == AddTodoMode {
		title = "➕ Add New Todo"
	} else {
		title = "✏️  Edit Todo"
	}

	// Render title field
	titleLabel := "Title:"
	titleValue := m.inputState.title

	if m.inputState.editField == 0 {
		// Show cursor in title field
		if m.inputState.cursor <= len(titleValue) {
			titleValue = titleValue[:m.inputState.cursor] + "│" + titleValue[m.inputState.cursor:]
		}
		titleLabel = selectedItemStyle.Render(titleLabel)
	} else {
		titleLabel = normalItemStyle.Render(titleLabel)
	}

	// Render description field
	descLabel := "Description:"
	descValue := m.inputState.description

	if m.inputState.editField == 1 {
		// Show cursor in description field
		if m.inputState.cursor <= len(descValue) {
			descValue = descValue[:m.inputState.cursor] + "│" + descValue[m.inputState.cursor:]
		}
		descLabel = selectedItemStyle.Render(descLabel)
	} else {
		descLabel = normalItemStyle.Render(descLabel)
	}

	// Build the form
	form := []string{
		title,
		"",
		titleLabel,
		"  " + titleValue,
		"",
		descLabel,
		"  " + descValue,
		"",
		mutedStyle.Render("Press Tab to switch fields, Enter to save, Esc to cancel"),
	}

	return baseStyle.Render(strings.Join(form, "\n"))
}
