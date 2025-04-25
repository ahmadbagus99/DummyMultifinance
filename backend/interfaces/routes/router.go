package interfaces

import (
	handler "DummyMultifinance/interfaces/handlers"
	"DummyMultifinance/usecases"
	"net/http"
)

// // Struct UseCases untuk menyatukan use case terkait
// type UseCases struct {
// 	// UserUseCase        *usecases.UserUseCase
// 	TransactionUseCase usecases.TransactionUseCase // Pastikan penamaan sesuai field
// }

// NewRouter hanya mengatur routing
func NewRouter(useCases *usecases.UseCases) *http.ServeMux {
	router := http.NewServeMux()

	// Menyiapkan handler dengan use case yang telah di-inject
	// userHandler := handler.NewUserHandler(useCases.UserUseCase)                      // hanya UserUseCase
	transactionHandler := handler.NewTransactionHandler(useCases.TransactionUseCase) // Pastikan nama field benar

	// Menambahkan routing
	// router.HandleFunc("/register", userHandler.RegisterUser)
	// router.HandleFunc("/login", userHandler.Login)
	router.HandleFunc("/insert-transactions", transactionHandler.CreateTransaction) // untuk transaksi
	router.HandleFunc("/get-transaction", transactionHandler.GetTransaction)        // untuk transaksi

	return router
}
