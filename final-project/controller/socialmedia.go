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

type socialMediaController struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) *socialMediaController {
	return &socialMediaController{socialMediaService}
}

// SocialMediaCreate godoc
// @Summary create a new social media
// @Tags social-media
// @Accept json
// @Produce json
// @Security BearerToken
// @Param request body dto.SocialMediaCreate true "required body"
// @Success 201 {object} helper.Response[dto.SocialMediaCreateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /socialmedias [post]
func (c *socialMediaController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.SocialMediaRequest
		resp = helper.NewResponse[dto.SocialMediaCreateResponse](helper.SocialMediaCreate)
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

	socialMedia, err := c.socialMediaService.Create(r.Context(), data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(socialMedia).Code(http.StatusCreated).Send(w)
}

// SocialMediaGetAll godoc
// @Summary get all social media
// @Tags social-media
// @Produce json
// @Security BearerToken
// @Success 200 {object} helper.Response[[]dto.SocialMediaResponse]
// @Failure 500 {object} helper.Response[any]
// @Router /socialmedias [get]
func (c *socialMediaController) GetAll(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[[]dto.SocialMediaResponse](helper.SocialMediaGetAll)

	socialMedias, err := c.socialMediaService.GetAll(r.Context())
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Data(socialMedias).Success(true).Code(http.StatusOK).Send(w)
}

// SocialMediaUpdate godoc
// @Summary update social media
// @Tags social-media
// @Accept json
// @Produce json
// @Security BearerToken
// @Param socialMediaID path int true "social media ID"
// @Param request body dto.SocialMediaUpdate true "required body"
// @Success 200 {object} helper.Response[dto.SocialMediaUpdateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /socialmedias/{socialMediaID} [put]
func (c *socialMediaController) Update(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.SocialMediaRequest
		resp = helper.NewResponse[dto.SocialMediaUpdateResponse](helper.SocialMediaUpdate)
	)

	socialMediaIDStr := r.PathValue("socialMediaID")
	socialMediaID, err := strconv.ParseUint(socialMediaIDStr, 10, 64)
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

	socialMedia, err := c.socialMediaService.Update(r.Context(), socialMediaID, data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(socialMedia).Code(http.StatusOK).Send(w)
}

// SocialMediaDelete godoc
// @Summary delete social media
// @Tags social-media
// @Produce json
// @Security BearerToken
// @Param socialMediaID path int true "social media ID"
// @Success 200 {object} helper.Response[any]
// @Failure 400 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /socialmedias/{socialMediaID} [delete]
func (c *socialMediaController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[any](helper.SocialMediaDelete)

	socialMediaIDStr := r.PathValue("socialMediaID")
	socialMediaID, err := strconv.ParseUint(socialMediaIDStr, 10, 64)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = c.socialMediaService.Delete(r.Context(), socialMediaID)
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

// SocialMediaGetByID godoc
// @Summary get social media by ID
// @Tags social-media
// @Produce json
// @Security BearerToken
// @Param socialMediaID path int true "social media ID"
// @Success 200 {object} helper.Response[dto.SocialMediaResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /socialmedias/{socialMediaID} [get]
func (c *socialMediaController) GetByID(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[dto.SocialMediaResponse](helper.SocialMediaGetByID)

	socialMediaIDStr := r.PathValue("socialMediaID")
	socialMediaID, err := strconv.ParseUint(socialMediaIDStr, 10, 64)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	socialMedia, err := c.socialMediaService.GetByID(r.Context(), socialMediaID)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Data(socialMedia).Success(true).Code(http.StatusOK).Send(w)
}
