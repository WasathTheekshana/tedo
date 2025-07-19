package main

import (
	"fmt"
	"log"

	"github.com/WasathTheekshana/tedo/internal/models"
	"github.com/WasathTheekshana/tedo/internal/storage"
)

func main() {
	fmt.Println("---------------------------------------------------------")

	fmt.Println("Tedo - Testing Storage")

	repo := storage.NewRepository()
	today := models.TodayString()

	// Test adding todos
	generalTodo := models.NewTodo("Learn Go Patterns", "Study design patterns in Go", nil)
	todayTodo := models.NewTodo("Code Review", "Review PR #123", &today)

	fmt.Println("Adding todos...")
	if err := repo.AddTodo(generalTodo); err != nil {
		log.Fatalf("Failed to add general todo: %v", err)
	}

	if err := repo.AddTodo(todayTodo); err != nil {
		log.Fatalf("Failed to add today's todo: %v", err)
	}

	// Test loading todos
	fmt.Println("\nLoading general todos...")
	generalTodos, err := repo.GetGeneralTodos()
	if err != nil {
		log.Fatalf("Failed to load general todos: %v", err)
	}
	fmt.Printf("Found %d general todos\n", len(generalTodos))
	for _, todo := range generalTodos {
		fmt.Printf("- %s: %s (Completed: %v)\n", todo.Title, todo.Description, todo.Completed)
	}

	fmt.Println("\nLoading today's todos...")
	todayTodos, err := repo.GetTodosForDate(today)
	if err != nil {
		log.Fatalf("Failed to load today's todos: %v", err)
	}
	fmt.Printf("Found %d todos for %s\n", len(todayTodos), today)
	for _, todo := range todayTodos {
		fmt.Printf("- %s: %s (Completed: %v)\n", todo.Title, todo.Description, todo.Completed)
	}

	// Test toggle and update
	if len(todayTodos) > 0 {
		fmt.Println("\nToggling first todo...")
		todayTodos[0].Toggle()
		if err := repo.UpdateTodo(todayTodos[0]); err != nil {
			log.Fatalf("Failed to update todo: %v", err)
		}
		fmt.Printf("Todo '%s' is now completed: %v\n", todayTodos[0].Title, todayTodos[0].Completed)
	}

	// Test count
	count, err := repo.GetTodoCountForDate(today)
	if err != nil {
		log.Fatalf("Failed to get todo count: %v", err)
	}
	fmt.Printf("\nTotal todos for %s: %d\n", today, count)
}
