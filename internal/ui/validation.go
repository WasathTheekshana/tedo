package ui

import (
	"fmt"
	"strings"
	"unicode"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateTodoInput validates todo input fields
func ValidateTodoInput(title, description string) []ValidationError {
	var errors []ValidationError

	// Validate title
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		errors = append(errors, ValidationError{
			Field:   "Title",
			Message: "cannot be empty",
		})
	} else if len(title) > 100 {
		errors = append(errors, ValidationError{
			Field:   "Title",
			Message: "cannot exceed 100 characters",
		})
	} else if !isValidText(title) {
		errors = append(errors, ValidationError{
			Field:   "Title",
			Message: "contains invalid characters",
		})
	}

	// Validate description (optional but has limits)
	description = strings.TrimSpace(description)
	if len(description) > 500 {
		errors = append(errors, ValidationError{
			Field:   "Description",
			Message: "cannot exceed 500 characters",
		})
	} else if description != "" && !isValidText(description) {
		errors = append(errors, ValidationError{
			Field:   "Description",
			Message: "contains invalid characters",
		})
	}

	return errors
}

// isValidText checks if text contains only valid characters
func isValidText(text string) bool {
	for _, r := range text {
		// Allow printable characters, spaces, and common punctuation
		if !unicode.IsPrint(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// CleanInput cleans and trims input text
func CleanInput(input string) string {
	// Trim whitespace
	cleaned := strings.TrimSpace(input)

	// Replace multiple spaces with single space
	cleaned = strings.Join(strings.Fields(cleaned), " ")

	return cleaned
}
