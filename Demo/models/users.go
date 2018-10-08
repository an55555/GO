package Models

import (
	"GoLang-WEB/Demo/DB"
	"GoLang-WEB/Demo/checkErr"
	"GoLang-WEB/Demo/utils/encrypt"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var db = DB.Db

type BaseUser struct {
	userName string `db: "username"`
	sex      int    `db: "sex"`
	created  int64  `db: "created"'`
}
type User struct {
	BaseUser
	passWord string `db: "password"`
}

func ParamsToSqlPrepare(params map[string]interface{}) (string, []interface{}) {
	sqlString := []string{}
	execValue := []interface{}{}
	for i, v := range params {
		sqlString = append(sqlString, i+"= ?")
		execValue = append(execValue, v)
	}
	return strings.Join(sqlString, ", "), execValue
}

func InsertUser(params map[string]interface{}) (int64, error) {
	count, err := UserCount(params["username"].(string))
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errors.New("该用户名已经存在")
	}
	params["password"] = encrypt.EncryptSailt(params["password"].(string))
	tx, err := db.Begin()
	sqlString, exec := ParamsToSqlPrepare(params)
	defer DB.ClearTrnsaction(tx, err)
	doSQLPrepare := "INSERT userlist set " + sqlString
	stmt, err := tx.Prepare(doSQLPrepare)
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	res, err := stmt.Exec(exec...)
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	affect, _ := res.LastInsertId()
	if err := tx.Commit(); err != nil {
		checkErr.Check(err)
	}
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	return affect, nil
}

func UpdateUser(uid int, params map[string]interface{}) (int64, error) {
	sqlString, exec := ParamsToSqlPrepare(params)
	exec = append(exec, uid)
	if len(sqlString) == 0 {
		return 0, errors.New("没有需要修改的值")
	}
	tx, err := db.Begin()
	defer DB.ClearTrnsaction(tx, err)
	doSQLPrepare := "update userlist set " + sqlString + " where uid=?"
	stmt, err := tx.Prepare(doSQLPrepare)
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	res, err := stmt.Exec(exec...)
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	affect, _ := res.RowsAffected()
	if err := tx.Commit(); err != nil {
		checkErr.Check(err)
	}
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	if affect == 0 {
		return 0, errors.New("没有需要修改的值")
	}
	return affect, nil
}

func DeleteUser(uid int) (int64, error) {
	tx, err := db.Begin()
	defer DB.ClearTrnsaction(tx, err)
	stmt, err := tx.Prepare("delete from userlist where uid = ?")
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	res, err := stmt.Exec(uid)
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	affect, _ := res.RowsAffected()
	if err := tx.Commit(); err != nil {
		checkErr.Check(err)
	}
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	if affect == 0 {
		return 0, errors.New("删除失败")
	}
	return affect, nil
}

func UserDetail(uid int) (int64, string, error) {
	tx, err := db.Begin()
	var getUid string
	defer DB.ClearTrnsaction(tx, err)
	err = tx.QueryRow("select uid from userlist where uid = ?", uid).Scan(&getUid)
	if err == sql.ErrNoRows {
		return 0, getUid, errors.New("No user with that ID.")
	} else if err != nil {
		checkErr.Check(err)
		return 0, getUid, err
	}

	if err := tx.Commit(); err != nil {
		checkErr.Check(err)
	}
	if err != nil {
		checkErr.Check(err)
		return 0, getUid, err
	}
	return 1, getUid, nil
}

func UserCount(username string) (int, error) {
	var count int
	tx, err := db.Begin()
	defer DB.ClearTrnsaction(tx, err)
	err = tx.QueryRow("select count(username) from userlist where username = ?", username).Scan(&count)
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		checkErr.Check(err)
	}
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	return count, nil
}

func VerifyUser(params map[string]interface{}) (int, error) {
	userName := params["username"].(string)
	password := params["password"].(string)
	var count int
	tx, err := db.Begin()
	defer DB.ClearTrnsaction(tx, err)
	err = tx.QueryRow("select count(uid) from userlist where username = ? AND password = ?", userName, password).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, errors.New("No Find User")
	} else if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		checkErr.Check(err)
	}
	if err != nil {
		checkErr.Check(err)
		return 0, err
	}
	fmt.Println("count", count)
	return count, nil
}
