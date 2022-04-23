package user

import "time"

//go:generate go run github.com/boourns/scaffold model -config user.json

type User struct {
	ID         int64
	Name       string
	Email      string
	ResetToken string

	CreatedAt time.Time `sqlType:"DATETIME"`
}
