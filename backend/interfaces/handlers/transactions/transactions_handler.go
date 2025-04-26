package interfaces

import (
	handlers "DummyMultifinance/interfaces/handlers"
	transactionUseCase "DummyMultifinance/usecases/transactions"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	TransactionUseCase transactionUseCase.TransactionUseCase
}

func NewTransactionHandler(uc transactionUseCase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		TransactionUseCase: uc,
	}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	var requestData struct {
		ConsumerID int     `json:"consumer_id"`
		AssetName  string  `json:"asset_name"`
		Amount     float64 `json:"amount"`
		Tenor      int     `json:"tenor"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, err.Error(), nil)
		return
	}

	loanRequestTransaction, err := h.TransactionUseCase.RequestTransaction(r.Context(), requestData.ConsumerID, requestData.AssetName, requestData.Tenor, requestData.Amount)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, err.Error(), nil)
		return
	}

	createdTx, err := h.TransactionUseCase.CreateTransaction(r.Context(), loanRequestTransaction)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, handlers.Success, "Transaction created successfully", createdTx)
}

func (h *TransactionHandler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Missing transaction ID", nil)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Invalid transaction ID format", nil)
		return
	}

	tx, err := h.TransactionUseCase.GetTransactionById(r.Context(), id)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.SystemError, err.Error(), nil)
		return
	}

	if tx == nil {
		handlers.GeneralResponse(w, http.StatusNotFound, handlers.DataFound, fmt.Sprintf("Transaction with ID %d not found", id), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, handlers.Success, "Transaction retrieved successfully", tx)
}
