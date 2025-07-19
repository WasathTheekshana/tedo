package ui

import tea "github.com/charmbracelet/bubbletea"

// handleTodayViewKeys handles keys specific to today view
func (m Model) handleTodayViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
	case "i":
		m.inputState.StartAddMode()
		return m, nil
	case "e":
		return m.editCurrentTodo()
	case "d":
		return m.deleteCurrentTodo()
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
	case "i":
		m.inputState.StartAddMode()
		return m, nil
	case "e":
		return m.editCurrentGeneralTodo()
	case "d":
		return m.deleteCurrentGeneralTodo()
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

// editCurrentTodo starts editing the current today todo
func (m Model) editCurrentTodo() (tea.Model, tea.Cmd) {
	if len(m.todayTodos) > 0 && m.cursor < len(m.todayTodos) {
		m.inputState.StartEditMode(&m.todayTodos[m.cursor])
	}
	return m, nil
}

// editCurrentGeneralTodo starts editing the current general todo
func (m Model) editCurrentGeneralTodo() (tea.Model, tea.Cmd) {
	if len(m.generalTodos) > 0 && m.cursor < len(m.generalTodos) {
		m.inputState.StartEditMode(&m.generalTodos[m.cursor])
	}
	return m, nil
}

// deleteCurrentTodo deletes the current today todo
func (m Model) deleteCurrentTodo() (tea.Model, tea.Cmd) {
	if len(m.todayTodos) > 0 && m.cursor < len(m.todayTodos) {
		todo := m.todayTodos[m.cursor]
		if err := m.repository.DeleteTodo(todo.ID, todo.Date); err != nil {
			m.err = err
			return m, nil
		}

		// Reload todos and adjust cursor
		m.todayTodos, _ = m.repository.GetTodosForDate(m.selectedDate)
		if m.cursor >= len(m.todayTodos) && len(m.todayTodos) > 0 {
			m.cursor = len(m.todayTodos) - 1
		} else if len(m.todayTodos) == 0 {
			m.cursor = 0
		}
	}
	return m, nil
}

// deleteCurrentGeneralTodo deletes the current general todo
func (m Model) deleteCurrentGeneralTodo() (tea.Model, tea.Cmd) {
	if len(m.generalTodos) > 0 && m.cursor < len(m.generalTodos) {
		todo := m.generalTodos[m.cursor]
		if err := m.repository.DeleteTodo(todo.ID, nil); err != nil {
			m.err = err
			return m, nil
		}

		// Reload todos and adjust cursor
		m.generalTodos, _ = m.repository.GetGeneralTodos()
		if m.cursor >= len(m.generalTodos) && len(m.generalTodos) > 0 {
			m.cursor = len(m.generalTodos) - 1
		} else if len(m.generalTodos) == 0 {
			m.cursor = 0
		}
	}
	return m, nil
}
