package interfaces

import (
	"DummyMultifinance/domain/models"
	consumerUseCase "DummyMultifinance/usecases/consumers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ConsumerHandler struct {
	ConsumerUseCase consumerUseCase.ConsumerUseCase
}

func NewConsumerHandler(uc consumerUseCase.ConsumerUseCase) *ConsumerHandler {
	return &ConsumerHandler{
		ConsumerUseCase: uc,
	}
}

func (h *ConsumerHandler) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		GeneralResponse(w, http.StatusMethodNotAllowed, "BadRequest", "Invalid method", nil)
		return
	}

	var tx models.Consumers
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		GeneralResponse(w, http.StatusBadRequest, BadRequest, err.Error(), nil)
		return
	}

	createdTx, err := h.ConsumerUseCase.CreateConsumer(r.Context(), &tx)
	if err != nil {
		GeneralResponse(w, http.StatusInternalServerError, "ServerError", err.Error(), nil)
		return
	}

	GeneralResponse(w, http.StatusOK, "Success", Success, createdTx)
}

func (h *ConsumerHandler) GetConsumer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		GeneralResponse(w, http.StatusMethodNotAllowed, "BadRequest", "Invalid method", nil)
		return
	}

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
	tx, err := h.ConsumerUseCase.GetConsumerById(r.Context(), id)
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
