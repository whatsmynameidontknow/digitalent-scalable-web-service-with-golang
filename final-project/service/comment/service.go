package commentservice

import (
	"context"
	"database/sql"
	"errors"
	"final-project/dto"
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
	"final-project/service"
	"net/http"
)

type commentService struct {
	commentRepo repository.CommentRepository
	photoRepo   repository.PhotoRepository
	db          *sql.DB
}

func New(commentRepo repository.CommentRepository, photoRepo repository.PhotoRepository, db *sql.DB) service.CommentService {
	return &commentService{commentRepo, photoRepo, db}
}

func (s *commentService) Create(ctx context.Context, data dto.CommentRequest) (dto.CommentCreateResponse, error) {
	var (
		resp dto.CommentCreateResponse
		err  error
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	_, err = s.photoRepo.FindByID(ctx, data.PhotoID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp, helper.NewResponseError(helper.ErrPhotoNotFound, http.StatusNotFound)
		}
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	comment := model.Comment{
		PhotoID: data.PhotoID,
		UserID:  uint64(userID),
		Message: data.Message,
	}

	comment, err = s.commentRepo.Create(ctx, comment)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	resp = dto.CommentCreateResponse{
		ID:        comment.ID,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		Message:   comment.Message,
		CreatedAt: comment.CreatedAt,
	}

	return resp, nil
}

func (s *commentService) GetAll(ctx context.Context) ([]dto.CommentResponse, error) {
	var resp []dto.CommentResponse

	comments, err := s.commentRepo.FindAll(ctx)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	resp = make([]dto.CommentResponse, 0, len(comments))

	for _, comment := range comments {
		resp = append(resp, dto.CommentResponse{
			ID:        comment.ID,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
			UpdateAt:  comment.UpdatedAt,
			User: dto.User{
				Username: comment.User.Username,
				Email:    comment.User.Email,
			},
			Photo: dto.Photo{
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption.String,
				PhotoURL: comment.Photo.URL,
				UserID:   comment.Photo.UserID,
			},
		})
	}

	return resp, nil
}

func (s *commentService) Update(ctx context.Context, commentID uint64, data dto.CommentRequest) (dto.CommentUpdateResponse, error) {
	var (
		comment model.Comment
		resp    dto.CommentUpdateResponse
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}
	defer helper.RollbackOrCommit(tx, &err)

	comment.Message = data.Message
	comment.ID = commentID
	comment, err = s.commentRepo.Update(ctx, tx, comment)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp, helper.NewResponseError(helper.ErrCommentNotFound, http.StatusNotFound)
		}
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	if comment.UserID != uint64(userID) {
		return resp, helper.NewResponseError(helper.ErrUnauthorized, http.StatusUnauthorized)
	}

	resp = dto.CommentUpdateResponse{
		ID:        comment.ID,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		Message:   comment.Message,
		UpdatedAt: comment.UpdatedAt,
	}

	return resp, nil
}

func (s *commentService) Delete(ctx context.Context, commentID uint64) error {
	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}
	defer helper.RollbackOrCommit(tx, &err)

	ownerID, err := s.commentRepo.Delete(ctx, tx, commentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return helper.NewResponseError(helper.ErrCommentNotFound, http.StatusNotFound)
		}
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	if ownerID != uint64(userID) {
		return helper.NewResponseError(helper.ErrUnauthorized, http.StatusUnauthorized)
	}

	return nil
}
