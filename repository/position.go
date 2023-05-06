package repository

import (
	"context"
	"fmt"
	"self-payroll/config"
	"self-payroll/model"
)

type positionRepository struct {
	Cfg config.Config
}

func NewPositionRepository(cfg config.Config) model.PositionRepository {
	return &positionRepository{Cfg: cfg}
}

func (p *positionRepository) FindByID(ctx context.Context, id int) (*model.Position, error) {
	position := new(model.Position)

	if err := p.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		First(position).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return position, nil
}

func (p *positionRepository) Create(ctx context.Context, position *model.Position) (*model.Position, error) {
	if err := p.Cfg.Database().WithContext(ctx).Create(&position).Error; err != nil {
		return nil, err
	}
	return position, nil
}

func (p *positionRepository) UpdateByID(ctx context.Context, id int, position *model.Position) (*model.Position, error) {
	if err := p.Cfg.Database().WithContext(ctx).
		Model(&model.Position{ID: id}).Updates(position).Find(position).Error; err != nil {
		return nil, err
	}
	return position, nil
}

func (p *positionRepository) Delete(ctx context.Context, id int) error {
	position := new(model.Position)

	if _, err := p.FindByID(ctx, id); err != nil {
		return err
	}
	if err := p.Cfg.Database().
		WithContext(ctx).
		Delete(position, id).Error; err != nil {
		return err
	}
	return nil

}

func (p *positionRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Position, error) {
	var positions []*model.Position
	if err := p.Cfg.Database().WithContext(ctx).Limit(limit).Offset(offset).Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}
