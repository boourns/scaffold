package user

import (
	"database/sql"
	"fmt"
	"github.com/boourns/dbutil"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var db *sql.DB

func init() {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		panic("TEST_DATABASE_URL is not set, expected sqlite3://test.db")
	}
	db = dbutil.Connect(url)
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
		t.Errorf("Expected 1 user, received %d", len(users))
	}

	fmt.Printf("Received user %#v", users[0])
}
