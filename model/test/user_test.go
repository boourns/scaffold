package user

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	db = database
}

func TestUserCreateTable(t *testing.T) {
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

	err = v.Insert(db)
	if err != nil {
		t.Errorf("error inserting a new user: %s", err)
	}

	users, err := SelectUser(db, "ORDER BY ID DESC LIMIT 1")

	if err != nil {
		t.Errorf("error selecting all users back out: %s", err)
	}

	if len(users) != 1 {
		t.Fatalf("expected 1 user, received %d", len(users))
	}

	if users[0].ID == 0 {
		t.Error("expected user to have a valid ID")
	}

	if users[0].Name != "Tom" {
		t.Errorf("expected Tom, received %s", users[0].Name)
	}

	if users[0].Email != "tom@tom.com" {
		t.Errorf("expected tom@tom.com, received %s", users[0].Email)
	}

	if users[0].ResetToken != "asdf1234" {
		t.Errorf("expected asdf1234, received %s", users[0].ResetToken)
	}
}
