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
		fmt.Println("open database error:", err)
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into user(name, email) values (?, ?)", "kevin", "test@email.com")
	if err != nil {
		fmt.Println("insert error:", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("insert error:", err)
		return
	}
	fmt.Println("insert id:", id)
}
