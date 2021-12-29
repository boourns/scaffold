package user

//go:generate scaffold model

type User struct {
	ID         int64
	Name       string
	Email      string
	ResetToken string
}
