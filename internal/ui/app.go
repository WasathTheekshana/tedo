package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/WasathTheekshana/tedo/internal/models"
	"github.com/WasathTheekshana/tedo/internal/storage"
)

// ViewType represents different views in the app
type ViewType int

const (
	TodayView ViewType = iota
	CalendarView
	GeneralView
)

// Model represents the main application state
type Model struct {
	currentView ViewType
	repository  *storage.Repository

	// View states
	todayTodos   []models.Todo
	generalTodos []models.Todo
	selectedDate string // Currently selected date (YYYY-MM-DD)
	cursor       int    // Current cursor position in lists

	// Input state
	inputState InputState

	// UI state
	width  int
	height int
	err    error
}

// NewModel creates a new application model
func NewModel() Model {
	repo := storage.NewRepository()
	today := models.TodayString()

	// Load initial data
	todayTodos, _ := repo.GetTodosForDate(today)
	generalTodos, _ := repo.GetGeneralTodos()

	return Model{
		currentView:  TodayView,
		repository:   repo,
		todayTodos:   todayTodos,
		generalTodos: generalTodos,
		selectedDate: today,
		cursor:       0,
		inputState:   NewInputState(),
	}
}

// Init implements tea.Model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model - handles all key presses and messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyPress(msg)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case error:
		m.err = msg
		return m, nil
	}

	return m, nil
}

// handleKeyPress processes keyboard input
func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle input mode first
	if m.inputState.mode != NavigationMode {
		return m.handleInputMode(msg)
	}

	// Global navigation keys
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "tab", "l", "right":
		return m.switchToNextView(), nil
	case "shift+tab", "h", "left":
		return m.switchToPrevView(), nil
	case "1":
		m.currentView = TodayView
		m.cursor = 0
		return m, nil
	case "2":
		m.currentView = CalendarView
		m.cursor = 0
		return m, nil
	case "3":
		m.currentView = GeneralView
		m.cursor = 0
		return m, nil
	}

	// Handle view-specific keys
	switch m.currentView {
	case TodayView:
		return m.handleTodayViewKeys(msg)
	case CalendarView:
		return m.handleCalendarViewKeys(msg)
	case GeneralView:
		return m.handleGeneralViewKeys(msg)
	}

	return m, nil
}

// handleInputMode handles keys when in input mode
func (m Model) handleInputMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.inputState.ExitInputMode()
		return m, nil
	case "enter":
		return m.handleSaveTodo()
	case "tab":
		m.inputState.SwitchField()
		return m, nil
	default:
		m.inputState.HandleInput(msg.String())
		return m, nil
	}
}

// handleSaveTodo saves the current input as a todo
func (m Model) handleSaveTodo() (tea.Model, tea.Cmd) {
	if !m.inputState.IsValid() {
		m.err = fmt.Errorf("title is required")
		return m, nil
	}

	switch m.inputState.mode {
	case AddTodoMode:
		return m.saveNewTodo()
	case EditTodoMode:
		return m.saveEditedTodo()
	}

	return m, nil
}

// saveNewTodo creates and saves a new todo
func (m Model) saveNewTodo() (tea.Model, tea.Cmd) {
	var date *string
	if m.currentView == TodayView {
		date = &m.selectedDate
	}
	// For GeneralView, date remains nil

	newTodo := models.NewTodo(m.inputState.title, m.inputState.description, date)

	if err := m.repository.AddTodo(newTodo); err != nil {
		m.err = err
		return m, nil
	}

	// Reload the appropriate todo list
	if date != nil {
		m.todayTodos, _ = m.repository.GetTodosForDate(*date)
	} else {
		m.generalTodos, _ = m.repository.GetGeneralTodos()
	}

	m.inputState.ExitInputMode()
	return m, nil
}

// saveEditedTodo updates an existing todo
func (m Model) saveEditedTodo() (tea.Model, tea.Cmd) {
	if m.inputState.editingTodo == nil {
		m.err = fmt.Errorf("no todo being edited")
		return m, nil
	}

	// Update the todo
	m.inputState.editingTodo.Title = m.inputState.title
	m.inputState.editingTodo.Description = m.inputState.description

	if err := m.repository.UpdateTodo(*m.inputState.editingTodo); err != nil {
		m.err = err
		return m, nil
	}

	// Reload the appropriate todo list
	if m.inputState.editingTodo.IsGeneral() {
		m.generalTodos, _ = m.repository.GetGeneralTodos()
	} else {
		m.todayTodos, _ = m.repository.GetTodosForDate(*m.inputState.editingTodo.Date)
	}

	m.inputState.ExitInputMode()
	return m, nil
}

// View implements tea.Model - renders the current view
func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\nPress q to quit.", m.err)
	}

	var content string

	switch m.currentView {
	case TodayView:
		content = m.renderTodayView()
	case CalendarView:
		content = m.renderCalendarView()
	case GeneralView:
		content = m.renderGeneralView()
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderHeader(),
		content,
		m.renderFooter(),
	)
}

// Helper methods for view switching
func (m Model) switchToNextView() Model {
	switch m.currentView {
	case TodayView:
		m.currentView = CalendarView
	case CalendarView:
		m.currentView = GeneralView
	case GeneralView:
		m.currentView = TodayView
	}
	m.cursor = 0 // Reset cursor when switching views
	return m
}

func (m Model) switchToPrevView() Model {
	switch m.currentView {
	case TodayView:
		m.currentView = GeneralView
	case CalendarView:
		m.currentView = TodayView
	case GeneralView:
		m.currentView = CalendarView
	}
	m.cursor = 0 // Reset cursor when switching views
	return m
}
