package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// handleTodayViewKeys handles keys specific to today view
func (m Model) handleTodayViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Debug: let's see what keys we're getting
	switch msg.String() {
	case "j", "down":
		if len(m.todayTodos) > 0 && m.cursor < len(m.todayTodos)-1 {
			m.cursor++
		}
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "x":
		return m.toggleCurrentTodo(), nil
	}
	return m, nil
}

// handleCalendarViewKeys handles keys specific to calendar view
func (m Model) handleCalendarViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Placeholder - will implement in next task
	return m, nil
}

// handleGeneralViewKeys handles keys specific to general view
func (m Model) handleGeneralViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "j", "down":
		if len(m.generalTodos) > 0 && m.cursor < len(m.generalTodos)-1 {
			m.cursor++
		}
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "x":
		return m.toggleCurrentGeneralTodo(), nil
	}
	return m, nil
}

// toggleCurrentTodo toggles completion of current today todo
func (m Model) toggleCurrentTodo() Model {
	if len(m.todayTodos) > 0 && m.cursor < len(m.todayTodos) {
		m.todayTodos[m.cursor].Toggle()
		if err := m.repository.UpdateTodo(m.todayTodos[m.cursor]); err != nil {
			m.err = err
		}
	}
	return m
}

// toggleCurrentGeneralTodo toggles completion of current general todo
func (m Model) toggleCurrentGeneralTodo() Model {
	if len(m.generalTodos) > 0 && m.cursor < len(m.generalTodos) {
		m.generalTodos[m.cursor].Toggle()
		if err := m.repository.UpdateTodo(m.generalTodos[m.cursor]); err != nil {
			m.err = err
		}
	}
	return m
}
