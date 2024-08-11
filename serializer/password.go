package serializer

import (
	"regexp"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

// PasswordRequest represents the request body for password validation
type PasswordRequest struct {
	InitPassword string `json:"init_password"`
}

// Validate checks if the PasswordRequest has valid fields
func (d *PasswordRequest) Validate() error {
	return v.ValidateStruct(
		d,
		v.Field(&d.InitPassword,
			v.Required,
			v.Length(1, 40), // Length between 1 and 40 characters
			v.Match(regexp.MustCompile(`^[a-zA-Z0-9.!]+$`)), // Only letters, digits, dots, and exclamation marks
		),
	)
}

type PasswordResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}
