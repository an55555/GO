package CTL

import (
	"../httpRouter"
	"GoLang-WEB/Demo/models"
)

func CheckUser(c *odserver.Context) {
	resp := Resp{
		RetCode: DEFAULT_RETCODE,
		Msg:     "登录成功",
		Data:    "",
	}
	_resp := &resp
	defer RESP(c, _resp)
	body := c.PostParams()
	result, err := Models.VerifyUser(body)
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     "修改信息失败",
			Data:    err.Error(),
		}
		return
	}
	if result == 0 {
		resp = Resp{
			RetCode: "0",
			Msg:     "用户名或密码不对",
			Data:    "",
		}
		return
	}
}
