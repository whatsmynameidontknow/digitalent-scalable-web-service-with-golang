package userrepository

import (
	"context"
	"database/sql"
	"final-project/model"
	"final-project/repository"
	"time"
)

type userRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) repository.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Create(ctx context.Context, data model.User) (model.User, error) {
	var (
		user model.User
		now  = time.Now()
		stmt = `INSERT INTO users(username, email, password, age, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id, username, email, age`
	)

	row := r.DB.QueryRowContext(ctx, stmt, data.Username, data.Email, data.Password, data.Age, now, now)
	if err := row.Err(); err != nil {
		return user, err
	}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Age)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var (
		user model.User
		stmt = `SELECT id, password FROM users WHERE email = $1`
	)

	row := r.DB.QueryRowContext(ctx, stmt, email)
	if err := row.Err(); err != nil {
		return user, err
	}

	err := row.Scan(&user.ID, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, data model.User) (model.User, error) {
	var (
		user model.User
		now  = time.Now()
		stmt = `UPDATE users SET email=$1, password=$2, updated_at=$3 WHERE id=$4 RETURNING id, username, email, age, updated_at`
	)

	row := r.DB.QueryRowContext(ctx, stmt, data.Email, data.Password, now, data.ID)
	if err := row.Err(); err != nil {
		return user, err
	}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, userID uint) error {
	var (
		stmt = `DELETE FROM users WHERE id=$1`
	)
	res, err := r.DB.ExecContext(ctx, stmt, userID)
	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n == 0 {
		return sql.ErrNoRows
	}

	return nil
}
