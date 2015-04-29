package ast

type Scaffold interface {
	Description() string
	Generate(*Model) (error)
}
