package odserver

import "net/http"

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
