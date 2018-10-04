package main

import (
	"fmt"
)

func test(s string, n ...interface{}) string {
	var x int
	for k, i := range n {
		fmt.Println(k, ":", i)
	}

	return fmt.Sprintf(s, x)
}

type BaseUser struct {
	userName string `db: "username"`
	sex      int    `db: "sex"`
	created  int64  `db: "created"'`
}
type User struct {
	BaseUser
	passWord string `db: "password"`
}

func main() {
	Any(2)
	Any("666")
}
func Any(v interface{}) {

	if v2, ok := v.(string); ok {
		println(v2)
	} else if v3, ok2 := v.(int); ok2 {
		println(v3)
	}
}
