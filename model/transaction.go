package model

import (
	"context"
	"time"
)

const (
	TransactionTypeDebit   = "debit"
	TransactionsTypeCredit = "credit"
)

type (
	Transaction struct {
		ID        int       `json:"id"`
		Amount    int       `json:"amount"`
		Note      string    `json:"note"`
		Type      string    `json:"type"`
		UserID    int       `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	TransactionRepository interface {
		Fetch(ctx context.Context, limit, offset int) ([]*Transaction, error)
		FetchByUserID(ctx context.Context, userID, limit, offset int) ([]*Transaction, error)
		CheckLastWithdrawal(ctx context.Context, userID int) error
	}

	TransactionUsecase interface {
		Fetch(ctx context.Context, limit, offset int) ([]*Transaction, int, error)
		FetchByUserID(ctx context.Context, userID, limit, offset int) ([]*Transaction, int, error)
	}
)
