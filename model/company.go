package model

import (
	"context"
	"self-payroll/request"
	"time"
)

type (
	Company struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Address   string    `json:"address"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CompanyRepository interface {
		Get(ctx context.Context) (*Company, error)
		CreateOrUpdate(ctx context.Context, Company *Company) (*Company, error)
		AddBalance(ctx context.Context, userID, balance int) (*Company, error)
		DebitBalance(ctx context.Context, amount, userID int, note string) error
	}

	CompanyUsecase interface {
		GetCompanyInfo(ctx context.Context) (*Company, int, error)
		CreateOrUpdateCompany(ctx context.Context, req request.CompanyRequest) (*Company, int, error)
		TopupBalance(ctx context.Context, userID int, req request.TopupCompanyBalance) (*Company, int, error)
	}
)
