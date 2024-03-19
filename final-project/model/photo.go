package model

import (
	"database/sql"
	"time"
)

type Photo struct {
	ID, UserID           uint
	Title, URL           string
	Caption              sql.NullString
	CreatedAt, UpdatedAt time.Time
}
