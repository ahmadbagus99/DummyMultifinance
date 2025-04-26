package usecases

import (
	consumerUseCase "DummyMultifinance/usecases/consumers"
	limitUseCase "DummyMultifinance/usecases/limits"
	transactionUseCase "DummyMultifinance/usecases/transactions"
	userUseCase "DummyMultifinance/usecases/users"
)

type UseCases struct {
	UserUseCase        userUseCase.UserUseCase
	TransactionUseCase transactionUseCase.TransactionUseCase
	ConsumerUseCase    consumerUseCase.ConsumerUseCase
	LimitUseCase       limitUseCase.LimiUseCase
}

func NewUseCases(userUC userUseCase.UserUseCase, transactionUC transactionUseCase.TransactionUseCase, consumerUC consumerUseCase.ConsumerUseCase, limitUC limitUseCase.LimiUseCase) *UseCases {
	return &UseCases{
		UserUseCase:        userUC,
		TransactionUseCase: transactionUC,
		ConsumerUseCase:    consumerUC,
		LimitUseCase:       limitUC,
	}
}
