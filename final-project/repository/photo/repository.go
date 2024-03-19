package photorepository

import (
	"context"
	"database/sql"
	"final-project/model"
	"final-project/repository"
	"time"
)

type photoRepository struct {
	db *sql.DB
}

func New(db *sql.DB) repository.PhotoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Create(ctx context.Context, data model.Photo) (model.Photo, error) {
	var (
		photo model.Photo
		now   = time.Now()
		stmt  = `INSERT INTO 
		photos(title, caption, photo_url, user_id, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6) RETURNING id, title, caption, photo_url, user_id, created_at`
	)

	row := r.db.QueryRowContext(ctx, stmt, data.Title, data.Caption, data.URL, data.UserID, now, now)
	if err := row.Err(); err != nil {
		return photo, err
	}

	err := row.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.URL, &photo.UserID, &photo.CreatedAt)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *photoRepository) FindAll(ctx context.Context) ([]model.Photo, error) {
	var (
		photos []model.Photo
		stmt   = `
		SELECT
			p.id,
			p.title,
			p.caption,
			p.photo_url,
			p.user_id,
			p.created_at,
			p.updated_at,
			u.email,
			u.username
		FROM photos p
		INNER JOIN users u ON p.user_id=u.id`
	)

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var photo model.Photo
		rows.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.URL, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt, &photo.User.Email, &photo.User.Username)
		photos = append(photos, photo)
	}

	return photos, nil
}

func (r *photoRepository) Update(ctx context.Context, tx *sql.Tx, data model.Photo) (model.Photo, error) {
	var (
		photo model.Photo
		now   = time.Now()
		stmt  = `
		UPDATE 
			photos 
		SET 
			title=$1,
			caption=$2,
			photo_url=$3,
			updated_at=$4
		WHERE id=$5
		RETURNING id, title, caption, photo_url, user_id, updated_at`
	)

	row := tx.QueryRowContext(ctx, stmt, data.Title, data.Caption, data.URL, now, data.ID)
	if err := row.Err(); err != nil {
		return photo, err
	}

	err := row.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.URL, &photo.UserID, &photo.UpdatedAt)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *photoRepository) Delete(ctx context.Context, tx *sql.Tx, id uint64) (uint64, error) {
	var ownerID uint64

	stmt := `DELETE FROM photos WHERE id=$1 RETURNING user_id`
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