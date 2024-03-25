package socialmediarepository

import (
	"context"
	"database/sql"
	"final-project/model"
	"final-project/repository"
)

type socialMediaRepository struct {
	db *sql.DB
}

func New(db *sql.DB) repository.SocialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) Create(ctx context.Context, data model.SocialMedia) (model.SocialMedia, error) {
	var (
		socialMedia model.SocialMedia
		stmt        = `
		INSERT INTO
			social_media(user_id, name, url)
			VALUES($1, $2, $3)
		RETURNING
			id,
			name,
			url,
			user_id,
			created_at
		`
	)

	row := r.db.QueryRowContext(ctx, stmt, data.UserID, data.Name, data.URL)
	if err := row.Err(); err != nil {
		return socialMedia, err
	}

	err := row.Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.URL, &socialMedia.UserID, &socialMedia.CreatedAt)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (r *socialMediaRepository) FindAll(ctx context.Context) ([]model.SocialMedia, error) {
	var (
		socialMedias []model.SocialMedia
		stmt         = `
		SELECT
			s.id,
			s.name,
			s.url,
			s.user_id,
			s.created_at,
			s.updated_at,
			u.username,
			u.email
		FROM social_media s
		INNER JOIN user_ u ON s.user_id = u.id
		ORDER BY s.created_at DESC
		`
	)

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var socialMedia model.SocialMedia
		rows.Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.URL, &socialMedia.UserID, &socialMedia.CreatedAt, &socialMedia.UpdatedAt, &socialMedia.User.Username, &socialMedia.User.Email)
		socialMedias = append(socialMedias, socialMedia)
	}

	return socialMedias, nil
}

func (r *socialMediaRepository) Update(ctx context.Context, tx *sql.Tx, data model.SocialMedia) (model.SocialMedia, error) {
	var (
		socialMedia model.SocialMedia
		stmt        = `
		UPDATE
			social_media
		SET
			name=$1,
			url=$2
		WHERE id=$3
		RETURNING
			id,
			name,
			url,
			user_id,
			updated_at
	`
	)

	row := tx.QueryRowContext(ctx, stmt, data.Name, data.URL, data.ID)
	if err := row.Err(); err != nil {
		return socialMedia, err
	}

	err := row.Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.URL, &socialMedia.UserID, &socialMedia.UpdatedAt)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (r *socialMediaRepository) Delete(ctx context.Context, tx *sql.Tx, id uint64) (uint64, error) {
	var (
		ownerID uint64
		stmt    = `
		DELETE FROM
			social_media
		WHERE id=$1
		RETURNING user_id
		`
	)

	row := tx.QueryRowContext(ctx, stmt, id)
	if err := row.Err(); err != nil {
		return ownerID, err
	}

	err := row.Scan(&ownerID)
	if err != nil {
		return ownerID, err
	}

	return ownerID, nil
}

func (r *socialMediaRepository) FindByID(ctx context.Context, id uint64) (model.SocialMedia, error) {
	var (
		socialMedia model.SocialMedia
		stmt        = `
		SELECT
			s.id,
			s.name,
			s.url,
			s.user_id,
			s.created_at,
			s.updated_at,
			u.username,
			u.email
		FROM social_media s
		INNER JOIN user_ u ON s.user_id = u.id
		WHERE s.id=$1
		`
	)

	row := r.db.QueryRowContext(ctx, stmt, id)
	if err := row.Err(); err != nil {
		return socialMedia, err
	}

	err := row.Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.URL, &socialMedia.UserID, &socialMedia.CreatedAt, &socialMedia.UpdatedAt, &socialMedia.User.Username, &socialMedia.User.Email)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}
