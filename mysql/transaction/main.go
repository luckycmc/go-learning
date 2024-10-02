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
		fmt.Println("open err:", err)
		return
	}
	Db = db
}

func main() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin err:", err)
		return
	}
	r, err := conn.Exec("insert into user(name,email)values (?,?)", "tom", "tom@email.com")
	if err != nil {
		fmt.Println("insert err:", err)
		err := conn.Rollback()
		if err != nil {
			return
		}
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("insert err:", err)
		err := conn.Rollback()
		if err != nil {
			return
		}
		return
	}
	fmt.Println("id:", id)

	r, err = conn.Exec("insert into user(name,email) values(?,?)", "jerry", "jerry@email.com")
	if err != nil {
		fmt.Println("insert err:", err)
		err := conn.Rollback()
		if err != nil {
			return
		}
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("insert err:", err)
		err := conn.Rollback()
		if err != nil {
			return
		}
		return
	}
	fmt.Println("id:", id)
	err = conn.Commit()
	if err != nil {
		return
	}
}
