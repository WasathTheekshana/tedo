package ui

import (
	"fmt"
	"strings"
	"time"
)

// ErrorState manages error display
type ErrorState struct {
	message   string
	timestamp time.Time
	isVisible bool
}

// SetError sets an error message
func (e *ErrorState) SetError(err error) {
	if err != nil {
		e.message = err.Error()
		e.timestamp = time.Now()
		e.isVisible = true
	}
}

// SetErrorMessage sets a custom error message
func (e *ErrorState) SetErrorMessage(msg string) {
	e.message = msg
	e.timestamp = time.Now()
	e.isVisible = true
}

// ClearError clears the current error
func (e *ErrorState) ClearError() {
	e.isVisible = false
	e.message = ""
}

// GetError returns the current error message if visible
func (e *ErrorState) GetError() string {
	if !e.isVisible {
		return ""
	}

	// Auto-clear errors after 5 seconds
	if time.Since(e.timestamp) > 5*time.Second {
		e.ClearError()
		return ""
	}

	return e.message
}

// FormatValidationErrors formats validation errors nicely
func FormatValidationErrors(errors []ValidationError) string {
	if len(errors) == 0 {
		return ""
	}

	var messages []string
	for _, err := range errors {
		messages = append(messages, err.Error())
	}

	return fmt.Sprintf("Validation errors: %s", strings.Join(messages, "; "))
}
