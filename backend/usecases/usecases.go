package usecases

import (
	consumerUseCase "DummyMultifinance/usecases/consumers"
	transactionUseCase "DummyMultifinance/usecases/transactions"
	userUseCase "DummyMultifinance/usecases/users"
)

type UseCases struct {
	UserUseCase        userUseCase.UserUseCase
	TransactionUseCase transactionUseCase.TransactionUseCase
	ConsumerUseCase    consumerUseCase.ConsumerUseCase
}

func NewUseCases(userUC userUseCase.UserUseCase, transactionUC transactionUseCase.TransactionUseCase, consumerUC consumerUseCase.ConsumerUseCase) *UseCases {
	return &UseCases{
		// UserUseCase:        userUC,
		TransactionUseCase: transactionUC,
		ConsumerUseCase:    consumerUC,
	}
}
