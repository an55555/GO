package CTL

import (
	"../httpRouter"
	"encoding/json"
	"fmt"
)

type Resp struct {
	RetCode string      `json:"retCode"`
	Msg     string      `json:msg, omitempty`
	Data    interface{} `json:"data, omitempty"`
}

func RESP(c *odserver.Context, resData *Resp) {
	fmt.Println("执行")
	c.GoResW().Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(resData)
	c.GoResW().Write(res)
}
