package usecases

import (
	TransactionUseCase "DummyMultifinance/usecases/transaction"
	userUseCase "DummyMultifinance/usecases/user"
)

type UseCases struct {
	UserUseCase        userUseCase.UserUseCase
	TransactionUseCase TransactionUseCase.TransactionUseCase
}

func NewUseCases(userUC userUseCase.UserUseCase, txUC TransactionUseCase.TransactionUseCase) *UseCases {
	return &UseCases{
		// UserUseCase:        userUC,
		TransactionUseCase: txUC,
	}
}
