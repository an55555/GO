package CTL

import (
	"../httpRouter"
	"GoLang-WEB/Demo/models"
	"GoLang-WEB/Demo/utils/jwt"
	"fmt"
	"net/http"
	"time"
)

const setCookieName string = "iToken"

func WriteCookieServer(c *odserver.Context, username string) {
	nowTime := time.Now()
	addTime, _ := time.ParseDuration("10m")
	nowTime = nowTime.Add(addTime)
	tokenPayload := jwt.NewPayload(username)
	tokenTest := jwt.JwtCode(tokenPayload)
	cookie := http.Cookie{Name: setCookieName, Value: tokenTest, Expires: nowTime, Path: "/"}
	http.SetCookie(c.GoResW(), &cookie)
	fmt.Println("cooooo")
}

func ReadCookieServer(c *odserver.Context) {
	cookie, err := c.GoReq().Cookie(setCookieName)
	if err == nil {
		cookieValue := cookie.Value
		userInfo, err := jwt.JwtDecode(cookieValue)
		if !err {
			fmt.Println("<b>cookie的值是：" + userInfo.UserName + "<b/>\n")
		}

	} else {
		fmt.Println("<b>读取错误" + err.Error() + "</b>\n")
	}
}

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
	if result == "" {
		resp = Resp{
			RetCode: "0",
			Msg:     "用户名或密码不对",
			Data:    "",
		}
		return
	}
	WriteCookieServer(c, result)
}
