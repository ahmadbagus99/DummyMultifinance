package interfaces

import (
	handlerConsumer "DummyMultifinance/interfaces/handlers/consumers"
	handlerLimit "DummyMultifinance/interfaces/handlers/limits"
	handlerTransaction "DummyMultifinance/interfaces/handlers/transactions"
	handlerUser "DummyMultifinance/interfaces/handlers/users"
	"DummyMultifinance/usecases"
	"net/http"
)

func NewRouter(useCases *usecases.UseCases) *http.ServeMux {
	router := http.NewServeMux()

	userHandler := handlerUser.NewUserHandler(useCases.UserUseCase)
	consumerHandler := handlerConsumer.NewConsumerHandler(useCases.ConsumerUseCase)
	transactionHandler := handlerTransaction.NewTransactionHandler(useCases.TransactionUseCase)
	limitHandler := handlerLimit.NewLimitHandler(useCases.LimitUseCase)

	router.HandleFunc("/register", userHandler.CreateUser)
	router.HandleFunc("/login", userHandler.Login)

	router.HandleFunc("/insert-consumer", consumerHandler.CreateConsumer)
	router.HandleFunc("/get-consumer", consumerHandler.GetConsumer)
	router.HandleFunc("/get-consumer-limit", consumerHandler.GetConsumerLimit)

	router.HandleFunc("/insert-transaction", transactionHandler.CreateTransaction)
	router.HandleFunc("/get-transaction", transactionHandler.GetTransaction)

	router.HandleFunc("/insert-limits", limitHandler.CreateLimit)
	router.HandleFunc("/get-limit", limitHandler.GetLimit)

	return router
}
