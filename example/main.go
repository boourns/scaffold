package main

import (
	"database/sql"
	"github.com/boourns/dbutil"
	"github.com/boourns/scaffold/example/user"
	"log"
	"os"
)

var db *sql.DB

func init() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		panic("set DATABASE_URL like user:pass@tcp(127.0.0.1:3306)/dbname")
	}
	db = dbutil.Connect(url)
}

func main() {
	err := user.CreateUserTable(db)
	if err != nil {
		log.Fatalf("Error creating table: %s", err)
	}

	u := &user.User{Name: "Tom", Admin: true}
	err = u.Insert(db)
	if err != nil {
		log.Fatalf("Error saving user: %s", err)
	}

	results, err := user.Where(db, "Name = ?", "Tom")
	if err != nil {
		log.Fatalf("Error loading users: %s", err)
	}

	if len(results) == 0 {
		log.Fatalf("Could not find user we just saved")
	}

	existing := results[0]
	log.Printf("From database: %v\n", existing)

	existing.Admin = false
	err = existing.Update(db)
	if err != nil {
		log.Fatalf("Error saving user: %s", err)
	}
}
