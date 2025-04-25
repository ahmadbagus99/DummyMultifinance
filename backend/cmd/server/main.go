package main

import (
	"DummyMultifinance/infrastructure/config" // Pastikan ini mengarah ke router yang benar
	infrastructureConsumer "DummyMultifinance/infrastructure/repositories/consumers"
	infrastructureTransaction "DummyMultifinance/infrastructure/repositories/transactions"
	routes "DummyMultifinance/interfaces/routes"
	"DummyMultifinance/usecases"
	consumerUseCase "DummyMultifinance/usecases/consumers"
	transactionUseCase "DummyMultifinance/usecases/transactions"
	"fmt"
	"net/http"
)

func main() {
	config.LoadEnv()
	db := config.NewDB()

	transactionsRepo := infrastructureTransaction.NewMysqlTransactionRepo(db)
	consumersRepo := infrastructureConsumer.NewMysqlConsumerRepo(db)

	transactionUseCase := transactionUseCase.NewTransactionUsecase(transactionsRepo)
	consumerUseCase := consumerUseCase.NewConsumerUsecase(consumersRepo)

	useCases := &usecases.UseCases{
		// UserUseCase:        userUseCase,
		TransactionUseCase: transactionUseCase,
		ConsumerUseCase:    consumerUseCase,
	}

	router := routes.NewRouter(useCases)
	port := config.GetEnv("APP_PORT", "8080")
	fmt.Println("Starting server on port:", port)

	http.ListenAndServe(":"+port, router)
}
