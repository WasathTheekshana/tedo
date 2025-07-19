package storage

import (
	"fmt"

	"github.com/WasathTheekshana/tedo/internal/models"
)

// Repository provides high-level operations for todo management
type Repository struct {
	storage *JSONStorage
}

// NewRepository creates a new repository instance
func NewRepository() *Repository {
	return &Repository{
		storage: NEWJSONStorage(),
	}
}

// GetTodosForDate retrieves todos for a specific date
func (r *Repository) GetTodosForDate(date string) ([]models.Todo, error) {
	return r.storage.LoadTodos(&date)
}

// GetGeneralTodos retrieves general todos (no specific date)
func (r *Repository) GetGeneralTodos() ([]models.Todo, error) {
	return r.storage.LoadTodos(nil)
}

// AddTodo adds a new todo and saves it
func (r *Repository) AddTodo(todo models.Todo) error {
	var todos []models.Todo
	var err error

	if todo.IsGeneral() {
		todos, err = r.GetGeneralTodos()
	} else {
		todos, err = r.GetTodosForDate(*todo.Date)
	}

	if err != nil {
		return fmt.Errorf("failed to load existing todos: %w", err)
	}

	todos = append(todos, todo)
	return r.storage.SaveTodos(todos, todo.Date)
}

// UpdateTodo updates an existing todo
func (r *Repository) UpdateTodo(updatedTodo models.Todo) error {
	var todos []models.Todo
	var err error

	if updatedTodo.IsGeneral() {
		todos, err = r.GetGeneralTodos()
	} else {
		todos, err = r.GetTodosForDate(*updatedTodo.Date)
	}

	if err != nil {
		return fmt.Errorf("failed to load existing todos: %w", err)
	}

	// Find and update the todo
	found := false
	for i, todo := range todos {
		if todo.ID == updatedTodo.ID {
			todos[i] = updatedTodo
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("todo with ID %s not found", updatedTodo.ID)
	}

	return r.storage.SaveTodos(todos, updatedTodo.Date)
}

// DeleteTodo removes a todo
func (r *Repository) DeleteTodo(todoID string, date *string) error {
	todos, err := r.storage.LoadTodos(date)
	if err != nil {
		return fmt.Errorf("failed to load todos: %w", err)
	}

	// Find and remove the todo
	found := false
	for i, todo := range todos {
		if todo.ID == todoID {
			todos = append(todos[:i], todos[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("todo with ID %s not found", todoID)
	}

	return r.storage.SaveTodos(todos, date)
}

// GetTodoCountForDate returns the number of todos for a specific date
func (r *Repository) GetTodoCountForDate(date string) (int, error) {
	todos, err := r.GetTodosForDate(date)
	if err != nil {
		return 0, err
	}
	return len(todos), nil
}
