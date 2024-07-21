package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthenticationHandler struct {
	KorisnikService *service.KorisnikService
}

func NewAuthenticationHandler(korisnikService *service.KorisnikService) *AuthenticationHandler {
	return &AuthenticationHandler{KorisnikService: korisnikService}
}

func (handler *AuthenticationHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := handler.KorisnikService.GetByEmail(credentials.Username)
	if err != nil {
		http.Error(w, "Error querying user", http.StatusInternalServerError)
		return
	}

	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Lozinka), []byte(credentials.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := model.GenerateJWT(user.Email)
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
