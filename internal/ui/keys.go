package ui

import tea "github.com/charmbracelet/bubbletea"

// handleTodayViewKeys handles keys specific to today view
func (m Model) handleTodayViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle tab switching first
	switch msg.String() {
	case "l", "right":
		return m.switchToNextView(), nil
	case "h", "left":
		return m.switchToPrevView(), nil
	}

	paginatedTodos, currentPage, totalPages := m.getPaginatedTodos()

	switch msg.String() {
	case "j", "down":
		if len(paginatedTodos) > 0 && m.cursor < len(paginatedTodos)-1 {
			m.cursor++
		} else if len(paginatedTodos) > 0 && m.cursor == len(paginatedTodos)-1 && currentPage < totalPages-1 {
			// Go to next page
			m.todayPage++
			m.cursor = 0
		}
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		} else if m.cursor == 0 && currentPage > 0 {
			// Go to previous page
			m.todayPage--
			newPaginatedTodos, _, _ := m.getPaginatedTodos()
			m.cursor = len(newPaginatedTodos) - 1
		}
	case "ctrl+f", "page_down":
		if currentPage < totalPages-1 {
			m.todayPage++
			m.cursor = 0
		}
	case "ctrl+b", "page_up":
		if currentPage > 0 {
			m.todayPage--
			m.cursor = 0
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
	case "c":
		// Press 'c' to go to calendar
		m.currentView = CalendarView
		return m, nil
	}
	return m, nil
}

// handleCalendarViewKeys handles keys specific to calendar view
func (m Model) handleCalendarViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle calendar-specific navigation (these take priority over global keys)
	switch msg.String() {
	case "j", "down":
		m.calendarState.moveCursor(1, 0)
		return m, nil
	case "k", "up":
		m.calendarState.moveCursor(-1, 0)
		return m, nil
	case "h", "left":
		m.calendarState.moveCursor(0, -1)
		return m, nil
	case "l", "right":
		m.calendarState.moveCursor(0, 1)
		return m, nil
	case "n", ">":
		m.calendarState.moveToNextMonth()
		return m, nil
	case "p", "<":
		m.calendarState.moveToPrevMonth()
		return m, nil
	case "t":
		m.calendarState.moveToToday()
		return m, nil
	case "enter":
		// Switch to today view with selected date
		m.selectedDate = m.calendarState.getSelectedDate()
		m.todayTodos, _ = m.repository.GetTodosForDate(m.selectedDate)
		m.currentView = TodayView
		m.cursor = 0
		m.todayPage = 0
		return m, nil
	case "i":
		// Add todo for selected date
		m.selectedDate = m.calendarState.getSelectedDate()
		m.inputState.StartAddMode()
		return m, nil
	case "tab":
		// Tab switches to next view
		return m.switchToNextView(), nil
	case "shift+tab":
		// Shift+tab switches to previous view
		return m.switchToPrevView(), nil
	}
	return m, nil
}

// handleGeneralViewKeys handles keys specific to general view
func (m Model) handleGeneralViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle tab switching first
	switch msg.String() {
	case "l", "right":
		return m.switchToNextView(), nil
	case "h", "left":
		return m.switchToPrevView(), nil
	}

	paginatedTodos, currentPage, totalPages := m.getPaginatedTodos()

	switch msg.String() {
	case "j", "down":
		if len(paginatedTodos) > 0 && m.cursor < len(paginatedTodos)-1 {
			m.cursor++
		} else if len(paginatedTodos) > 0 && m.cursor == len(paginatedTodos)-1 && currentPage < totalPages-1 {
			// Go to next page
			m.generalPage++
			m.cursor = 0
		}
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		} else if m.cursor == 0 && currentPage > 0 {
			// Go to previous page
			m.generalPage--
			newPaginatedTodos, _, _ := m.getPaginatedTodos()
			m.cursor = len(newPaginatedTodos) - 1
		}
	case "ctrl+f", "page_down":
		if currentPage < totalPages-1 {
			m.generalPage++
			m.cursor = 0
		}
	case "ctrl+b", "page_up":
		if currentPage > 0 {
			m.generalPage--
			m.cursor = 0
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
	case "c":
		// Press 'c' to go to calendar
		m.currentView = CalendarView
		return m, nil
	}
	return m, nil
}

