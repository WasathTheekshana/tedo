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
