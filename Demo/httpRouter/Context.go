package odserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Context struct {
	Req Request
	Rw  responseWriter
	//对应restful的参数值
	Params     []string
	ParamsName []string
}

func NewContext(r *http.Request, rw http.ResponseWriter) *Context {
	return &Context{
		Req: NewRequest(r),
		Rw:  NewResponse(rw),
	}
}

//获取http包的Request
func (c *Context) GoReq() *http.Request {
	return c.Req.Request
}

//获取http包的ResponseWriter
func (c *Context) GoResW() http.ResponseWriter {
	return c.Rw.ResponseWriter
}

func (c *Context) GetParams() map[string]string {
	paramsMap := make(map[string]string)
	for k, v := range c.ParamsName {
		paramsMap[v] = c.Params[k]
	}
	return paramsMap
}

// 可获取formData和application/json形式的数据提交
func (c *Context) PostParams() map[string]interface{} {
	jsonData := make(map[string]interface{})
	contentType := c.GoReq().Header.Get("Content-Type")
	if contentType == "application/json" {
		body, _ := ioutil.ReadAll(c.GoReq().Body)
		json.Unmarshal(body, &jsonData)
		return jsonData
	}
	c.GoReq().ParseForm()
	for i, k := range c.GoReq().PostForm {
		if len(c.GoReq().PostForm[i]) > 0 {
			jsonData[i] = k[0]
		}
	}
	return jsonData
}
