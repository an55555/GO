package main

import (
	"Golang-WEB/Demo/httpRouter"
	"fmt"
	"net/http"
	"time"
)

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
	cookie, err := req.Cookie("testCookieName")
	if err == nil {
		cookieValue := cookie.Value
		w.Write([]byte("<b>cookie的值是：" + cookieValue + "<b/>\n"))
	} else {
		w.Write([]byte("<b>读取错误" + err.Error() + "</b>\n"))
	}
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	nowTime := time.Now()
	fmt.Println("当时时间%v", nowTime)
	addTime, _ := time.ParseDuration("1m")
	nowTime = nowTime.Add(addTime)
	fmt.Println("增加后的时间%v", nowTime)
	cookie := http.Cookie{Name: "testCookieName", Value: "testCookieValue", Expires: nowTime, Path: "/"}
	http.SetCookie(w, &cookie)
	w.Write([]byte("<b>设置cookie成功。</b>\n"))
}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "testCookieName", MaxAge: -1}
	http.SetCookie(w, &cookie)
	w.Write([]byte("<b>删除cookie成功。</b>\n"))
}

func main() {
	o := odserver.Default()

	o.SetStaticPath("/static/", "static")

	o.Target("/").Get(SayHello)

	o.Start("/{test}/main/").Target("/number/{number}").
		Get(SayHello).Post(SayHello)

	o.Start("/cookie").
		Target("/read").Get(ReadCookieServer).And().
		Target("/write").Get(WriteCookieServer).And().
		Target("/delete").Get(DeleteCookieServer)
	http.ListenAndServe(":6543", o)

}
