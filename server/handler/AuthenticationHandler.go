package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthenticationHandler struct {
	UserService *service.UserService
	TeamService *service.TeamService
}

func NewAuthenticationHandler(userService *service.UserService, teamService *service.TeamService) *AuthenticationHandler {
	return &AuthenticationHandler{UserService: userService, TeamService: teamService}
}

func (handler *AuthenticationHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := handler.UserService.GetByEmail(credentials.Username)
	if err != nil {
		http.Error(w, "Error querying user", http.StatusInternalServerError)
		return
	}

	if user == nil || user.Password != credentials.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if user.Role == 0 { // Recruit role
		token, err := model.GenerateRecruitJWT(user.ID, user.Email, 0)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		authResponse := model.AuthenticationResponse{Token: token}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(authResponse); err != nil {
			http.Error(w, "Error generating response", http.StatusInternalServerError)
		}
	} else { // Employee role
		team, err := handler.TeamService.GetByUserID(int(user.ID))
		if err != nil {
			http.Error(w, "Error querying team", http.StatusInternalServerError)
			return
		}

		token, err := model.GenerateJWT(user.ID, user.Email, team.ID, 1)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		authResponse := model.AuthenticationResponse{Token: token}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(authResponse); err != nil {
			http.Error(w, "Error generating response", http.StatusInternalServerError)
		}
	}
}

func (handler *AuthenticationHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userDTO model.UserCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	user := &model.User{}
	user.FromCreateDTO(&userDTO)

	err := handler.UserService.Create(user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	token, err := model.GenerateRecruitJWT(user.ID, user.Email, 0)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	authResponse := model.AuthenticationResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(authResponse); err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
	}
}
