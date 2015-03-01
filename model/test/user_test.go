package user

import (
	"database/sql"
	"github.com/boourns/dbutil"
	"os"
	"testing"
)

var db *sql.DB

func init() {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		panic("TEST_DATABASE_URL is not set, expected user:pass@tcp(127.0.0.1:3306)/dbname")
	}
	db = dbutil.Connect(url)
}

func TestUserCreateTable(t *testing.T) {
	err := CreateUserTable(db)

	if err != nil {
		t.Errorf("error creating user table: %s", err)
	}

	v := User{}
	err = v.Insert(db)

	if err != nil {
		t.Errorf("error inserting a new user: %s", err)
	}

}
