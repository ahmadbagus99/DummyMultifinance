package interfaces

import (
	"DummyMultifinance/usecases"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserHandler(uc *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: uc,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.UserUseCase.CreateUser(input.Username, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	GeneralResponse(w, http.StatusOK, Success, "Successful", user)
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

	token, expiredTime, err := h.UserUseCase.Login(input.Username, input.Password)
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
