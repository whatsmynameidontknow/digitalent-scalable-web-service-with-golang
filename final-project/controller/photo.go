package controller

import (
	"encoding/json"
	"errors"
	"final-project/dto"
	"final-project/helper"
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
	var (
		data dto.PhotoRequest
		resp helper.Response[dto.PhotoCreateResponse]
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = data.ValidateCreate()
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	photo, err := c.photoService.Create(r.Context(), data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(photo).Code(http.StatusCreated).Send(w)
}

func (c *photoController) GetAll(w http.ResponseWriter, r *http.Request) {
	var resp helper.Response[[]dto.PhotoResponse]

	photos, err := c.photoService.GetAll(r.Context())
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Data(photos).Success(true).Code(http.StatusOK).Send(w)
}

func (c *photoController) Update(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.PhotoRequest
		resp helper.Response[dto.PhotoUpdateResponse]
	)

	photoIDStr := r.PathValue("photoID")
	photoID, err := strconv.ParseUint(photoIDStr, 10, 64)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = data.ValidateUpdate()
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	photo, err := c.photoService.Update(r.Context(), photoID, data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(photo).Code(http.StatusOK).Send(w)
}

func (c *photoController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp helper.Response[any]

	photoIDStr := r.PathValue("photoID")
	photoID, err := strconv.ParseUint(photoIDStr, 10, 64)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = c.photoService.Delete(r.Context(), photoID)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Code(http.StatusOK).Send(w)
}