// toggleCurrentTodo toggles completion of current today todo
func (m Model) toggleCurrentTodo() Model {
	paginatedTodos, _, _ := m.getPaginatedTodos()
	if len(paginatedTodos) > 0 && m.cursor < len(paginatedTodos) {
		absoluteIndex := m.getAbsoluteCursor()
		if absoluteIndex < len(m.todayTodos) {
			m.todayTodos[absoluteIndex].Toggle()
			if err := m.repository.UpdateTodo(m.todayTodos[absoluteIndex]); err != nil {
				m.err = err
			}
		}
	}
	return m
}

// toggleCurrentGeneralTodo toggles completion of current general todo
func (m Model) toggleCurrentGeneralTodo() Model {
	paginatedTodos, _, _ := m.getPaginatedTodos()
	if len(paginatedTodos) > 0 && m.cursor < len(paginatedTodos) {
		absoluteIndex := m.getAbsoluteCursor()
		if absoluteIndex < len(m.generalTodos) {
			m.generalTodos[absoluteIndex].Toggle()
			if err := m.repository.UpdateTodo(m.generalTodos[absoluteIndex]); err != nil {
				m.err = err
			}
		}
	}
	return m
}

// editCurrentTodo starts editing the current today todo
func (m Model) editCurrentTodo() (tea.Model, tea.Cmd) {
	paginatedTodos, _, _ := m.getPaginatedTodos()
	if len(paginatedTodos) > 0 && m.cursor < len(paginatedTodos) {
		absoluteIndex := m.getAbsoluteCursor()
		if absoluteIndex < len(m.todayTodos) {
			m.inputState.StartEditMode(&m.todayTodos[absoluteIndex])
		}
	}
	return m, nil
}

// editCurrentGeneralTodo starts editing the current general todo
func (m Model) editCurrentGeneralTodo() (tea.Model, tea.Cmd) {
	paginatedTodos, _, _ := m.getPaginatedTodos()
	if len(paginatedTodos) > 0 && m.cursor < len(paginatedTodos) {
		absoluteIndex := m.getAbsoluteCursor()
		if absoluteIndex < len(m.generalTodos) {
			m.inputState.StartEditMode(&m.generalTodos[absoluteIndex])
		}
	}
	return m, nil
}

// deleteCurrentTodo deletes the current today todo
func (m Model) deleteCurrentTodo() (tea.Model, tea.Cmd) {
	paginatedTodos, _, _ := m.getPaginatedTodos()
	if len(paginatedTodos) > 0 && m.cursor < len(paginatedTodos) {
		absoluteIndex := m.getAbsoluteCursor()
		if absoluteIndex < len(m.todayTodos) {
			todo := m.todayTodos[absoluteIndex]
			if err := m.repository.DeleteTodo(todo.ID, todo.Date); err != nil {
				m.err = err
				return m, nil
			}

			// Reload todos and reset pagination
			m.todayTodos, _ = m.repository.GetTodosForDate(m.selectedDate)
			m.resetPagination()
		}
	}
	return m, nil
}

// deleteCurrentGeneralTodo deletes the current general todo
func (m Model) deleteCurrentGeneralTodo() (tea.Model, tea.Cmd) {
	paginatedTodos, _, _ := m.getPaginatedTodos()
	if len(paginatedTodos) > 0 && m.cursor < len(paginatedTodos) {
		absoluteIndex := m.getAbsoluteCursor()
		if absoluteIndex < len(m.generalTodos) {
			todo := m.generalTodos[absoluteIndex]
			if err := m.repository.DeleteTodo(todo.ID, nil); err != nil {
				m.err = err
				return m, nil
			}

			// Reload todos and reset pagination
			m.generalTodos, _ = m.repository.GetGeneralTodos()
			m.resetPagination()
		}
	}
	return m, nil
}
