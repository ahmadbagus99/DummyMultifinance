package interfaces

import (
	"DummyMultifinance/domain/models"
	handlers "DummyMultifinance/interfaces/handlers"
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
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	var tx models.Consumers
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, err.Error(), nil)
		return
	}

	createdTx, err := h.ConsumerUseCase.CreateConsumer(r.Context(), &tx)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, handlers.Success, "Success", createdTx)
}

func (h *ConsumerHandler) GetConsumerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Missing consumer ID", nil)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Invalid consumer ID format", nil)
		return
	}

	tx, err := h.ConsumerUseCase.GetConsumerById(r.Context(), id)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	if tx == nil {
		handlers.GeneralResponse(w, http.StatusNotFound, handlers.DataNotFound, fmt.Sprintf("Consumer with ID %d not found", id), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, "Success", "Consumer retrieved successfully", tx)
}

func (h *ConsumerHandler) GetAllConsumer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	tx, err := h.ConsumerUseCase.GetAllConsumer(r.Context())
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	if tx == nil {
		handlers.GeneralResponse(w, http.StatusNotFound, handlers.DataNotFound, "not found", nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, "Success", "Consumer retrieved successfully", tx)
}

func (h *ConsumerHandler) GetConsumerLimit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Missing consumer ID", nil)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Invalid consumer ID format", nil)
		return
	}

	tx, err := h.ConsumerUseCase.GetConsumerLimit(r.Context(), id)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	if tx == nil {
		handlers.GeneralResponse(w, http.StatusNotFound, handlers.DataNotFound, fmt.Sprintf("Consumer with ID %d not found", id), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, "Success", "Consumer retrieved successfully", tx)
}
