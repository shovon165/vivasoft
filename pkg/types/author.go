package types

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ReadAuthorResponse struct {
	ID          uint   `json:"id"`
	AuthorName  string `json:"authorName"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

type CreateAuthorRequest struct {
	ReadAuthorResponse
	ID uint `json:"-"`
}

type UpdateAuthorRequest struct {
	ReadAuthorResponse
	ID uint `json:"-"`
}

// Validates the CreateAuthorRequest request
func (request CreateAuthorRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.AuthorName,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(2, 64)),
		validation.Field(&request.Address,
			validation.Length(2, 128)),
		validation.Field(&request.PhoneNumber,
			validation.Length(8, 32)))

}

// Validates the UpdateAuthorRequest request
func (request UpdateAuthorRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.AuthorName,
			validation.Length(2, 64)),
		validation.Field(&request.Address,
			validation.Length(2, 128)),
		validation.Field(&request.PhoneNumber,
			validation.Length(8, 32)))
}
