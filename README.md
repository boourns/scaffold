# scaffold ![build](https://github.com/boourns/scaffold/actions/workflows/go.yml/badge.svg)

code generator for go

# Install
```bash
go get github.com/boourns/scaffold
```

# Create a struct
```bash
cat > user.go <<EOF
package user

type User struct {
	ID int
	Name  string
	Admin bool
}
EOF
```

# Scaffold some code
```
scaffold model -in=user.go
- Parsing user.go:User
- Generating model
- Saving as user_model.go
```

# Use it
```go
	err := CreateUserTable(db)
	if err != nil {
		t.Errorf("error creating user table: %s", err)
	}

	v := User{}
	v.Name = "Tom"
	v.Email = "tom@tom.com"
	v.ResetToken = "asdf1234"

	err = v.Insert(db)
	if err != nil {
		t.Errorf("error inserting a new user: %s", err)
	}
	users, err := Select(db, "ORDER BY ID DESC LIMIT 1")

    if err != nil {
    	t.Errorf("error selecting all users back out: %s", err)
    }

    if len(users) != 1 {
    	t.Errorf("Expected 1 user, received %d", len(users))
    }

    fmt.Printf("Received user %#v", users[0])
```


