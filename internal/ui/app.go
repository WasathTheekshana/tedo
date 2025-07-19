package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/WasathTheekshana/tedo/internal/models"
	"github.com/WasathTheekshana/tedo/internal/storage"
)

// ViewType represents different views in the app
type ViewType int

const (
	TodayView ViewType = iota
	UpcomingView
	CalendarView
	GeneralView
)

// Pagination for the app
const (
	TodosPerPage = 10
)

// Model represents the main application state
type Model struct {
	currentView ViewType
	repository  *storage.Repository

	// View states
	todayTodos    []models.Todo
	upcomingTodos []models.Todo
	generalTodos  []models.Todo
	selectedDate  string
	cursor        int
	calendarState CalendarState

	// Pagination
	todayPage    int
	upcomingPage int
	generalPage  int

	// Input state
	inputState InputState

	// Error handling
	errorState ErrorState

	// UI state
	width  int
	height int

	// Performance optimization
	lastRefresh time.Time
}

// NewModel creates a new application model

func NewModel() Model {
	repo := storage.NewRepository()
	today := models.TodayString()

	// Load initial data
	todayTodos, _ := repo.GetTodosForDate(today)
	upcomingTodos := loadUpcomingTodos(repo, today)
	generalTodos, _ := repo.GetGeneralTodos()

	return Model{
		currentView:   TodayView,
		repository:    repo,
		todayTodos:    todayTodos,
		upcomingTodos: upcomingTodos,
		generalTodos:  generalTodos,
		selectedDate:  today,
		cursor:        0,
		calendarState: NewCalendarState(),
		todayPage:     0,
		upcomingPage:  0,
		generalPage:   0,
		inputState:    NewInputState(),
		errorState:    ErrorState{},
		lastRefresh:   time.Now(),
	}
}

// loadUpcomingTodos loads all todos that are not for today (future dates)
func loadUpcomingTodos(repo *storage.Repository, today string) []models.Todo {
	// For now, we'll load todos from the next 30 days
	// In a real app, you might want to scan the data directory
	var upcomingTodos []models.Todo

	todayTime, _ := time.Parse("2006-01-02", today)

	for i := 1; i <= 30; i++ {
		futureDate := todayTime.AddDate(0, 0, i)
		futureDateStr := futureDate.Format("2006-01-02")

		todos, err := repo.GetTodosForDate(futureDateStr)
		if err == nil && len(todos) > 0 {
			upcomingTodos = append(upcomingTodos, todos...)
		}
	}

	return upcomingTodos
}

// getPaginatedTodos returns the todos for the current page
func (m Model) getPaginatedTodos() ([]models.Todo, int, int) {
	var todos []models.Todo
	var currentPage int

	switch m.currentView {
	case TodayView:
		todos = m.todayTodos
		currentPage = m.todayPage
	case UpcomingView:
		todos = m.upcomingTodos
		currentPage = m.upcomingPage
	case GeneralView:
		todos = m.generalTodos
		currentPage = m.generalPage
	default:
		return []models.Todo{}, 0, 0
	}

	totalPages := (len(todos) + TodosPerPage - 1) / TodosPerPage
	if totalPages == 0 {
		totalPages = 1
	}

	start := currentPage * TodosPerPage
	end := start + TodosPerPage
	if end > len(todos) {
		end = len(todos)
	}

	if start >= len(todos) {
		return []models.Todo{}, currentPage, totalPages
	}

	return todos[start:end], currentPage, totalPages
}

// getAbsoluteCursor returns the absolute cursor position (across all pages)
func (m Model) getAbsoluteCursor() int {
	switch m.currentView {
	case TodayView:
		return m.todayPage*TodosPerPage + m.cursor
	case UpcomingView:
		return m.upcomingPage*TodosPerPage + m.cursor
	case GeneralView:
		return m.generalPage*TodosPerPage + m.cursor
	default:
		return m.cursor
	}
}

