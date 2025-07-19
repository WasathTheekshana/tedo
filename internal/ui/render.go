package ui

import (
	"fmt"
	"strings"
)

// renderHeader renders the top navigation bar
func (m Model) renderHeader() string {
	var tabs []string

	views := []ViewType{TodayView, CalendarView, GeneralView}

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

	help := []string{
		"j/k: navigate",
		"h/l: switch tabs",
		"x: toggle",
		"d: delete",
		"e: edit",
		"i: add",
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

	if len(m.todayTodos) == 0 {
		return baseStyle.Render(
			fmt.Sprintf("📅 %s\n\nNo todos for today!\n\nPress 'i' to add a new todo.", m.selectedDate),
		)
	}

	var items []string
	items = append(items, fmt.Sprintf("📅 %s\n", m.selectedDate))

	for i, todo := range m.todayTodos {
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

		line := fmt.Sprintf("%s %s %s", cursor, checkbox, todo.Title)
		if todo.Description != "" {
			line += fmt.Sprintf("\n    %s", todo.Description)
		}

		items = append(items, style.Render(line))
	}

	return baseStyle.Render(strings.Join(items, "\n"))
}

// renderCalendarView renders the calendar view (placeholder for now)
func (m Model) renderCalendarView() string {
	return baseStyle.Render("📅 Calendar View\n\nComing soon! Use arrow keys to navigate dates.")
}

// renderGeneralView renders the general todos view
func (m Model) renderGeneralView() string {
	// If in input mode, show the input form
	if m.inputState.mode != NavigationMode {
		return m.renderInputForm()
	}

	if len(m.generalTodos) == 0 {
		return baseStyle.Render("📝 General Todos\n\nNo general todos!\n\nPress 'i' to add a new todo.")
	}

	var items []string
	items = append(items, "📝 General Todos\n")

	for i, todo := range m.generalTodos {
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

		line := fmt.Sprintf("%s %s %s", cursor, checkbox, todo.Title)
		if todo.Description != "" {
			line += fmt.Sprintf("\n    %s", todo.Description)
		}

		items = append(items, style.Render(line))
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
