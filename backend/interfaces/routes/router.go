package interfaces

import (
	handler "DummyMultifinance/interfaces/handlers"
	"DummyMultifinance/usecases"
	"net/http"
)

func NewRouter(useCases *usecases.UseCases) *http.ServeMux {
	router := http.NewServeMux()

	transactionHandler := handler.NewTransactionHandler(useCases.TransactionUseCase)
	consumerHandler := handler.NewConsumerHandler(useCases.ConsumerUseCase)

	// Menambahkan routing
	// router.HandleFunc("/register", userHandler.RegisterUser)
	// router.HandleFunc("/login", userHandler.Login)
	router.HandleFunc("/insert-transactions", transactionHandler.CreateTransaction)
	router.HandleFunc("/get-transaction", transactionHandler.GetTransaction)

	router.HandleFunc("/insert-consumer", consumerHandler.CreateConsumer)
	router.HandleFunc("/get-consumer", consumerHandler.GetConsumer)

	return router
}
