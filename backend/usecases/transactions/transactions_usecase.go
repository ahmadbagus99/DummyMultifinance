package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
	"fmt"
	"time"
)

type TransactionUseCase interface {
	CreateTransaction(ctx context.Context, tx *models.Transactions) (*models.Transactions, error)
	RequestTransaction(ctx context.Context, consumer_id int, asset_name string, tenor int, amount float64) (*models.Transactions, error)
	GetTransactionById(ctx context.Context, id int) (*models.Transactions, error)
}

type transactionUsecase struct {
	repo         repositories.TransactionRepository
	repoConsumer repositories.ConsumerRepository
	repoLimit    repositories.LimitRepository
}

func NewTransactionUsecase(r repositories.TransactionRepository, repoConsumer repositories.ConsumerRepository, repoLimit repositories.LimitRepository) TransactionUseCase {
	return &transactionUsecase{
		repo:         r,
		repoConsumer: repoConsumer,
		repoLimit:    repoLimit,
	}
}

func (uc *transactionUsecase) GetTransactionById(ctx context.Context, id int) (*models.Transactions, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *transactionUsecase) CreateTransaction(ctx context.Context, tx *models.Transactions) (*models.Transactions, error) {
	tx.TransactionDate = time.Now()
	return uc.repo.Insert(ctx, tx)
}

func (uc *transactionUsecase) RequestTransaction(ctx context.Context, consumer_id int, asset_name string, tenor int, amount float64) (*models.Transactions, error) {
	listConsumerLimit, err := uc.repoConsumer.GetConsumerLimit(ctx, consumer_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get consumer limits: %v", err)
	}

	var limit float64
	var interestRate float64

	for _, consumerLimit := range listConsumerLimit {
		if consumerLimit.TenorID == tenor {
			limit = consumerLimit.Amount
			break
		}
	}

	if limit == 0 {
		return nil, fmt.Errorf("no limit found for the selected tenor")
	}

	switch tenor {
	case 1:
		interestRate = 0.05
	case 2:
		interestRate = 0.06
	case 3:
		interestRate = 0.07
	case 6:
		interestRate = 0.08
	default:
		return nil, fmt.Errorf("invalid tenor selected")
	}

	if amount > limit {
		return nil, fmt.Errorf("insufficient limit for the selected tenor")
	}

	interest := amount * interestRate
	totalAmount := amount + interest
	cicilan := totalAmount / float64(tenor)

	loanRequestTransaction := &models.Transactions{
		ConsumerID:      consumer_id,
		OTR:             amount,
		AdminFee:        2000,
		Installment:     cicilan,
		Interest:        interest,
		AssetName:       asset_name,
		Approved:        true,
		TransactionDate: time.Now(),
	}

	createdTx, err := uc.repo.Insert(ctx, loanRequestTransaction)
	if err != nil {
		return nil, fmt.Errorf("failed to save transaction: %v", err)
	}

	err = uc.repoLimit.UpdateLimit(ctx, consumer_id, tenor, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to update consumer limit: %v", err)
	}

	return createdTx, nil
}