// resetPagination resets pagination when todos are modified
func (m *Model) resetPagination() {
	switch m.currentView {
	case TodayView:
		totalPages := (len(m.todayTodos) + TodosPerPage - 1) / TodosPerPage
		if totalPages == 0 {
			totalPages = 1
		}
		if m.todayPage >= totalPages {
			m.todayPage = totalPages - 1
		}
		if m.todayPage < 0 {
			m.todayPage = 0
		}
	case UpcomingView:
		totalPages := (len(m.upcomingTodos) + TodosPerPage - 1) / TodosPerPage
		if totalPages == 0 {
			totalPages = 1
		}
		if m.upcomingPage >= totalPages {
			m.upcomingPage = totalPages - 1
		}
		if m.upcomingPage < 0 {
			m.upcomingPage = 0
		}
	case GeneralView:
		totalPages := (len(m.generalTodos) + TodosPerPage - 1) / TodosPerPage
		if totalPages == 0 {
			totalPages = 1
		}
		if m.generalPage >= totalPages {
			m.generalPage = totalPages - 1
		}
		if m.generalPage < 0 {
			m.generalPage = 0
		}
	}

	// Reset cursor if out of bounds
	paginatedTodos, _, _ := m.getPaginatedTodos()
	if m.cursor >= len(paginatedTodos) {
		m.cursor = 0
		if len(paginatedTodos) > 0 {
			m.cursor = len(paginatedTodos) - 1
		}
	}
}

// reloadTodos reloads todos after changes
func (m *Model) reloadTodos() {
	// Only reload if significant time has passed or forced
	if time.Since(m.lastRefresh) < 100*time.Millisecond {
		return
	}

	today := models.TodayString()
	m.todayTodos, _ = m.repository.GetTodosForDate(today)
	m.upcomingTodos = loadUpcomingTodos(m.repository, today)
	m.generalTodos, _ = m.repository.GetGeneralTodos()
	m.lastRefresh = time.Now()
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
		m.errorState.SetErrorMessage("Error while updating")
		return m, nil
	}

	return m, nil
}

// Update handleKeyPress to include UpcomingView
func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle input mode first
	if m.inputState.mode != NavigationMode {
		return m.handleInputMode(msg)
	}

	// Handle QUIT keys FIRST
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	}

	// Handle ARROW KEYS for menu/tab navigation ONLY
	switch msg.String() {
	case "left", "right":
		if msg.String() == "right" {
			return m.switchToNextView(), nil
		} else {
			return m.switchToPrevView(), nil
		}
	}

	// Handle view-specific keys (these will use hjkl)
	switch m.currentView {
	case TodayView:
		return m.handleTodayViewKeys(msg)
	case UpcomingView:
		return m.handleUpcomingViewKeys(msg)
	case CalendarView:
		return m.handleCalendarViewKeys(msg)
	case GeneralView:
		return m.handleGeneralViewKeys(msg)
	}

	// Handle remaining global navigation keys
	switch msg.String() {
	case "tab":
		return m.switchToNextView(), nil
	case "shift+tab":
		return m.switchToPrevView(), nil
	case "1":
		m.currentView = TodayView
		m.cursor = 0
		return m, nil
	case "2":
		m.currentView = UpcomingView
		m.cursor = 0
		return m, nil
	case "3":
		m.currentView = CalendarView
		m.cursor = 0
		return m, nil
	case "4":
		m.currentView = GeneralView
		m.cursor = 0
		return m, nil
	}

	return m, nil
}

// handleInputMode handles keys when in input mode
func (m Model) handleInputMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.inputState.ExitInputMode()
		m.errorState.ClearError() // Clear errors when canceling
		return m, nil
	case "ctrl+s", "enter":
		return m.handleSaveTodo()
	case "tab":
		m.inputState.SwitchField()
		m.errorState.ClearError() // Clear errors when switching fields
		return m, nil
	case "ctrl+a":
		// Select all text in current field
		m.inputState.cursor = len(*m.inputState.getCurrentField())
		return m, nil
	default:
		m.inputState.HandleInput(msg.String())
		m.errorState.ClearError() // Clear errors when typing
		return m, nil
	}
}

