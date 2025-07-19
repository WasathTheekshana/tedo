package ui

import (
	"fmt"
	"runtime"
	"time"
)

// PerformanceInfo holds performance metrics
type PerformanceInfo struct {
	MemoryUsage    uint64
	LastUpdateTime time.Duration
	TodoCount      int
}

// GetPerformanceInfo returns current performance metrics
func (m Model) GetPerformanceInfo() PerformanceInfo {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	totalTodos := len(m.todayTodos) + len(m.upcomingTodos) + len(m.generalTodos)

	return PerformanceInfo{
		MemoryUsage:    memStats.Alloc / 1024, // KB
		LastUpdateTime: time.Since(m.lastRefresh),
		TodoCount:      totalTodos,
	}
}

// RenderDebugInfo renders debug information (hidden feature)
func (m Model) RenderDebugInfo() string {
	info := m.GetPerformanceInfo()

	debugInfo := fmt.Sprintf(
		"Debug: Memory: %dKB | Todos: %d | View: %s | Cursor: %d",
		info.MemoryUsage,
		info.TodoCount,
		getViewName(m.currentView),
		m.cursor,
	)

	return mutedStyle.Render(debugInfo)
}
