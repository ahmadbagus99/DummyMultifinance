package main

import (
	"DummyMultifinance/infrastructure/config" // Pastikan ini mengarah ke router yang benar
	infrastructure "DummyMultifinance/infrastructure/repositories/transaction"
	routes "DummyMultifinance/interfaces/routes"
	"DummyMultifinance/usecases"
	transactionUseCase "DummyMultifinance/usecases/transaction"
	"fmt"
	"net/http"
)

func main() {
	config.LoadEnv()
	db := config.NewDB()

	txRepo := infrastructure.NewMysqlTransactionRepo(db)
	txUseCase := transactionUseCase.NewTransactionUsecase(txRepo)

	useCases := &usecases.UseCases{
		// UserUseCase:        userUseCase,
		TransactionUseCase: txUseCase,
	}

	router := routes.NewRouter(useCases)
	port := config.GetEnv("APP_PORT", "8080")
	fmt.Println("Starting server on port:", port)

	http.ListenAndServe(":"+port, router)
}
