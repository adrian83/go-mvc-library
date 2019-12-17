package book

import "github.com/adrian83/library/pkg/api"

// CreateBook represents data in request body.
type CreateBook struct {
	Title string `json:"title"`
}

func (cb *CreateBook) Validate() *api.ValidationError {
	violations := make([]*api.Violation, 0)

	if cb.Title == "" {
		violations = append(violations, api.NewViolation("title", "title cannot be empty"))
	}

	if len(violations) > 0 {
		return api.NewValidationError(violations...)
	}
	return nil
}

type UpdateBook struct {
	Title string `json:"title"`
}

func (cb *UpdateBook) Validate() *api.ValidationError {
	violations := make([]*api.Violation, 0)

	if cb.Title == "" {
		violations = append(violations, api.NewViolation("title", "title cannot be empty"))
	}

	if len(violations) > 0 {
		return api.NewValidationError(violations...)
	}
	return nil
}