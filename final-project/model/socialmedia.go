package model

import "time"

type SocialMedia struct {
	ID, UserID           uint
	Name, URL            string
	CreatedAt, UpdatedAt time.Time
}
