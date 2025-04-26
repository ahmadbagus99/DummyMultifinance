package interfaces

import (
	"DummyMultifinance/domain/models"
	limitUseCase "DummyMultifinance/usecases/limits"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type LimitHandler struct {
	LimitUseCase limitUseCase.LimiUseCase
}

func NewLimitHandler(uc limitUseCase.LimiUseCase) *LimitHandler {
	return &LimitHandler{
		LimitUseCase: uc,
	}
}

func (h *LimitHandler) CreateLimit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		GeneralResponse(w, http.StatusMethodNotAllowed, "BadRequest", "Invalid method", nil)
		return
	}

	var tx models.Limits
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		GeneralResponse(w, http.StatusBadRequest, BadRequest, err.Error(), nil)
		return
	}

	createdTx, err := h.LimitUseCase.CreateLimit(r.Context(), &tx)
	if err != nil {
		GeneralResponse(w, http.StatusInternalServerError, "ServerError", err.Error(), nil)
		return
	}

	GeneralResponse(w, http.StatusOK, "Success", Success, createdTx)
}

func (h *LimitHandler) GetLimit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		GeneralResponse(w, http.StatusMethodNotAllowed, "BadRequest", "Invalid method", nil)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		GeneralResponse(w, http.StatusBadRequest, "BadRequest", "Missing transaction ID", nil)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		GeneralResponse(w, http.StatusBadRequest, "BadRequest", "Invalid transaction ID format", nil)
		return
	}

	tx, err := h.LimitUseCase.GetLimitById(r.Context(), id)
	if err != nil {
		GeneralResponse(w, http.StatusInternalServerError, "ServerError", err.Error(), nil)
		return
	}

	if tx == nil {
		GeneralResponse(w, http.StatusNotFound, "NotFound", fmt.Sprintf("Transaction with ID %d not found", id), nil)
		return
	}

	GeneralResponse(w, http.StatusOK, "Success", "Transaction retrieved successfully", tx)
}
