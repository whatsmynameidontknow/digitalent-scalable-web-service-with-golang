package photoservice

import (
	"context"
	"database/sql"
	"errors"
	"final-project/dto"
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
	"final-project/service"
	"log/slog"
	"net/http"
)

type photoService struct {
	photoRepo repository.PhotoRepository
	db        *sql.DB
	logger    *slog.Logger
}

func New(photoRepo repository.PhotoRepository, db *sql.DB, logger *slog.Logger) service.PhotoService {
	return &photoService{photoRepo, db, logger}
}

func (s *photoService) Create(ctx context.Context, data dto.PhotoRequest) (dto.PhotoCreateResponse, error) {
	var (
		resp dto.PhotoCreateResponse
		err  error
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		s.logger.ErrorContext(ctx, "userID is not float64", "cause", "ctx.Value(helper.UserIDKey).(float64)")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	photo := model.Photo{
		Title:  data.Title,
		URL:    data.PhotoURL,
		UserID: uint64(userID),
	}

	if data.Caption != "" {
		photo.Caption.String = data.Caption
		photo.Caption.Valid = true
	}

	photo, err = s.photoRepo.Create(ctx, photo)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.photoRepo.Create")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	resp = dto.PhotoCreateResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		PhotoURL:  photo.URL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	if photo.Caption.Valid {
		resp.Caption = photo.Caption.String
	}

	return resp, nil
}

func (s *photoService) GetAll(ctx context.Context) ([]dto.PhotoResponse, error) {
	var resp []dto.PhotoResponse

	photos, err := s.photoRepo.FindAll(ctx)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.photoRepo.FindAll")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	resp = make([]dto.PhotoResponse, 0, len(photos))

	for _, photo := range photos {
		item := dto.PhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			PhotoURL:  photo.URL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: dto.User{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		}

		if photo.Caption.Valid {
			item.Caption = photo.Caption.String
		}

		resp = append(resp, item)
	}

	return resp, nil
}

func (s *photoService) Update(ctx context.Context, id uint64, data dto.PhotoRequest) (resp dto.PhotoUpdateResponse, err error) {
	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		s.logger.ErrorContext(ctx, "userID is not float64", "cause", "ctx.Value(helper.UserIDKey).(float64)")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	photo := model.Photo{
		ID:    id,
		Title: data.Title,
		URL:   data.PhotoURL,
		Caption: sql.NullString{
			String: data.Caption,
			Valid:  data.Caption != "",
		},
	}
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.db.BeginTx")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}
	defer helper.RollbackOrCommit(tx, &err)

	photo, err = s.photoRepo.Update(ctx, tx, photo)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.photoRepo.Update")
		if errors.Is(err, sql.ErrNoRows) {
			return resp, helper.NewResponseError(helper.ErrPhotoNotFound, http.StatusNotFound)
		}
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	if photo.UserID != uint64(userID) {
		s.logger.ErrorContext(ctx, "user is not the owner of the photo", "cause", "photo.UserID != uint64(userID)")
		return resp, helper.NewResponseError(helper.ErrUnauthorized, http.StatusUnauthorized)
	}

	resp = dto.PhotoUpdateResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption.String,
		PhotoURL:  photo.URL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	return resp, nil
}

func (s *photoService) Delete(ctx context.Context, id uint64) (err error) {
	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		s.logger.ErrorContext(ctx, "userID is not float64", "cause", "ctx.Value(helper.UserIDKey).(float64)")
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.db.BeginTx")
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}
	defer helper.RollbackOrCommit(tx, &err)

	ownerID, err := s.photoRepo.Delete(ctx, tx, id)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.photoRepo.Delete")
		if errors.Is(err, sql.ErrNoRows) {
			return helper.NewResponseError(helper.ErrPhotoNotFound, http.StatusNotFound)
		}
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	if ownerID != uint64(userID) {
		s.logger.ErrorContext(ctx, "user is not the owner of the photo", "cause", "ownerID != uint64(userID)")
		return helper.NewResponseError(helper.ErrUnauthorized, http.StatusUnauthorized)
	}

	return nil
}
