# scaffold [![Build Status](https://travis-ci.org/boourns/scaffold.svg?branch=master)](https://travis-ci.org/boourns/scaffold)

code generator for go

# Install
```bash
go get github.com/boourns/scaffold/cmd/scaffold
```

# Create a struct
```bash
cat > model.go <<EOF
type User struct {
	Name  string
	Admin bool
}
EOF
```

# Generate some scaffolds
```
scaffold controller model.go User
-> generated UserController
```
