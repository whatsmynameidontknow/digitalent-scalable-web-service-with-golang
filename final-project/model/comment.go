package model

import "time"

type Comment struct {
	ID, UserID, PhotoID  uint
	Message              string
	CreatedAt, UpdatedAt time.Time
}
