package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8mb4")
}

type User struct {
	Name string
}

func Query() (User, error) {
	var user User
	err := db.QueryRow("select name from user where id = 1").Scan(&user.Name)
	if err != nil {
		return user, errors.Wrap(err, "scan error")
	}
	return user, nil
}