package interfaces

import (
	"DummyMultifinance/domain/models"
	userUseCase "DummyMultifinance/usecases/users"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserUseCase userUseCase.UserUseCase
}

func NewUserHandler(uc userUseCase.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: uc,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		GeneralResponse(w, http.StatusMethodNotAllowed, "BadRequest", "Invalid method", nil)
		return
	}

	var tx models.Users
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		GeneralResponse(w, http.StatusBadRequest, BadRequest, err.Error(), nil)
		return
	}

	createdTx, err := h.UserUseCase.CreateUser(r.Context(), &tx)
	if err != nil {
		GeneralResponse(w, http.StatusInternalServerError, "ServerError", err.Error(), nil)
		return
	}

	GeneralResponse(w, http.StatusOK, "Success", Success, createdTx)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		GeneralResponse(w, http.StatusInternalServerError, ServerError, err.Error(), nil)
		return
	}

	token, expiredTime, err := h.UserUseCase.Login(r.Context(), input.Username, input.Password)
	if err != nil {
		GeneralResponse(w, http.StatusUnauthorized, Unauthorized, err.Error(), nil)
		return
	}

	responseData := map[string]interface{}{
		"username":    input.Username,
		"accessToken": token,
		"expiredAt":   expiredTime,
	}

	GeneralResponse(w, http.StatusOK, Success, "Successful", responseData)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		GeneralResponse(w, http.StatusMethodNotAllowed, "BadRequest", "Invalid method", nil)
		return
	}

	// Ambil ID transaksi dari query parameter
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		GeneralResponse(w, http.StatusBadRequest, "BadRequest", "Missing transaction ID", nil)
		return
	}

	// Convert ID dari string ke integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		GeneralResponse(w, http.StatusBadRequest, "BadRequest", "Invalid transaction ID format", nil)
		return
	}

	// Panggil usecase untuk mendapatkan transaksi berdasarkan ID
	tx, err := h.UserUseCase.GetUserById(r.Context(), id)
	if err != nil {
		GeneralResponse(w, http.StatusInternalServerError, "ServerError", err.Error(), nil)
		return
	}

	if tx == nil {
		GeneralResponse(w, http.StatusNotFound, "NotFound", fmt.Sprintf("Transaction with ID %d not found", id), nil)
		return
	}

	// Response transaksi yang ditemukan
	GeneralResponse(w, http.StatusOK, "Success", "Transaction retrieved successfully", tx)
}
