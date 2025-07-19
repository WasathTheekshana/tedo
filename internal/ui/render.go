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
	help := []string{
		"j/k: navigate",
		"h/l: switch tabs",
		"x: toggle",
		"d: delete",
		"i: add",
		"q: quit",
	}

	return footerStyle.Render(strings.Join(help, " • "))
}

// renderTodayView renders the today's todos view
func (m Model) renderTodayView() string {
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
