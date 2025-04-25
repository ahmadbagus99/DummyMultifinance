package interfaces

import (
	handler "DummyMultifinance/interfaces/handlers"
	"DummyMultifinance/usecases"
	"net/http"
)

func NewRouter(useCases *usecases.UseCases) *http.ServeMux {
	router := http.NewServeMux()

	userHandler := handler.NewUserHandler(useCases.UserUseCase)
	consumerHandler := handler.NewConsumerHandler(useCases.ConsumerUseCase)
	transactionHandler := handler.NewTransactionHandler(useCases.TransactionUseCase)

	router.HandleFunc("/register", userHandler.CreateUser)
	router.HandleFunc("/login", userHandler.Login)

	router.HandleFunc("/insert-transactions", transactionHandler.CreateTransaction)
	router.HandleFunc("/get-transaction", transactionHandler.GetTransaction)

	router.HandleFunc("/insert-consumer", consumerHandler.CreateConsumer)
	router.HandleFunc("/get-consumer", consumerHandler.GetConsumer)

	return router
}
