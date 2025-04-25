package main

import (
	"DummyMultifinance/infrastructure/config"
	infrastructureConsumer "DummyMultifinance/infrastructure/repositories/consumers"
	infrastructureTransaction "DummyMultifinance/infrastructure/repositories/transactions"
	infrastructureUser "DummyMultifinance/infrastructure/repositories/users"
	routes "DummyMultifinance/interfaces/routes"
	"DummyMultifinance/usecases"
	consumerUseCase "DummyMultifinance/usecases/consumers"
	transactionUseCase "DummyMultifinance/usecases/transactions"
	userUseCase "DummyMultifinance/usecases/users"
	"fmt"
	"net/http"
)

func main() {
	config.LoadEnv()
	db := config.NewDB()

	userRepo := infrastructureUser.NewMysqlUserRepo(db)
	consumersRepo := infrastructureConsumer.NewMysqlConsumerRepo(db)
	transactionsRepo := infrastructureTransaction.NewMysqlTransactionRepo(db)

	userUseCase := userUseCase.NewUserUsecase(userRepo)
	consumerUseCase := consumerUseCase.NewConsumerUsecase(consumersRepo)
	transactionUseCase := transactionUseCase.NewTransactionUsecase(transactionsRepo)

	useCases := &usecases.UseCases{
		UserUseCase:        userUseCase,
		TransactionUseCase: transactionUseCase,
		ConsumerUseCase:    consumerUseCase,
	}

	router := routes.NewRouter(useCases)
	port := config.GetEnv("APP_PORT", "8080")
	fmt.Println("Starting server on port:", port)

	http.ListenAndServe(":"+port, router)
}
