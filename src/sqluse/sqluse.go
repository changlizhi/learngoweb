package sqluse

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func FindByUser(pk int) int {
	var num int = 0
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hanfuxin?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtOut, err := db.Prepare("select id from user where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(pk).Scan(&num)
	fmt.Println(num)
	if err != nil {
		panic(err.Error())
	}
	return num
}
