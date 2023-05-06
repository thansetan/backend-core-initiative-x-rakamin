package repository

import (
	"context"
	"fmt"
	"self-payroll/config"
	"self-payroll/model"
	"time"
)

type userRepository struct {
	Cfg config.Config
}

func NewUserRepository(cfg config.Config) model.UserRepository {
	return &userRepository{Cfg: cfg}
}

func (p *userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {

	user := new(model.User)
	if err := p.Cfg.Database().
		WithContext(ctx).
		Preload("Position").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := p.Cfg.Database().
		WithContext(ctx).
		Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func (p *userRepository) UpdateByID(ctx context.Context, id int, user *model.User) (*model.User, error) {
	if err := p.Cfg.Database().
		WithContext(ctx).
		Model(&model.User{ID: id}).
		Updates(user).
		Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func (p *userRepository) Delete(ctx context.Context, id int) error {

	if _, err := p.FindByID(ctx, id); err != nil {
		return err
	}

	if err := p.Cfg.Database().WithContext(ctx).
		Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (p *userRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.User, error) {
	var data []*model.User

	if err := p.Cfg.Database().WithContext(ctx).Preload("Position").
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (p *userRepository) CheckLastWithdraw(ctx context.Context, name_id string) error {
	t := new(model.Transaction)
	if err := p.Cfg.Database().
		WithContext(ctx).
		Where("note LIKE ? and created_at >= ?",
			fmt.Sprintf("%s%%", name_id),
			time.Now().AddDate(0, -1, 0)).
		Last(&t).Error; err != nil {
		return nil
	}
	err := fmt.Errorf("you have withdrawn your salary for this month at %s %d, %d", t.CreatedAt.Month().String(), t.CreatedAt.Day(), t.CreatedAt.Year())
	return err
}
