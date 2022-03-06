package user

import "time"

type User struct {
	Name      string
	Admin     bool      `field:yolo`
	UpdatedAt time.Time `json:"updatedAt" sqlType:"DATETIME"`
}

func Yolo() {
}
