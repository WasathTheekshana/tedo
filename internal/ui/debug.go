package ui

import "github.com/WasathTheekshana/tedo/internal/models"

// AddTestData adds some test todos for development
func (m *Model) AddTestData() {
	today := models.TodayString()

	// Add test today todos
	todo1 := models.NewTodo("Morning Exercise", "30 minutes of jogging", &today)
	todo2 := models.NewTodo("Code Review", "Review PR #123", &today)
	todo3 := models.NewTodo("Team Meeting", "Daily standup at 10 AM", &today)

	m.repository.AddTodo(todo1)
	m.repository.AddTodo(todo2)
	m.repository.AddTodo(todo3)

	// Add test general todos
	general1 := models.NewTodo("Learn Go", "Complete Bubble Tea tutorial", nil)
	general2 := models.NewTodo("Read Book", "Finish 'Clean Code'", nil)

	m.repository.AddTodo(general1)
	m.repository.AddTodo(general2)

	// Reload data
	m.todayTodos, _ = m.repository.GetTodosForDate(today)
	m.generalTodos, _ = m.repository.GetGeneralTodos()
}
