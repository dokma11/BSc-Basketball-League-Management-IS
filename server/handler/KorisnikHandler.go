package handler

import (
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type KorisnikHandler struct {
	KorisnikService *service.KorisnikService
}

func NewKorisnikHandler(korisnikService *service.KorisnikService) *KorisnikHandler {
	return &KorisnikHandler{KorisnikService: korisnikService}
}

func (handler *KorisnikHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	users, err := handler.KorisnikService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users) // Proveriti samo da li valja
}

func (handler *KorisnikHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := handler.KorisnikService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(user)
}
