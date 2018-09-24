package DB

import (
	"GoLang-WEB/Demo/checkErr"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:55555yyy@tcp(127.0.0.1:3306)/demo")
	Db = db
	fmt.Println(err)
}

func ClearTrnsaction(tx *sql.Tx, error error) {
	fmt.Println("执行回滚")
	if error := recover(); error != nil {
		fmt.Println("Panic info is: ", error)
	}
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		checkErr.Check(err)
	}
}
