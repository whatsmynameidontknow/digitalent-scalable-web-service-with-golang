package photorepository

import (
	"context"
	"database/sql"
	"final-project/model"
)

type photoRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Create(ctx context.Context, data model.Photo) (model.Photo, error) {
	var (
		photo model.Photo
		stmt  = `
		INSERT INTO 
			photo(title, caption, url, user_id)
			VALUES($1, $2, $3, $4)
		RETURNING
			id, 
			title, 
			caption, 
			url, 
			user_id, 
			created_at
		`
	)

	row := r.db.QueryRowContext(ctx, stmt, data.Title, data.Caption, data.URL, data.UserID)
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
			p.url,
			p.user_id,
			p.created_at,
			p.updated_at,
			u.email,
			u.username
		FROM photo p
		INNER JOIN user_ u ON p.user_id=u.id
		ORDER BY p.created_at DESC
		`
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
		stmt  = `
		UPDATE 
			photo
		SET 
			title=$1,
			caption=$2,
			url=$3,
			updated_at=NOW()
		WHERE id=$4 AND updated_at=$5
		RETURNING 
			id, 
			title, 
			caption, 
			url, 
			user_id, 
			updated_at
		`
	)

	row := tx.QueryRowContext(ctx, stmt, data.Title, data.Caption, data.URL, data.ID, data.UpdatedAt)
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
	var (
		ownerID uint64
		stmt    = `
		DELETE FROM
			photo
		WHERE id=$1
		RETURNING
			user_id
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

func (r *photoRepository) FindByID(ctx context.Context, id uint64) (model.Photo, error) {
	var (
		photo model.Photo
		stmt  = `
		SELECT
			p.id,
			p.title,
			p.caption,
			p.url,
			p.user_id,
			p.created_at,
			p.updated_at,
			u.email,
			u.username
		FROM photo p
		INNER JOIN user_ u ON p.user_id=u.id
		WHERE p.id=$1`
	)

	row := r.db.QueryRowContext(ctx, stmt, id)
	if err := row.Err(); err != nil {
		return photo, err
	}

	err := row.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.URL, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt, &photo.User.Email, &photo.User.Username)
	if err != nil {
		return photo, err
	}

	return photo, nil
}
