package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(192.168.72.130:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	Db = database
}

func main() {
	res, err := Db.Exec("update user set email=? where id=?", "tes111t@email.com", 1)
	if err != nil {
		fmt.Println("update failed", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows affected failed", err)
	}
	fmt.Println("update", row)
}
