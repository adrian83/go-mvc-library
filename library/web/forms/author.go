package forms

import (
	"fmt"

	"github.com/adrian83/go-mvc-library/library/domain/author"
	"github.com/adrian83/go-mvc-library/library/domain/common"
)

var (
	emptyAuthorLastName = common.ValidationError{
		Field:   "LastName",
		Code:    "author.lastName",
		Message: "Author's last name cannot be empty"}
)

// AuthorForm represents form for creating / updating authors.
type AuthorForm struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

// String returns struct's string representation.
func (a *AuthorForm) String() string {
	return fmt.Sprintf("AuthorForm {ID: %v, FirstName: %v, LastName: %v}", a.ID, a.FirstName, a.LastName)
}

// Validate implements Validable interface.
func (a *AuthorForm) Validate() common.ValidationErrors {
	errs := make([]*common.ValidationError, 0)

	if common.IsStringEmpty(a.LastName) {
		errs = append(errs, &emptyAuthorLastName)
	}

	return errs
}

// ToAuthor transforms AuthorForm to Author struct.
func (a AuthorForm) ToAuthor() *author.Author {
	return &author.Author{
		ID:        a.ID,
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}
