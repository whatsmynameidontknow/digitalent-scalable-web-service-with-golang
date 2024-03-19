package controller

import (
	"encoding/json"
	"final-project/dto"
	"final-project/service"
	"net/http"
	"strconv"
)

type photoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *photoController {
	return &photoController{photoService}
}

func (c *photoController) Create(w http.ResponseWriter, r *http.Request) {
	var data dto.PhotoRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.ValidateCreate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := c.photoService.Create(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (c *photoController) GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := c.photoService.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *photoController) Update(w http.ResponseWriter, r *http.Request) {
	photoIDStr := r.PathValue("photoID")
	photoID, err := strconv.ParseUint(photoIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data dto.PhotoRequest

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.ValidateUpdate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := c.photoService.Update(r.Context(), photoID, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *photoController) Delete(w http.ResponseWriter, r *http.Request) {
	photoIDStr := r.PathValue("photoID")
	photoID, err := strconv.ParseUint(photoIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.photoService.Delete(r.Context(), photoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.DeleteResponse{
		Message: "your photo has benn successfully deleted",
	})
}
