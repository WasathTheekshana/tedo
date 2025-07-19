package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/WasathTheekshana/tedo/internal/models"
	"github.com/charmbracelet/lipgloss"
)

// CalendarState holds calendar-specific state
type CalendarState struct {
	currentMonth time.Time
	selectedDay  int
	cursor       int // 0-6 for days of week, then by weeks
	cursorRow    int // week row (0-5)
	cursorCol    int // day column (0-6)
}

// NewCalendarState creates a new calendar state for the current month
func NewCalendarState() CalendarState {
	now := time.Now()
	currentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	return CalendarState{
		currentMonth: currentMonth,
		selectedDay:  now.Day(),
		cursor:       0,
		cursorRow:    0,
		cursorCol:    int(now.Weekday()),
	}
}

// getDaysInCurrentMonth returns the number of days in the current month
func (c *CalendarState) getDaysInCurrentMonth() int {
	return models.GetDaysInMonth(c.currentMonth.Year(), c.currentMonth.Month())
}

// getFirstWeekday returns the weekday of the first day of the month (0=Sunday, 6=Saturday)
func (c *CalendarState) getFirstWeekday() int {
	_, weekday := models.GetFirstDayOfMonth(c.currentMonth.Year(), c.currentMonth.Month())
	return int(weekday)
}

// moveToNextMonth moves to the next month
func (c *CalendarState) moveToNextMonth() {
	c.currentMonth = c.currentMonth.AddDate(0, 1, 0)
	c.selectedDay = 1
	c.cursorRow = 0
	c.cursorCol = c.getFirstWeekday()
}

// moveToPrevMonth moves to the previous month
func (c *CalendarState) moveToPrevMonth() {
	c.currentMonth = c.currentMonth.AddDate(0, -1, 0)
	c.selectedDay = 1
	c.cursorRow = 0
	c.cursorCol = c.getFirstWeekday()
}

// moveToToday moves cursor to today's date if it's in current month
func (c *CalendarState) moveToToday() {
	now := time.Now()
	if now.Year() == c.currentMonth.Year() && now.Month() == c.currentMonth.Month() {
		c.selectedDay = now.Day()
		c.updateCursorPosition()
	}
}

// updateCursorPosition updates cursor position based on selected day
func (c *CalendarState) updateCursorPosition() {
	firstWeekday := c.getFirstWeekday()
	dayIndex := c.selectedDay - 1 + firstWeekday
	c.cursorRow = dayIndex / 7
	c.cursorCol = dayIndex % 7
}

// moveCursor moves the cursor and updates selected day
func (c *CalendarState) moveCursor(deltaRow, deltaCol int) {
	newRow := c.cursorRow + deltaRow
	newCol := c.cursorCol + deltaCol

	// Handle horizontal wrapping
	if newCol < 0 {
		newCol = 6
		newRow--
	} else if newCol > 6 {
		newCol = 0
		newRow++
	}

	// Handle vertical bounds
	if newRow < 0 {
		newRow = 0
		newCol = c.getFirstWeekday()
	} else if newRow > 5 {
		newRow = 5
		newCol = 6
	}

	c.cursorRow = newRow
	c.cursorCol = newCol

	// Calculate the day based on cursor position
	firstWeekday := c.getFirstWeekday()
	dayIndex := c.cursorRow*7 + c.cursorCol - firstWeekday + 1

	daysInMonth := c.getDaysInCurrentMonth()
	if dayIndex >= 1 && dayIndex <= daysInMonth {
		c.selectedDay = dayIndex
	}
}

// getSelectedDate returns the currently selected date as YYYY-MM-DD string
func (c *CalendarState) getSelectedDate() string {
	selectedDate := time.Date(
		c.currentMonth.Year(),
		c.currentMonth.Month(),
		c.selectedDay,
		0, 0, 0, 0, time.UTC,
	)
	return models.FormatDate(selectedDate)
}

// renderCalendar renders the calendar grid
func (m Model) renderCalendar() string {
	cal := m.calendarState
	var lines []string

	// Month/Year header
	monthYear := cal.currentMonth.Format("January 2006")
	header := selectedItemStyle.Render(fmt.Sprintf("📅 %s", monthYear))
	lines = append(lines, header)
	lines = append(lines, "")

	// Day headers
	dayHeaders := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	headerLine := "  " + strings.Join(dayHeaders, "  ")
	lines = append(lines, mutedStyle.Render(headerLine))

	// Calendar grid
	firstWeekday := cal.getFirstWeekday()
	daysInMonth := cal.getDaysInCurrentMonth()

	for week := 0; week < 6; week++ {
		var weekDays []string
		hasValidDay := false

		for day := 0; day < 7; day++ {
			dayNum := week*7 + day - firstWeekday + 1

			if dayNum < 1 || dayNum > daysInMonth {
				weekDays = append(weekDays, "  ")
			} else {
				hasValidDay = true
				dayStr := fmt.Sprintf("%2d", dayNum)

				// Get todo count for this date
				dateStr := time.Date(cal.currentMonth.Year(), cal.currentMonth.Month(), dayNum, 0, 0, 0, 0, time.UTC)
				todoCount, _ := m.repository.GetTodoCountForDate(models.FormatDate(dateStr))

				// Style the day
				style := normalItemStyle
				if week == cal.cursorRow && day == cal.cursorCol && dayNum == cal.selectedDay {
					style = selectedItemStyle
					dayStr = ">" + dayStr[:1] + "<"
				} else if todoCount > 0 {
					style = accentStyle
					dayStr = dayStr + "•"
				} else {
					dayStr = dayStr + " "
				}

				// Today's date highlighting
				now := time.Now()
				if now.Year() == cal.currentMonth.Year() &&
					now.Month() == cal.currentMonth.Month() &&
					now.Day() == dayNum {
					style = style.Copy().Background(primaryColor).Foreground(lipgloss.Color("0"))
				}

				weekDays = append(weekDays, style.Render(dayStr))
			}
		}

		if hasValidDay {
			lines = append(lines, strings.Join(weekDays, " "))
		}
	}

	// Selected date info
	lines = append(lines, "")
	selectedDateStr := cal.getSelectedDate()
	selectedTodos, _ := m.repository.GetTodosForDate(selectedDateStr)

	if len(selectedTodos) > 0 {
		lines = append(lines, selectedItemStyle.Render(fmt.Sprintf("📝 %s (%d todos)", selectedDateStr, len(selectedTodos))))

		// Show first few todos
		for i, todo := range selectedTodos {
			if i >= 3 { // Show only first 3
				lines = append(lines, mutedStyle.Render("  ... and more"))
				break
			}
			checkbox := "☐"
			if todo.Completed {
				checkbox = "✓"
			}
			lines = append(lines, fmt.Sprintf("  %s %s", checkbox, todo.Title))
		}
	} else {
		lines = append(lines, mutedStyle.Render(fmt.Sprintf("📝 %s (no todos)", selectedDateStr)))
	}

	return strings.Join(lines, "\n")
}
