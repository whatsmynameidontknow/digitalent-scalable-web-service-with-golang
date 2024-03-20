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

// PhotoCreate godoc
// @Summary create a new photo
// @Tags photos
// @Accept json
// @Produce json
// @Security BearerToken
// @Param request body dto.PhotoCreate true "required body"
// @Success 201 {object} helper.Response[dto.PhotoCreateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /photos [post]
func (c *photoController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.PhotoRequest
		resp = helper.NewResponse[dto.PhotoCreateResponse](helper.PhotoCreate)
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

// PhotoGetAll godoc
// @Summary get all photos
// @Tags photos
// @Accept json
// @Produce json
// @Security BearerToken
// @Success 200 {object} helper.Response[[]dto.PhotoResponse]
// @Failure 500 {object} helper.Response[any]
// @Router /photos [get]
func (c *photoController) GetAll(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[[]dto.PhotoResponse](helper.PhotoGetAll)

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

// PhotoUpdate godoc
// @Summary update a photo
// @Tags photos
// @Accept json
// @Produce json
// @Security BearerToken
// @Param photoID path int true "photo id"
// @Param request body dto.PhotoUpdate true "required body"
// @Success 200 {object} helper.Response[dto.PhotoUpdateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 401 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /photos/{photoID} [put]
func (c *photoController) Update(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.PhotoRequest
		resp = helper.NewResponse[dto.PhotoUpdateResponse](helper.PhotoUpdate)
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

// PhotoDelete godoc
// @Summary delete a photo
// @Tags photos
// @Produce json
// @Security BearerToken
// @Param photoID path int true "photo id"
// @Success 200 {object} helper.Response[any]
// @Failure 400 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /photos/{photoID} [delete]
func (c *photoController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[any](helper.PhotoDelete)

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

// PhotoGetByID godoc
// @Summary get a photo by id
// @Tags photos
// @Produce json
// @Security BearerToken
// @Param photoID path int true "photo id"
// @Success 200 {object} helper.Response[dto.PhotoResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 401 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /photos/{photoID} [get]
func (c *photoController) GetByID(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[dto.PhotoResponse](helper.PhotoGetByID)

	photoIDStr := r.PathValue("photoID")
	photoID, err := strconv.ParseUint(photoIDStr, 10, 64)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	photo, err := c.photoService.GetByID(r.Context(), photoID)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Data(photo).Success(true).Code(http.StatusOK).Send(w)
}
