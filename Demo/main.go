package main

import (
	"./httpRouter"
	"./utils/encrypt"
	"./utils/jwt"
	"fmt"
	"net/http"
	"time"
)

const tokenCookieName = "tokenTest"

func GetParams(c *odserver.Context) {
	fmt.Println(c.Params)
	//fmt.Fprintf(w, "Hello astaxie!")
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func ReadCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(tokenCookieName)
	if err == nil {
		cookieValue := cookie.Value
		userInfo, err := jwt.JwtDecode(cookieValue)
		if !err {
			w.Write([]byte("<b>cookie的值是：" + userInfo.UserName + "<b/>\n"))
		}

	} else {
		w.Write([]byte("<b>读取错误" + err.Error() + "</b>\n"))
	}
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	nowTime := time.Now()
	fmt.Println("当时时间%v", nowTime)
	addTime, _ := time.ParseDuration("10m")
	nowTime = nowTime.Add(addTime)
	fmt.Println("增加后的时间%v", nowTime)
	tokenPayload := jwt.NewPayload("lanTest")
	tokenTest := jwt.JwtCode(tokenPayload)
	cookie := http.Cookie{Name: tokenCookieName, Value: tokenTest, Expires: nowTime, Path: "/"}
	http.SetCookie(w, &cookie)
	w.Write([]byte("<b>设置cookie成功。</b>\n"))
}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: tokenCookieName, MaxAge: -1}
	http.SetCookie(w, &cookie)
	w.Write([]byte("<b>删除cookie成功。</b>\n"))
}

func main() {
	encryptData := encrypt.EncryptSailt("我是谁")
	fmt.Println("%x", encryptData)

	route := odserver.Default()

	route.SetStaticPath("/static/", "static")

	route.Target("/").GoGet(SayHello)
	route.Get("/get", SayHello).Get("/get2", SayHello)
	route.Start("/new").Get("/1", SayHello).Get("/2", SayHello)
	route.Target("/params/{id}").GoGet(GetParams)

	route.Start("/{test}/main/").Target("/number/{number}").
		GoGet(SayHello).GoPost(SayHello)

	route.Start("/cookie").
		Target("/read").GoGet(ReadCookieServer).And().
		Target("/write").GoGet(WriteCookieServer).And().
		Target("/delete").GoGet(DeleteCookieServer)
	http.ListenAndServe(":6543", route)

}
