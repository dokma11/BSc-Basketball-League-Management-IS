package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type RecruitHandler struct {
	RecruitService *service.RecruitService
	DraftService   *service.DraftService
}

func NewRecruitHandler(RecruitService *service.RecruitService, DraftService *service.DraftService) *RecruitHandler {
	return &RecruitHandler{RecruitService: RecruitService, DraftService: DraftService}
}

func (handler *RecruitHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	recruits, err := handler.RecruitService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var recruitResponseDTOs []model.RecruitResponseDTO
	for _, recruit := range *recruits {
		var recruitResponseDTO model.RecruitResponseDTO
		recruit.FromModel(&recruitResponseDTO)
		recruitResponseDTOs = append(recruitResponseDTOs, recruitResponseDTO)
	}

	json.NewEncoder(w).Encode(recruitResponseDTOs)
}

func (handler *RecruitHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	recruit, err := handler.RecruitService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recruit == nil {
		http.NotFound(w, r)
		return
	}

	var recruitResponseDTO model.RecruitResponseDTO
	recruit.FromModel(&recruitResponseDTO)
	json.NewEncoder(w).Encode(recruitResponseDTO)
}

func (handler *RecruitHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var recruitDTO model.RecruitCreateDTO
	err := json.NewDecoder(req.Body).Decode(&recruitDTO)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	recruit := &model.Recruit{}
	recruit.FromDTO(&recruitDTO)

	var draft, draftErr = handler.DraftService.GetLatest()
	if draftErr != nil {
		println("Error while getting latest draft")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	recruit.DraftId = draft.ID

	err = handler.RecruitService.Create(recruit)
	if err != nil {
		println("Error while creating a new recruit")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *RecruitHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var recruitDTO model.RecruitCreateDTO
	err := json.NewDecoder(req.Body).Decode(&recruitDTO)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	recruit := &model.Recruit{}
	recruit.FromDTO(&recruitDTO)

	err = handler.RecruitService.Update(recruit)
	if err != nil {
		println("Error while updating recruit")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *RecruitHandler) AddToWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		println(err)
		println(err.Error())
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var recruitDTO model.RecruitCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&recruitDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recruit := &model.Recruit{}
	recruit.FromDTO(&recruitDTO)

	err = handler.RecruitService.AddToWishlist(recruit, teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *RecruitHandler) RemoveFromWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var recruitDTO model.RecruitCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&recruitDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recruit := &model.Recruit{}
	recruit.FromDTO(&recruitDTO)

	err = handler.RecruitService.RemoveFromWishlist(recruit, teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *RecruitHandler) GetAllByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	recruits, err := handler.RecruitService.GetAllByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var recruitResponseDTOs []model.RecruitResponseDTO
	for _, recruit := range *recruits {
		var recruitResponseDTO model.RecruitResponseDTO
		recruit.FromModel(&recruitResponseDTO)
		recruitResponseDTOs = append(recruitResponseDTOs, recruitResponseDTO)
	}

	json.NewEncoder(w).Encode(recruitResponseDTOs)
}
