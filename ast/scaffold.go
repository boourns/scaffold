package ast

import (
	"flag"
)

type Scaffold interface {
	Description() string
	Generate(flags *flag.FlagSet) error
	Details() string
}
