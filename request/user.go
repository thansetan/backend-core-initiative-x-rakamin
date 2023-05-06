package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (
	UserRequest struct {
		ID         int    `json:"id"`
		SecretID   string `json:"secret_id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		Address    string `json:"address"`
		PositionID int    `json:"position_id"`
	}

	WithdrawRequest struct {
		ID       int    `json:"id"`
		SecretID string `json:"secret_id"`
	}
)

func (req WithdrawRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required),
		validation.Field(&req.SecretID, validation.Required, is.Alphanumeric),
	)
}

func (req UserRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.SecretID, validation.Required),
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Phone, validation.Required, is.UTFNumeric, validation.Length(10, 13)),
		validation.Field(&req.Address, validation.Required),
		validation.Field(&req.PositionID, validation.Required),
	)
}
