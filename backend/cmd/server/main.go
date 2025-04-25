package main

import (
	"DummyMultifinance/infrastructure"
	"DummyMultifinance/infrastructure/config"
	interfaces "DummyMultifinance/interfaces/handlers"
	"DummyMultifinance/usecases"
	"fmt"
	"net/http"
)

func main() {
	config.LoadEnv()

	// Inisialisasi dependency
	userRepo := infrastructure.NewInMemoryUserRepo()
	userUseCase := &usecases.UserUseCase{
		UserRepo: userRepo,
	}
	userHandler := &interfaces.UserHandler{
		UserUseCase: userUseCase,
	}

	// Routing
	http.HandleFunc("/register", userHandler.RegisterUser)
	http.HandleFunc("/login", userHandler.Login)

	// Menjalankan server
	port := config.GetEnv("APP_PORT", "8080")
	fmt.Println("Starting server on port:", port)

	http.ListenAndServe(":"+port, nil)
}
