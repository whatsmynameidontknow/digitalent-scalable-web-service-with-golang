package photoservice

import (
	"context"
	"database/sql"
	"final-project/dto"
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
	"final-project/service"
)

type photoService struct {
	photoRepo repository.PhotoRepository
	db        *sql.DB
}

func New(photoRepo repository.PhotoRepository, db *sql.DB) service.PhotoService {
	return &photoService{photoRepo, db}
}

func (s *photoService) Create(ctx context.Context, data dto.PhotoRequest) (dto.PhotoCreateResponse, error) {
	var (
		resp dto.PhotoCreateResponse
		err  error
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return resp, helper.ErrInternal
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
		return resp, err
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
	var (
		resp []dto.PhotoResponse
		err  error
	)

	photos, err := s.photoRepo.FindAll(ctx)
	if err != nil {
		return resp, err
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
			User: dto.UserResponse{
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

func (s *photoService) Update(ctx context.Context, id uint64, data dto.PhotoRequest) (dto.PhotoUpdateResponse, error) {
	var (
		resp dto.PhotoUpdateResponse
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return resp, helper.ErrInternal
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
		return resp, err
	}

	photo, err = s.photoRepo.Update(ctx, tx, photo)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	if photo.UserID != uint64(userID) {
		tx.Rollback()
		return resp, helper.ErrUnauthorized
	}

	resp = dto.PhotoUpdateResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption.String,
		PhotoURL:  photo.URL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	err = tx.Commit()
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *photoService) Delete(ctx context.Context, id uint64) error {
	var (
		err error
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return helper.ErrInternal
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ownerID, err := s.photoRepo.Delete(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if ownerID != uint64(userID) {
		tx.Rollback()
		return helper.ErrUnauthorized
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
