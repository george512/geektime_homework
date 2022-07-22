package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

type User struct {
	Id   int64
	Name string
}

var DB *sql.DB

func init() {
	var err error
	dsn := "root:1998@tcp(localhost:3306)/test"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Open DB failed!")
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Connect DB failed!")
	}
}

func Get(id int) (user User, err error) {
	query := "select * from user where id=?"
	row := DB.QueryRow(query, id)
	err = row.Scan(&user.Id, &user.Name)

	if err != nil {
		return user, errors.Wrap(err, "User.Get:scan failed")
	}
	return user, nil
}

func IsNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
