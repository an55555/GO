package userCtl

import (
	"../httpRouter"
	"fmt"
)

func PutUser(c *odserver.Context) {
	fmt.Println("PUTUSER", c)
	uid := c.GetParams()["uid"]
	fmt.Println("UID=" + uid)
	body := c.PostParams()
	fmt.Println("body===", body)
	for k, v := range body {
		fmt.Println("key111:", k)
		fmt.Println("val111:", v)
	}
}
