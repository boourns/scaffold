package user

type User struct {
	Name  string
	Admin bool `{"field": "yolo"}`
}

func Yolo() {
}
