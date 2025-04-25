package interfaces

import (
	interfaces "DummyMultifinance/interfaces/handlers"
	"DummyMultifinance/usecases"
	"net/http"
)

func NewRouter(userUseCase *usecases.UserUseCase) *http.ServeMux {
	router := http.NewServeMux()

	// Menyiapkan handler dengan use case yang telah diinject
	userHandler := interfaces.NewUserHandler(userUseCase)

	// Menambahkan routing
	router.HandleFunc("/register", userHandler.RegisterUser)
	router.HandleFunc("/login", userHandler.Login)

	return router
}
