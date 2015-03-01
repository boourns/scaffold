



package user

import (
	"github.com/boourns/dbutil"
  "database/sql"
  "fmt"
)

func sqlFieldsForUser() string {
  return "ID, Name, Email, ResetToken"
}

func loadUser(rows *sql.Rows) (*User, error) {
	ret := User{}

	err := rows.Scan(&ret.ID,&ret.Name,&ret.Email,&ret.ResetToken)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func Where(tx dbutil.DBLike, where string, whereFields ...interface{}) ([]*User, error) {
  ret := []*User{}
  sql := fmt.Sprintf("SELECT %s from User WHERE %s", sqlFieldsForUser(), where)
	rows, err := tx.Query(sql, whereFields)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
    item, err := loadUser(rows)
    if err != nil {
      return nil, err
    }
    ret = append(ret, item)
	}
  return ret, nil
}

func (s *User) Update(tx dbutil.DBLike) error {
		stmt, err := tx.Prepare(fmt.Sprintf("UPDATE User(%s) VALUES(?,?,?,?) WHERE User.ID = ?", sqlFieldsForUser()))

		if err != nil {
			return err
		}

    params := []interface{}{s.ID,s.Name,s.Email,s.ResetToken}
    params = append(params, s.ID)

		_, err = stmt.Exec(params)
		if err != nil {
			return err
		}

    return nil
}

func (s *User) Insert(tx dbutil.DBLike) error {
		stmt, err := tx.Prepare("INSERT INTO User(Name,Email,ResetToken) VALUES(?,?,?)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(s.Name,s.Email,s.ResetToken)
		if err != nil {
			return err
		}
	  return nil
}

func CreateUserTable(tx dbutil.DBLike) error {
		stmt, err := tx.Prepare(`



CREATE TABLE User (
  
    ID INT NOT NULL PRIMARY KEY,
  
    Name VARCHAR(255),
  
    Email VARCHAR(255),
  
    ResetToken VARCHAR(255)
  
);

`)
		if err != nil {
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	  return nil
}