// handleSaveTodo saves the current input as a todo
func (m Model) handleSaveTodo() (tea.Model, tea.Cmd) {
	if !m.inputState.IsValid() {
		m.errorState.SetErrorMessage("title is required")
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
	// Clean input
	title := CleanInput(m.inputState.title)
	description := CleanInput(m.inputState.description)

	// Validate input
	if errors := ValidateTodoInput(title, description); len(errors) > 0 {
		m.errorState.SetErrorMessage(FormatValidationErrors(errors))
		return m, nil
	}

	var date *string
	if m.currentView == TodayView {
		date = &m.selectedDate
	} else if m.currentView == CalendarView {
		selectedDate := m.calendarState.getSelectedDate()
		date = &selectedDate
	} else if m.currentView == UpcomingView {
		date = &m.selectedDate
	}

	newTodo := models.NewTodo(title, description, date)

	if err := m.repository.AddTodo(newTodo); err != nil {
		m.errorState.SetError(fmt.Errorf("failed to save todo: %w", err))
		return m, nil
	}

	// Success - reload and clear error
	m.errorState.ClearError()
	m.reloadTodos()
	m.resetPagination()
	m.inputState.ExitInputMode()
	return m, nil
}

// saveEditedTodo updates an existing todo
func (m Model) saveEditedTodo() (tea.Model, tea.Cmd) {
	if m.inputState.editingTodo == nil {
		m.errorState.SetErrorMessage("No todo being edited")
		return m, nil
	}

	// Clean input
	title := CleanInput(m.inputState.title)
	description := CleanInput(m.inputState.description)

	// Validate input
	if errors := ValidateTodoInput(title, description); len(errors) > 0 {
		m.errorState.SetErrorMessage(FormatValidationErrors(errors))
		return m, nil
	}

	// Update the todo
	m.inputState.editingTodo.Title = title
	m.inputState.editingTodo.Description = description

	if err := m.repository.UpdateTodo(*m.inputState.editingTodo); err != nil {
		m.errorState.SetError(fmt.Errorf("failed to update todo: %w", err))
		return m, nil
	}

	// Success - reload and clear error
	m.errorState.ClearError()
	m.reloadTodos()
	m.resetPagination()
	m.inputState.ExitInputMode()
	return m, nil
}

// View implements tea.Model - renders the current view
func (m Model) View() string {
	// Show global errors if not in input mode
	if m.inputState.mode == NavigationMode {
		if errorMsg := m.errorState.GetError(); errorMsg != "" {
			errorDisplay := errorStyle.Render("⚠ " + errorMsg)
			return lipgloss.JoinVertical(
				lipgloss.Left,
				m.renderHeader(),
				errorDisplay,
				"",
				m.getCurrentViewContent(),
				m.renderFooter(),
			)
		}
	}

	var content string
	switch m.currentView {
	case TodayView:
		content = m.renderTodayView()
	case UpcomingView:
		content = m.renderUpcomingView()
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

// Update view switching
func (m Model) switchToNextView() Model {
	switch m.currentView {
	case TodayView:
		m.currentView = UpcomingView
	case UpcomingView:
		m.currentView = CalendarView
	case CalendarView:
		m.currentView = GeneralView
	case GeneralView:
		m.currentView = TodayView
	}
	m.cursor = 0
	return m
}

func (m Model) switchToPrevView() Model {
	switch m.currentView {
	case TodayView:
		m.currentView = GeneralView
	case UpcomingView:
		m.currentView = TodayView
	case CalendarView:
		m.currentView = UpcomingView
	case GeneralView:
		m.currentView = CalendarView
	}
	m.cursor = 0
	return m
}
