package models

import (
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type User struct {
	ID          uuid.UUID   `json:"id" db:"id"`
	Name        string      `json:"name" db:"name"`
	LastName    string      `json:"last_name" db:"last_name"`
	Email       string      `json:"email" db:"email"`
	Age         string      `json:"age" db:"age"`
	Calculators Calculators `has_many:"calculators"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}

// Users is not required by pop and may be deleted.
type Users []User

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name", Message: "Name cannot be empty"},
		&validators.StringIsPresent{Field: u.LastName, Name: "LastName", Message: "Last Name cannot be empty"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email", Message: "Email cannot be empty"},
		&validators.StringIsPresent{Field: u.Age, Name: "Age", Message: "Age cannot be empty"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name", Message: "Name cannot be empty"},
		&validators.StringIsPresent{Field: u.LastName, Name: "LastName", Message: "Last Name cannot be empty"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email", Message: "Email cannot be empty"},
		&validators.StringIsPresent{Field: u.Age, Name: "Age", Message: "Age cannot be empty"},
	), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
