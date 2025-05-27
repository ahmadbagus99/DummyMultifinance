package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
	"database/sql"
	"fmt"
)

type LimitUseCase interface {
	Insert(ctx context.Context, dbTx *sql.Tx, tx *models.Limits) (*models.Limits, error)
	UpdateLimit(ctx context.Context, dbTx *sql.Tx, consumer_id int, tenor int, amount float64) error
	GetLimitById(ctx context.Context, id int) (*models.Limits, error)
	CreateLimit(ctx context.Context, limit *models.Limits) (*models.Limits, error)
}

type limitUsecase struct {
	repo repositories.LimitRepository
	DB   *sql.Tx
}

func NewLimitUsecase(r repositories.LimitRepository, db *sql.Tx) LimitUseCase {
	return &limitUsecase{
		repo: r,
		DB:   db,
	}
}

func (uc *limitUsecase) Insert(ctx context.Context, dbTx *sql.Tx, limit *models.Limits) (*models.Limits, error) {
	return uc.repo.Insert(ctx, dbTx, limit)
}

func (uc *limitUsecase) UpdateLimit(ctx context.Context, dbTx *sql.Tx, consumer_id int, tenor int, amount float64) error {
	return uc.repo.UpdateLimit(ctx, dbTx, consumer_id, tenor, amount)
}

func (uc *limitUsecase) GetLimitById(ctx context.Context, id int) (*models.Limits, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *limitUsecase) CreateLimit(ctx context.Context, limit *models.Limits) (*models.Limits, error) {
	tx, err := uc.repo.GetDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	limitData, err := uc.Insert(ctx, tx, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to save transaction: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return limitData, nil
}
