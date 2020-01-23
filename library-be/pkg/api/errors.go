package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	msgInternalError = "internal server error"
)

// Violation contains information about which structure's field is invalida and why.
type Violation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewViolation is a constructor for Violation struct.
func NewViolation(field, msg string) *Violation {
	return &Violation{
		Field:   field,
		Message: msg,
	}
}

type ValidationError struct {
	Violations []*Violation
}

func (e *ValidationError) Error() string {
	bts, _ := json.Marshal(e.Violations)
	return string(bts)
}

func NewValidationError(violations ...*Violation) *ValidationError {
	return &ValidationError{
		Violations: violations,
	}
}

func HandleError(err error, w http.ResponseWriter) {
	fmt.Printf("Error in API: %v", err)

	switch t := err.(type) {
	case *ValidationError:
		fmt.Printf("Violations: %v", t.Violations)
		ResponseJSON(http.StatusBadRequest, t.Violations, w)
	default:
		ResponseText(http.StatusInternalServerError, msgInternalError, w)
	}
}
