package socialmediaservice

import (
	"context"
	"database/sql"
	"final-project/dto"
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
	"final-project/service"
)

type socialMediaService struct {
	socialMediaRepo repository.SocialMediaRepository
	db              *sql.DB
}

func New(socialMediaRepo repository.SocialMediaRepository, db *sql.DB) service.SocialMediaService {
	return &socialMediaService{socialMediaRepo, db}
}

func (s *socialMediaService) Create(ctx context.Context, data dto.SocialMediaRequest) (dto.SocialMediaCreateResponse, error) {
	var (
		resp dto.SocialMediaCreateResponse
		err  error
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return resp, helper.ErrInternal
	}

	socialMedia := model.SocialMedia{
		UserID: uint64(userID),
		Name:   data.Name,
		URL:    data.URL,
	}

	socialMedia, err = s.socialMediaRepo.Create(ctx, socialMedia)
	if err != nil {
		return resp, err
	}

	resp = dto.SocialMediaCreateResponse{
		ID:        socialMedia.ID,
		UserID:    socialMedia.UserID,
		Name:      socialMedia.Name,
		URL:       socialMedia.URL,
		CreatedAt: socialMedia.CreatedAt,
	}

	return resp, nil
}

func (s *socialMediaService) GetAll(ctx context.Context) ([]dto.SocialMediaResponse, error) {
	var (
		resp []dto.SocialMediaResponse
		err  error
	)

	socialMedias, err := s.socialMediaRepo.FindAll(ctx)
	if err != nil {
		return resp, err
	}

	resp = make([]dto.SocialMediaResponse, 0, len(socialMedias))

	for _, socialMedia := range socialMedias {
		resp = append(resp, dto.SocialMediaResponse{
			ID:        socialMedia.ID,
			UserID:    socialMedia.UserID,
			Name:      socialMedia.Name,
			URL:       socialMedia.URL,
			CreatedAt: socialMedia.CreatedAt,
			UpdatedAt: socialMedia.UpdatedAt,
			User: dto.User{
				Username: socialMedia.User.Username,
			},
		})
	}

	return resp, nil
}

func (s *socialMediaService) Update(ctx context.Context, id uint64, data dto.SocialMediaRequest) (dto.SocialMediaUpdateResponse, error) {
	var (
		resp dto.SocialMediaUpdateResponse
		err  error
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return resp, helper.ErrInternal
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return resp, err
	}

	socialMedia := model.SocialMedia{
		ID:     id,
		UserID: uint64(userID),
		Name:   data.Name,
		URL:    data.URL,
	}

	socialMedia, err = s.socialMediaRepo.Update(ctx, tx, socialMedia)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	if socialMedia.UserID != uint64(userID) {
		tx.Rollback()
		return resp, helper.ErrUnauthorized
	}

	resp = dto.SocialMediaUpdateResponse{
		ID:        socialMedia.ID,
		UserID:    socialMedia.UserID,
		Name:      socialMedia.Name,
		URL:       socialMedia.URL,
		UpdatedAt: socialMedia.UpdatedAt,
	}

	err = tx.Commit()
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *socialMediaService) Delete(ctx context.Context, id uint64) error {
	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return helper.ErrInternal
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ownerID, err := s.socialMediaRepo.Delete(ctx, tx, id)
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
