package interfaces

import (
	handlerConsumer "DummyMultifinance/interfaces/handlers/consumers"
	handlerLimit "DummyMultifinance/interfaces/handlers/limits"
	handlerTransaction "DummyMultifinance/interfaces/handlers/transactions"
	handlerUser "DummyMultifinance/interfaces/handlers/users"
	"DummyMultifinance/interfaces/middlewares"
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

	router.HandleFunc("/insert-consumer", middlewares.TokenValidation(consumerHandler.CreateConsumer))
	router.HandleFunc("/get-consumer", middlewares.TokenValidation(consumerHandler.GetConsumer))
	router.HandleFunc("/get-consumer-limit", middlewares.TokenValidation(consumerHandler.GetConsumerLimit))

	router.HandleFunc("/insert-transaction", middlewares.TokenValidation(transactionHandler.CreateTransaction))
	router.HandleFunc("/get-transaction", middlewares.TokenValidation(transactionHandler.GetTransaction))

	router.HandleFunc("/insert-limits", middlewares.TokenValidation(limitHandler.InsertLimit))
	router.HandleFunc("/get-limit", middlewares.TokenValidation(limitHandler.GetLimit))

	return router
}
