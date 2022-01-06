package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func initDb() (db *sql.DB, err error) {
	dsn := `root:yihao8208817..@tcp(127.0.0.1:3306)/users`
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open fail", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping fail:", err)
		return
	}
	return
}
