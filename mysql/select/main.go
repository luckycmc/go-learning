package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

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
	var person []Person
	err := Db.Select(&person, "select id, name,email from user where id = ?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)
}
