package model

import (
	"context"
	"self-payroll/request"
	"self-payroll/response"
	"time"
)

type (
	User struct {
		ID           int            `json:"id"`
		SecretID     string         `json:"secret_id"`
		Name         string         `json:"name"`
		Email        string         `json:"email"`
		Password     string         `json:"password"`
		Phone        string         `json:"phone"`
		Address      string         `json:"address"`
		PositionID   int            `json:"position_id"`
		IsAdmin      bool           `gorm:"default:false"`
		Position     *Position      `json:"position"`
		CreatedAt    time.Time      `json:"created_at"`
		UpdatedAt    time.Time      `json:"updated_at"`
		Transactions []*Transaction `gorm:"foreignKey:UserID"`
	}

	UserRepository interface {
		Create(ctx context.Context, user *User) (*User, error)
		UpdateByID(ctx context.Context, id int, user *User) (*User, error)
		FindByID(ctx context.Context, id int) (*User, error)
		Delete(ctx context.Context, id int) error
		Fetch(ctx context.Context, limit, offset int) ([]*User, error)
		FindByEmail(ctx context.Context, email string) (*User, error)
	}

	UserUsecase interface {
		GetByID(ctx context.Context, id int) (*User, error)
		FetchUser(ctx context.Context, limit, offset int) ([]*User, error)
		DestroyUser(ctx context.Context, id int) error
		EditUser(ctx context.Context, id int, req *request.UpdateRequest) (*User, error)
		StoreUser(ctx context.Context, req *request.RegisterRequest) (*User, error)
		WithdrawSalary(ctx context.Context, userID int, req *request.WithdrawRequest) error
		Login(ctx context.Context, req *request.LoginRequest) (*response.LoginResponse, error)
		AdminRegister(ctx context.Context, req *request.AdminRequest) (*User, error)
	}
)
