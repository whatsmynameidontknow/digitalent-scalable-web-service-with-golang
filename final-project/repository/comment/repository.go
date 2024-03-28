package commentrepository

import (
	"context"
	"database/sql"
	"final-project/model"
)

type commentRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(ctx context.Context, data model.Comment) (model.Comment, error) {
	var (
		comment model.Comment
		stmt    = `
		INSERT INTO
			comment(message, photo_id, user_id)
			VALUES($1, $2, $3) 
		RETURNING
			id,
			message,
			photo_id,
			user_id,
			created_at
		`
	)

	row := r.db.QueryRowContext(ctx, stmt, data.Message, data.PhotoID, data.UserID)
	if err := row.Err(); err != nil {
		return comment, err
	}

	err := row.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.CreatedAt)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) FindAll(ctx context.Context) ([]model.Comment, error) {
	var (
		comments []model.Comment
		stmt     = `
		SELECT
			c.id,
			c.message,
			c.photo_id,
			c.user_id,
			c.created_at,
			c.updated_at,
			u.username,
			u.email,
			p.title,
			p.caption,
			p.url,
			p.user_id
		FROM comment c
		INNER JOIN user_ u ON c.user_id=u.id
		INNER JOIN photo p ON c.photo_id=p.id
		ORDER BY c.created_at DESC
		`
	)

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment

		err := rows.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.CreatedAt, &comment.UpdatedAt, &comment.User.Username, &comment.User.Email, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.URL, &comment.Photo.UserID)
		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *commentRepository) Update(ctx context.Context, tx *sql.Tx, data model.Comment) (model.Comment, error) {
	var (
		comment model.Comment
		stmt    = `
		UPDATE 
			comment
		SET 
			message=$1,
			updated_at=NOW()
		WHERE id=$2 AND updated_at=$3
		RETURNING 
			id, 
			message, 
			photo_id, 
			user_id, 
			updated_at
		`
	)

	row := tx.QueryRowContext(ctx, stmt, data.Message, data.ID, data.UpdatedAt)
	if err := row.Err(); err != nil {
		return comment, err
	}

	err := row.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.UpdatedAt)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) Delete(ctx context.Context, tx *sql.Tx, id uint64) (uint64, error) {
	var (
		ownerID uint64
		stmt    = `
		DELETE FROM
			comment
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

func (r *commentRepository) FindByID(ctx context.Context, id uint64) (model.Comment, error) {
	var (
		comment model.Comment
		stmt    = `
		SELECT
			c.id,
			c.message,
			c.photo_id,
			c.user_id,
			c.created_at,
			c.updated_at,
			u.username,
			u.email,
			p.title,
			p.caption,
			p.url,
			p.user_id
		FROM comment c
		INNER JOIN user_ u ON c.user_id=u.id
		INNER JOIN photo p ON c.photo_id=p.id
		WHERE c.id=$1
		`
	)

	row := r.db.QueryRowContext(ctx, stmt, id)
	if err := row.Err(); err != nil {
		return comment, err
	}

	err := row.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.CreatedAt, &comment.UpdatedAt, &comment.User.Username, &comment.User.Email, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.URL, &comment.Photo.UserID)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
