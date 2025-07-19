package models

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// generateID creates a unique ID for todos
func generateID() string {
	bytes := make([]byte, 4) // 8 char hex string
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// FormatDate formats a time.Time to YYYY-MM-DD string
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// ParseDate parses YYYY-MM-DD string to time.Time
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// TodayString return today's date a YYYY-MM-DD
func TodayString() string {
	return FormatDate(time.Now())
}

// GetDaysInMonth returns the number of days in a given month/year
func GetDaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// GetFirstDayOfMonth returns the first day of the month and its weekday
func GetFirstDayOfMonth(year int, month time.Month) (time.Time, time.Weekday) {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return firstDay, firstDay.Weekday()
}
