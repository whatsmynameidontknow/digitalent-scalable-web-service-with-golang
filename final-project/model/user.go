package model

import "time"

type User struct {
	ID                   uint
	Username, Email      string
	Password             []byte
	Age                  uint
	CreatedAt, UpdatedAt time.Time

	Photos []Photo
}
