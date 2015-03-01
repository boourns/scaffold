# scaffold [![Build Status](https://travis-ci.org/boourns/scaffold.svg?branch=master)](https://travis-ci.org/boourns/scaffold)

code generator for go

# Install
```bash
go get github.com/boourns/scaffold/cmd/scaffold
```

# Create a struct
```bash
cat > user.go <<EOF
package user

type User struct {
	Name  string
	Admin bool
}
EOF
```

# Scaffold some code
```
scaffold -in=user.go -out=user_model.go -struct=User -scaffold=model
- Parsing user.go:User
- Generating model
- Saving as user_model.go
```

# Use it

