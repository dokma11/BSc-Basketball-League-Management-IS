package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(UserService *service.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (handler *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := handler.UserService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userResponseDTOs []model.UserResponseDTO
	for _, user := range *users {
		var userResponseDTO model.UserResponseDTO
		user.FromModel(&userResponseDTO)
		userResponseDTOs = append(userResponseDTOs, userResponseDTO)
	}

	json.NewEncoder(w).Encode(userResponseDTOs)
}

func (handler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := handler.UserService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.NotFound(w, r)
		return
	}

	var userResponseDTO model.UserResponseDTO
	user.FromModel(&userResponseDTO)
	json.NewEncoder(w).Encode(userResponseDTO)
}

func (handler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var userDTO model.UserUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &model.User{}
	user.FromDTO(&userDTO)

	err := handler.UserService.Update(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
