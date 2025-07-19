package ui

import (
	"strings"

	"github.com/WasathTheekshana/tedo/internal/models"
)

// InputMode represents different input states
type InputMode int

const (
	NavigationMode InputMode = iota
	AddTodoMode
	EditTodoMode
)

// InputState holds the current input state
type InputState struct {
	mode        InputMode
	title       string
	description string
	editingTodo *models.Todo
	editField   int // 0 = title, 1 = description
	cursor      int // cursor position in input field
}

// NewInputState creates a new input state
func NewInputState() InputState {
	return InputState{
		mode:        NavigationMode,
		title:       "",
		description: "",
		editField:   0,
		cursor:      0,
	}
}

// StartAddMode starts adding a new todo
func (s *InputState) StartAddMode() {
	s.mode = AddTodoMode
	s.title = ""
	s.description = ""
	s.editField = 0
	s.cursor = 0
}

// StartEditMode starts editing an existing todo
func (s *InputState) StartEditMode(todo *models.Todo) {
	s.mode = EditTodoMode
	s.editingTodo = todo
	s.title = todo.Title
	s.description = todo.Description
	s.editField = 0
	s.cursor = len(s.title)
}

// ExitInputMode exits any input mode
func (s *InputState) ExitInputMode() {
	s.mode = NavigationMode
	s.title = ""
	s.description = ""
	s.editingTodo = nil
	s.editField = 0
	s.cursor = 0
}

// HandleInput processes input characters
func (s *InputState) HandleInput(key string) {
	currentField := s.getCurrentField()

	switch key {
	case "backspace":
		if s.cursor > 0 && len(*currentField) > 0 {
			*currentField = (*currentField)[:s.cursor-1] + (*currentField)[s.cursor:]
			s.cursor--
		}
	case "delete":
		if s.cursor < len(*currentField) {
			*currentField = (*currentField)[:s.cursor] + (*currentField)[s.cursor+1:]
		}
	case "left":
		if s.cursor > 0 {
			s.cursor--
		}
	case "right":
		if s.cursor < len(*currentField) {
			s.cursor++
		}
	case "home":
		s.cursor = 0
	case "end":
		s.cursor = len(*currentField)
	default:
		// Regular character input
		if len(key) == 1 && key >= " " && key <= "~" {
			*currentField = (*currentField)[:s.cursor] + key + (*currentField)[s.cursor:]
			s.cursor++
		}
	}
}

// SwitchField switches between title and description fields
func (s *InputState) SwitchField() {
	if s.editField == 0 {
		s.editField = 1
		s.cursor = len(s.description)
	} else {
		s.editField = 0
		s.cursor = len(s.title)
	}
}

// getCurrentField returns pointer to the currently edited field
func (s *InputState) getCurrentField() *string {
	if s.editField == 0 {
		return &s.title
	}
	return &s.description
}

// IsValid returns true if the input is valid for saving
func (s *InputState) IsValid() bool {
	return strings.TrimSpace(s.title) != ""
}
