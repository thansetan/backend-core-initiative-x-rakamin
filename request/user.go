package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (
	RegisterRequest struct {
		SecretID   string `json:"secret_id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Phone      string `json:"phone"`
		Address    string `json:"address"`
		PositionID int    `json:"position_id"`
	}

	UpdateRequest struct {
		SecretID string `json:"secret_id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
	}

	WithdrawRequest struct {
		SecretID string `json:"secret_id"`
	}

	AdminRequest struct {
		SecretID   string `json:"secret_id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Phone      string `json:"phone"`
		Address    string `json:"address"`
		PositionID int
	}
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (req WithdrawRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.SecretID, validation.Required, is.Alphanumeric),
	)
}

func (req RegisterRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.SecretID, validation.Required),
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 32)),
		validation.Field(&req.Phone, validation.Required, is.UTFNumeric, validation.Length(10, 13)),
		validation.Field(&req.Address, validation.Required),
		validation.Field(&req.PositionID, validation.Required),
	)
}

func (req LoginRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required),
	)
}

func (req AdminRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.SecretID, validation.Required),
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 32)),
		validation.Field(&req.Phone, validation.Required, is.UTFNumeric, validation.Length(10, 13)),
		validation.Field(&req.Address, validation.Required),
	)
}

func (req UpdateRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.SecretID),
		validation.Field(&req.Name),
		validation.Field(&req.Email, is.Email),
		validation.Field(&req.Password, validation.Length(8, 32)),
		validation.Field(&req.Phone, is.UTFNumeric, validation.Length(10, 13)),
		validation.Field(&req.Address),
	)
}
