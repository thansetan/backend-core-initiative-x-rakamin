package usecase

import (
	"context"
	"net/http"
	"self-payroll/model"
)

type transactionUsecase struct {
	transactionRepository model.TransactionRepository
}

func NewTransactionUsecase(transaction model.TransactionRepository) model.TransactionUsecase {
	return &transactionUsecase{transactionRepository: transaction}
}

func (t *transactionUsecase) Fetch(ctx context.Context, limit, offset int) ([]*model.Transaction, int, error) {
	transactions, err := t.transactionRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return transactions, http.StatusOK, nil

}

func (t *transactionUsecase) FetchByUserID(ctx context.Context, userID, limit, offset int) ([]*model.Transaction, int, error) {
	transactions, err := t.transactionRepository.FetchByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return transactions, http.StatusOK, nil
}
