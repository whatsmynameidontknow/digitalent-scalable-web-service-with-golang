package socialmediaservice

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
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	socialMedia := model.SocialMedia{
		UserID: uint64(userID),
		Name:   data.Name,
		URL:    data.URL,
	}

	socialMedia, err = s.socialMediaRepo.Create(ctx, socialMedia)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
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
	var resp []dto.SocialMediaResponse

	socialMedias, err := s.socialMediaRepo.FindAll(ctx)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
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
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}
	defer helper.RollbackOrCommit(tx, &err)

	socialMedia := model.SocialMedia{
		ID:     id,
		UserID: uint64(userID),
		Name:   data.Name,
		URL:    data.URL,
	}

	socialMedia, err = s.socialMediaRepo.Update(ctx, tx, socialMedia)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp, helper.NewResponseError(helper.ErrSocialMediaNotFound, http.StatusNotFound)
		}
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	if socialMedia.UserID != uint64(userID) {
		return resp, helper.NewResponseError(helper.ErrUnauthorized, http.StatusUnauthorized)
	}

	resp = dto.SocialMediaUpdateResponse{
		ID:        socialMedia.ID,
		UserID:    socialMedia.UserID,
		Name:      socialMedia.Name,
		URL:       socialMedia.URL,
		UpdatedAt: socialMedia.UpdatedAt,
	}

	return resp, nil
}

func (s *socialMediaService) Delete(ctx context.Context, id uint64) error {
	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}
	defer helper.RollbackOrCommit(tx, &err)

	ownerID, err := s.socialMediaRepo.Delete(ctx, tx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return helper.NewResponseError(helper.ErrSocialMediaNotFound, http.StatusNotFound)
		}
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	if ownerID != uint64(userID) {
		return helper.NewResponseError(helper.ErrUnauthorized, http.StatusUnauthorized)
	}

	return nil
}
