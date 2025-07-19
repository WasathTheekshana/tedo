package main

import (
	"fmt"

	"github.com/WasathTheekshana/tedo/internal/models"
)

func main() {
	fmt.Println("Tedo - Testing Models")

	// Test general todo
	generalTodo := models.NewTodo("Learn GO", "I need to learn golang. IDK why I need this!", nil)
	fmt.Printf("General Todo: %+v\n", generalTodo)
	fmt.Printf("Is General: %+v\n", generalTodo.IsGeneral())

	// Test dated todo
	today := models.TodayString()
	datedTodo := models.NewTodo("Learn Rust", "I need to learn rust in near future", &today)
	fmt.Printf("Dated Todo: %+v\n", datedTodo)
	fmt.Printf("Is General: %+v\n", datedTodo.IsGeneral())
	fmt.Printf("Before Toggle: %+v\n", datedTodo.Completed)

	// Test toggle
	datedTodo.Toggle()
	fmt.Printf("After Toggle: %+v\n", datedTodo.Completed)

	// Test date utils
	fmt.Printf("Today: %s\n", models.TodayString())
}
