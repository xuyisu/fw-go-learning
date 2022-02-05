package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//连接数据库
func main() {
	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/fw_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
