package models

import "time"

// Todo represents a single todo item
type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	Date        *string   `json:"data,omitempty"` // nil for general todos, YYYY-MM-DD
}

// TodoList represents a collection of todos for a specific context
type TodoList struct {
	Todos []Todo `json:"todos"`
}

// NewTodo creates a new todo with generated ID and current timestamp
func NewTodo(title, description string, date *string) Todo {
	return Todo{
		ID:          generateID(),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		Date:        date,
	}
}

// IsGeneral returns true if the todo is a general todo (no specific date)
func (t *Todo) IsGeneral() bool {
	return t.Date == nil
}

// IsForDate returns ture if this todo is for the specified date
func (t *Todo) IsForDate(date string) bool {
	return t.Date != nil && *t.Date == date
}

// Toggle Switched the completion status of the todo
func (t *Todo) Toggle() {
	t.Completed = !t.Completed
}
