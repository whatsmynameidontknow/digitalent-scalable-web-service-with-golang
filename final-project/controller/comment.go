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

func (c *commentController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.CommentRequest
		resp helper.Response[dto.CommentCreateResponse]
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

func (c *commentController) GetAll(w http.ResponseWriter, r *http.Request) {
	var resp helper.Response[[]dto.CommentResponse]

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

func (c *commentController) Update(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.CommentRequest
		resp helper.Response[dto.CommentUpdateResponse]
	)

	commentIDStr := r.PathValue("commentID")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
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

func (c *commentController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp helper.Response[any]

	commentIDStr := r.PathValue("commentID")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
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
