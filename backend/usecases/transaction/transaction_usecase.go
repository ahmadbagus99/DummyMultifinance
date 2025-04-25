package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
	"time"
)

type TransactionUseCase interface {
	CreateTransaction(ctx context.Context, tx *models.Transaction) (*models.Transaction, error)
	GetTransactionById(ctx context.Context, id int) (*models.Transaction, error)
}

type transactionUsecase struct {
	repo repositories.TransactionRepository
}

func NewTransactionUsecase(r repositories.TransactionRepository) TransactionUseCase {
	return &transactionUsecase{repo: r}
}

func (uc *transactionUsecase) GetTransactionById(ctx context.Context, id int) (*models.Transaction, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *transactionUsecase) CreateTransaction(ctx context.Context, tx *models.Transaction) (*models.Transaction, error) {
	tx.TransactionDate = time.Now()
	return uc.repo.Insert(ctx, tx)
}
