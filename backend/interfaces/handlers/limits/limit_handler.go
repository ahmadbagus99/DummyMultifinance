package interfaces

import (
	"DummyMultifinance/domain/models"
	handlers "DummyMultifinance/interfaces/handlers"
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
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	var tx models.Limits
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, err.Error(), nil)
		return
	}

	createdTx, err := h.LimitUseCase.CreateLimit(r.Context(), &tx)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, handlers.Success, "Success", createdTx)
}

func (h *LimitHandler) GetLimit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Missing limit ID", nil)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Invalid limit ID format", nil)
		return
	}

	tx, err := h.LimitUseCase.GetLimitById(r.Context(), id)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	if tx == nil {
		handlers.GeneralResponse(w, http.StatusNotFound, handlers.DataNotFound, fmt.Sprintf("Limit with ID %d not found", id), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, handlers.Success, "Limit retrieved successfully", tx)
}
