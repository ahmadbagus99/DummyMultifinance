package interfaces

import (
	"DummyMultifinance/domain/models"
	handlers "DummyMultifinance/interfaces/handlers"
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
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	var tx models.Users
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, err.Error(), nil)
		return
	}

	createdTx, err := h.UserUseCase.CreateUser(r.Context(), &tx)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, handlers.Success, "Success", createdTx)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	token, expiredTime, err := h.UserUseCase.Login(r.Context(), input.Username, input.Password)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusUnauthorized, handlers.Unauthorized, err.Error(), nil)
		return
	}

	responseData := map[string]interface{}{
		"username":    input.Username,
		"accessToken": token,
		"expiredAt":   expiredTime,
	}

	handlers.GeneralResponse(w, http.StatusOK, handlers.Success, "Success", responseData)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handlers.GeneralResponse(w, http.StatusMethodNotAllowed, handlers.BadRequest, "Invalid method", nil)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		handlers.GeneralResponse(w, http.StatusBadRequest, handlers.BadRequest, "Missing user ID", nil)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusBadRequest, "BadRequest", "Invalid user ID format", nil)
		return
	}

	tx, err := h.UserUseCase.GetUserById(r.Context(), id)
	if err != nil {
		handlers.GeneralResponse(w, http.StatusInternalServerError, handlers.ServerError, err.Error(), nil)
		return
	}

	if tx == nil {
		handlers.GeneralResponse(w, http.StatusNotFound, handlers.DataNotFound, fmt.Sprintf("User with ID %d not found", id), nil)
		return
	}

	handlers.GeneralResponse(w, http.StatusOK, "Success", "User retrieved successfully", tx)
}
