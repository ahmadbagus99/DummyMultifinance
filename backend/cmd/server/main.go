package main

import (
	"DummyMultifinance/infrastructure/config"
	"DummyMultifinance/usecases"
	"fmt"
	"net/http"

	routes "DummyMultifinance/interfaces/routes"

	infrastructureConsumer "DummyMultifinance/infrastructure/repositories/consumers"
	infrastructureLimit "DummyMultifinance/infrastructure/repositories/limits"
	infrastructureTransaction "DummyMultifinance/infrastructure/repositories/transactions"
	infrastructureUser "DummyMultifinance/infrastructure/repositories/users"

	consumerUseCase "DummyMultifinance/usecases/consumers"
	limitUseCase "DummyMultifinance/usecases/limits"
	transactionUseCase "DummyMultifinance/usecases/transactions"
	userUseCase "DummyMultifinance/usecases/users"
)

func main() {
	config.LoadEnv()
	db := config.NewDB()

	userRepo := infrastructureUser.NewMysqlUserRepo(db)
	consumersRepo := infrastructureConsumer.NewMysqlConsumerRepo(db)
	transactionsRepo := infrastructureTransaction.NewMysqlTransactionRepo(db)
	limitsRepo := infrastructureLimit.NewMysqlLimitRepo(db)

	userUseCase := userUseCase.NewUserUsecase(userRepo)
	consumerUseCase := consumerUseCase.NewConsumerUsecase(consumersRepo)
	limitUseCase := limitUseCase.NewTransactionUsecase(limitsRepo)
	transactionUseCase := transactionUseCase.NewTransactionUsecase(transactionsRepo, consumersRepo, limitsRepo)

	useCases := &usecases.UseCases{
		UserUseCase:        userUseCase,
		ConsumerUseCase:    consumerUseCase,
		TransactionUseCase: transactionUseCase,
		LimitUseCase:       limitUseCase,
	}

	router := routes.NewRouter(useCases)
	port := config.GetEnv("APP_PORT", "8080")
	fmt.Println("Starting server on port:", port)

	http.ListenAndServe(":"+port, router)
}
