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
func HelloServer2(w http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("username")
	fmt.Fprint(w, cookie)
}

func HelloServer3(c *odserver.Context) {

	fmt.Fprint(c.Rw, c.Params)
}
func HelloServer4(c *odserver.Context) {

	fmt.Fprint(c.Rw, "hello world HelloServer4")
}

func main() {
	o := odserver.Default()
	o.SetStaticPath("/static/", "static")
	o.Start("/main").
		Target("/test/").Get(HelloServer).Post(HelloServer).Delete(HelloServer).And().
		Target("/test2").Get(HelloServer2)
	o.Start("/{test}/main/").Target("/number/{number}").
		Get(HelloServer3).Post(HelloServer4)

	http.ListenAndServe(":8080", o)

}
