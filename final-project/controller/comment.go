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

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *commentController {
	return &commentController{commentService}
}

// CommentCreate godoc
// @Summary create a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerToken
// @Param request body dto.CommentCreate true "required body"
// @Success 201 {object} helper.Response[dto.CommentCreateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 401 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /comments [post]
func (c *commentController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.CommentRequest
		resp = helper.NewResponse[dto.CommentCreateResponse](helper.CommentCreate)
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

	comment, err := c.commentService.Create(r.Context(), data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(comment).Code(http.StatusCreated).Send(w)
}

// CommentGetAll godoc
// @Summary get all comments
// @Tags comments
// @Produce json
// @Security BearerToken
// @Success 200 {object} helper.Response[[]dto.CommentResponse]
// @Failure 401 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /comments [get]
func (c *commentController) GetAll(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[[]dto.CommentResponse](helper.CommentGetAll)

	comments, err := c.commentService.GetAll(r.Context())
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Data(comments).Success(true).Code(http.StatusOK).Send(w)
}

// CommentUpdate godoc
// @Summary update a comment
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerToken
// @Param commentID path int true "comment ID"
// @Param request body dto.CommentUpdate true "required body"
// @Success 200 {object} helper.Response[dto.CommentUpdateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 401 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /comments/{commentID} [put]
func (c *commentController) Update(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.CommentRequest
		resp = helper.NewResponse[dto.CommentUpdateResponse](helper.CommentUpdate)
	)

	commentIDStr := r.PathValue("commentID")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		resp.Error(helper.ErrInvalidID).Code(http.StatusBadRequest).Send(w)
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

	comment, err := c.commentService.Update(r.Context(), commentID, data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(comment).Code(http.StatusOK).Send(w)
}

// CommentDelete godoc
// @Summary delete a comment
// @Tags comments
// @Produce json
// @Security BearerToken
// @Param commentID path int true "comment ID"
// @Success 200 {object} helper.Response[any]
// @Failure 400 {object} helper.Response[any]
// @Failure 401 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /comments/{commentID} [delete]
func (c *commentController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[any](helper.CommentDelete)

	commentIDStr := r.PathValue("commentID")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		resp.Error(helper.ErrInvalidID).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = c.commentService.Delete(r.Context(), commentID)
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

// CommentGetByID godoc
// @Summary get a comment by ID
// @Tags comments
// @Produce json
// @Security BearerToken
// @Param commentID path int true "comment ID"
// @Success 200 {object} helper.Response[dto.CommentResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 401 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /comments/{commentID} [get]
func (c *commentController) GetByID(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[dto.CommentResponse](helper.CommentGetByID)

	commentIDStr := r.PathValue("commentID")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		resp.Error(helper.ErrInvalidID).Code(http.StatusBadRequest).Send(w)
		return
	}

	comment, err := c.commentService.GetByID(r.Context(), commentID)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Data(comment).Success(true).Code(http.StatusOK).Send(w)
}
