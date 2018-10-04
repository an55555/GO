/*package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"log"
)

func DBDemo() {
	db, err := sql.Open("mysql", "root:55555yyy@tcp(127.0.0.1:3306)/demo")
	checkErr(err)

	err = db.Ping()
	if err !=nil {
		fmt.Println("数据库未成功连接")
	}

	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	/*	stmt, err = db.Prepare("delete from userinfo where uid=?")
		checkErr(err)

		res, err = stmt.Exec(id)
		checkErr(err)

		affect, err = res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)*/

	db.Close()
}

func QueryRow()  {
	db, err := sql.Open("mysql", "root:55555yyy@tcp(127.0.0.1:3306)/demo")
	var username string
	err = db.QueryRow("select username from userinfo where uid = ?", 2).Scan(&username)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(username)
	db.Close()
}

func Inset()  {
	db, err := sql.Open("mysql", "root:55555yyy@tcp(127.0.0.1:3306)/demo")
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	for res.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = res.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	id, err := res.LastInsertId()
	checkErr(err)
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", id, rowCnt)
	db.Close()
}

func init() {
	Inset()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
*/