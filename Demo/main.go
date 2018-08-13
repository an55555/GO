package main

import (
	"./httpRouter"
	"./utils/jwt"
	"fmt"
	"net/http"
	"time"
)

const tokenCookieName = "tokenTest"

func HelloServer(w http.ResponseWriter, req *http.Request) {
	expiration := time.Now()
	fmt.Println("之前时间%v", expiration)
	//mm, _ := time.ParseDuration("1m")
	//expiration = expiration.Add(mm)
	fmt.Println("时间%v", expiration)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: time.Now().AddDate(0, 0, 1)}
	http.SetCookie(w, &cookie)

	fmt.Println("SetCookie")
}

func HelloServer4(c *odserver.Context) {

	fmt.Fprint(c.Rw, "hello world HelloServer4")
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
	route := odserver.Default()

	route.SetStaticPath("/static/", "static")

	route.Target("/").Get(SayHello)

	route.Start("/{test}/main/").Target("/number/{number}").
		Get(SayHello).Post(SayHello)

	route.Start("/cookie").
		Target("/read").Get(ReadCookieServer).And().
		Target("/write").Get(WriteCookieServer).And().
		Target("/delete").Get(DeleteCookieServer)
	http.ListenAndServe(":6543", route)

}
