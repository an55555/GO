package Models

import (
	"GoLang-WEB/Demo/DB"
	"GoLang-WEB/Demo/checkErr"
	"errors"
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