package repository

import (
	"context"
	"fmt"
	"self-payroll/config"
	"self-payroll/model"
	"time"
)

type transactionRepository struct {
	Cfg config.Config
}

func NewTransactionRepository(cfg config.Config) model.TransactionRepository {
	return &transactionRepository{Cfg: cfg}
}

func (t *transactionRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Transaction, error) {
	var data []*model.Transaction

	if err := t.Cfg.Database().WithContext(ctx).
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (t *transactionRepository) FetchByUserID(ctx context.Context, userID, limit, offset int) ([]*model.Transaction, error) {
	var data []*model.Transaction

	if err := t.Cfg.Database().WithContext(ctx).
		Limit(limit).Offset(offset).Where("user_id = ?", userID).
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (t *transactionRepository) CheckLastWithdrawal(ctx context.Context, userID int) error {
	trx := new(model.Transaction)
	fmt.Println("masuk sini")
	if err := t.Cfg.Database().WithContext(ctx).
		Where("user_id = ? and created_at >= ?", userID, time.Now().AddDate(0, -1, 0)).
		Last(&trx).Error; err != nil {
		fmt.Println("gaada")
		return nil
	}
	err := fmt.Errorf("you have withdrawn your salary for this month at %s %d, %d", trx.CreatedAt.Month().String(), trx.CreatedAt.Day(), trx.CreatedAt.Year())
	return err
}
