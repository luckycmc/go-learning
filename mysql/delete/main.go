package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:root@tcp(192.168.72.130:3306)/test")
	if err != nil {
		fmt.Println("database open failed: ", err)
		return
	}
	Db = db
}

func main() {
	res, err := Db.Exec("delete from user where id=?", 1)
	if err != nil {
		fmt.Println("delete failed: ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows affected failed: ", err)
	}
	fmt.Println("affected rows: ", rows)
}
