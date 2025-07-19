package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/WasathTheekshana/tedo/internal/models"
)

const (
	DataDir      = "data"
	GeneralFile  = "general.json"
	DatedFileExt = ".json"
)

// JSONStorage handles file-based JSON storage
type JSONStorage struct {
	dataDir string
}

// NewJSONStorage create a new JSON storage instance
func NEWJSONStorage() *JSONStorage {
	return &JSONStorage{
		dataDir: DataDir,
	}
}

// ensureDataDir creates the data directory is it doesn't exist
func (s *JSONStorage) ensureDataDir() error {
	if _, err := os.Stat(s.dataDir); os.IsNotExist(err) {
		return os.MkdirAll(s.dataDir, 0o755)
	}
	return nil
}

// getFilePath returns the file path for a given data or general todos
func (s *JSONStorage) getFilePath(data *string) string {
	if data == nil {
		return filepath.Join(s.dataDir, GeneralFile)
	}
	return filepath.Join(s.dataDir, *data+DatedFileExt)
}

// SaveTodos saves todos to the appropriate JSON file
func (s *JSONStorage) SaveTodos(todos []models.Todo, date *string) error {
	if err := s.ensureDataDir(); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	filePath := s.getFilePath(date)
	todoList := models.TodoList{Todos: todos}

	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal todos: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

// LoadTodos loads todos from the appropriate JSON file
func (s *JSONStorage) LoadTodos(date *string) ([]models.Todo, error) {
	filePath := s.getFilePath(date)

	// If file doesn't exist, return empty slice (not an error)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []models.Todo{}, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var todoList models.TodoList
	if err := json.Unmarshal(data, &todoList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal todos from %s: %w", filePath, err)
	}

	return todoList.Todos, nil
}
