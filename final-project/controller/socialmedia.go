package controller

import (
	"encoding/json"
	"final-project/dto"
	"final-project/service"
	"net/http"
	"strconv"
)

type socialMediaController struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) *socialMediaController {
	return &socialMediaController{socialMediaService}
}

func (c *socialMediaController) Create(w http.ResponseWriter, r *http.Request) {
	var data dto.SocialMediaRequest

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

	resp, err := c.socialMediaService.Create(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (c *socialMediaController) GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := c.socialMediaService.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *socialMediaController) Update(w http.ResponseWriter, r *http.Request) {
	socialMediaIDStr := r.PathValue("socialMediaID")
	socialMediaID, err := strconv.ParseUint(socialMediaIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data dto.SocialMediaRequest

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

	resp, err := c.socialMediaService.Update(r.Context(), socialMediaID, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *socialMediaController) Delete(w http.ResponseWriter, r *http.Request) {
	socialMediaIDStr := r.PathValue("socialMediaID")
	socialMediaID, err := strconv.ParseUint(socialMediaIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.socialMediaService.Delete(r.Context(), socialMediaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.DeleteResponse{
		Message: "Your social media has been successfully deleted",
	})
}
