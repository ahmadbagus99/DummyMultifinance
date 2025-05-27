package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"database/sql"

	"context"
	"fmt"
	"time"
)

type TransactionUseCase interface {
	CreateTransaction(ctx context.Context, dbTx *sql.Tx, tx *models.Transactions) (*models.Transactions, error)
	RequestTransaction(ctx context.Context, consumer_id int, asset_name string, tenor int, amount float64) (*models.Transactions, error)
	GetTransactionById(ctx context.Context, id int) (*models.Transactions, error)
}

type transactionUsecase struct {
	repo         repositories.TransactionRepository
	repoConsumer repositories.ConsumerRepository
	repoLimit    repositories.LimitRepository
	DB           *sql.Tx
}

func NewTransactionUsecase(repoTransaction repositories.TransactionRepository, repoConsumer repositories.ConsumerRepository, repoLimit repositories.LimitRepository, db *sql.Tx) TransactionUseCase {
	return &transactionUsecase{
		repo:         repoTransaction,
		repoConsumer: repoConsumer,
		repoLimit:    repoLimit,
		DB:           db,
	}
}

func (uc *transactionUsecase) GetTransactionById(ctx context.Context, id int) (*models.Transactions, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *transactionUsecase) CreateTransaction(ctx context.Context, dbTx *sql.Tx, transaction *models.Transactions) (*models.Transactions, error) {
	transaction.TransactionDate = time.Now()
	return uc.repo.Insert(ctx, dbTx, transaction)
}

func (uc *transactionUsecase) RequestTransaction(ctx context.Context, consumer_id int, asset_name string, tenor int, amount float64) (*models.Transactions, error) {
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

	createdTx, err := uc.CreateTransaction(ctx, tx, loanRequestTransaction)
	if err != nil {
		return nil, fmt.Errorf("failed to save transaction: %v", err)
	}

	err = uc.repoLimit.UpdateLimit(ctx, tx, consumer_id, tenor, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to update consumer limit: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return createdTx, nil
}
