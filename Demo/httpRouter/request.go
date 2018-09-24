package odserver

import (
	"fmt"
	"net/http"
)

type Request struct {
	*http.Request
	*HandlerObject
}
type FuncObject struct {
	params     []string
	paramsName []string
	//对应编写的接口，IHandlerFunc是个空接口
	f     IHandlerFunc
	exist bool
	*httpConfig
}

func NewFuncObject(f IHandlerFunc) FuncObject {
	return FuncObject{
		f:          f,
		exist:      true,
		httpConfig: &httpConfig{header: make(map[string]string)},
	}
}

//方法函数映射，0代表GET方法下的接口
type methodFuncs []FuncObject

/**
关键struct，代表每个实体的请求
*/
type HandlerObject struct {
	*Router
	//对应占位符的参数
	params []string
	//对应占位符的参数名
	paramsName []string
	//对该请求的http配置
	*httpConfig
	//请求路径 即start+target的路径
	path      string
	startPath string
	//方法函数映射
	methodFuncs methodFuncs
}

//接口函数单位，即我们编写代码逻辑的函数
type IHandlerFunc interface{}

type IHandler interface {
	Get(f IHandlerFunc) *HandlerObject
	Post(f IHandlerFunc) *HandlerObject
	Put(f IHandlerFunc) *HandlerObject
	Delete(f IHandlerFunc) *HandlerObject
}

func NewHandlerObject(r *Router, startPath string) *HandlerObject {
	return &HandlerObject{
		params:      make([]string, paramsSize),
		Router:      r,
		startPath:   startPath,
		methodFuncs: make([]FuncObject, TRACE),
	}
}

func NewRequest(r *http.Request) Request {
	return Request{Request: r}
}

func (ho *HandlerObject) GoGet(f IHandlerFunc) *HandlerObject {
	if ho.methodFuncs[GET].exist {
		panic("GetFunc has existed")
	}
	ho.methodFuncs[GET] = NewFuncObject(f)
	return ho
}
func (ho *HandlerObject) GoPost(f IHandlerFunc) *HandlerObject {
	if ho.methodFuncs[POST].exist {
		panic("PostFunc has existed")
	}
	ho.methodFuncs[POST] = NewFuncObject(f)
	return ho
}

func (ho *HandlerObject) GoPut(f IHandlerFunc) *HandlerObject {
	if ho.methodFuncs[PUT].exist {
		panic("PutFunc has existed")
	}
	ho.methodFuncs[PUT] = NewFuncObject(f)
	return ho
}
func (ho *HandlerObject) GoDelete(f IHandlerFunc) *HandlerObject {
	if ho.methodFuncs[DELETE].exist {
		panic("DeleteFunc has existed")
	}
	ho.methodFuncs[DELETE] = NewFuncObject(f)
	return ho
}

func (ho *HandlerObject) Get(url string, f IHandlerFunc) *HandlerObject {
	ho.Target(url)
	ho.GoGet(f)
	return NewHandlerObject(ho.Router, ho.startPath)
}
func (ho *HandlerObject) Post(url string, f IHandlerFunc) *HandlerObject {
	ho.Target(url)
	ho.GoPost(f)
	return NewHandlerObject(ho.Router, ho.startPath)
}

func (ho *HandlerObject) Put(url string, f IHandlerFunc) *HandlerObject {
	ho.Target(url)
	ho.GoPut(f)
	return NewHandlerObject(ho.Router, ho.startPath)
}
func (ho *HandlerObject) Delete(url string, f IHandlerFunc) *HandlerObject {
	ho.Target(url)
	ho.GoDelete(f)
	return NewHandlerObject(ho.Router, ho.startPath)
}

func (ho *HandlerObject) Func(method int) (FuncObject, bool) {
	switch method {
	case GET:
		return ho.getFunc()
	case DELETE:
		return ho.deleteFunc()
	case PUT:
		return ho.putFunc()
	case POST:
		return ho.postFunc()
	case TRACE:
		return ho.traceFunc()
	case PATCH:
		return ho.patchFunc()
	case OPTIONS:
		return ho.optionsFunc()
	case HEAD:
		return ho.headFunc()
	case CONNECTIBNG:
		return ho.connectingFunc()
	}
	return FuncObject{}, false

}

func (ho *HandlerObject) getFunc() (FuncObject, bool) {
	if ho.methodFuncs[GET].exist {
		ho.methodFuncs[GET].params = ho.params
		ho.methodFuncs[GET].paramsName = ho.paramsName
	}
	return ho.methodFuncs[GET], ho.methodFuncs[GET].exist
}
func (ho *HandlerObject) postFunc() (FuncObject, bool) {
	if ho.methodFuncs[POST].exist {
		ho.methodFuncs[POST].params = ho.params
		ho.methodFuncs[POST].paramsName = ho.paramsName
	}
	return ho.methodFuncs[POST], ho.methodFuncs[POST].exist
}
func (ho *HandlerObject) putFunc() (FuncObject, bool) {
	if ho.methodFuncs[PUT].exist {
		ho.methodFuncs[PUT].params = ho.params
		ho.methodFuncs[PUT].paramsName = ho.paramsName
	}
	return ho.methodFuncs[PUT], ho.methodFuncs[PUT].exist
}

func (ho *HandlerObject) deleteFunc() (FuncObject, bool) {
	if ho.methodFuncs[DELETE].exist {
		ho.methodFuncs[DELETE].params = ho.params
		ho.methodFuncs[DELETE].paramsName = ho.paramsName
	}
	return ho.methodFuncs[DELETE], ho.methodFuncs[DELETE].exist
}

func (ho *HandlerObject) connectingFunc() (FuncObject, bool) {
	if ho.methodFuncs[CONNECTIBNG].exist {
		ho.methodFuncs[CONNECTIBNG].params = ho.params
		ho.methodFuncs[CONNECTIBNG].paramsName = ho.paramsName
	}
	return ho.methodFuncs[CONNECTIBNG], ho.methodFuncs[CONNECTIBNG].exist
}
func (ho *HandlerObject) headFunc() (FuncObject, bool) {
	if ho.methodFuncs[HEAD].exist {
		ho.methodFuncs[HEAD].params = ho.params
		ho.methodFuncs[HEAD].paramsName = ho.paramsName
	}
	return ho.methodFuncs[HEAD], ho.methodFuncs[HEAD].exist
}

func (ho *HandlerObject) optionsFunc() (FuncObject, bool) {
	if ho.methodFuncs[OPTIONS].exist {
		ho.methodFuncs[OPTIONS].params = ho.params
		ho.methodFuncs[OPTIONS].paramsName = ho.paramsName
	}
	return ho.methodFuncs[OPTIONS], ho.methodFuncs[OPTIONS].exist
}

func (ho *HandlerObject) patchFunc() (FuncObject, bool) {
	if ho.methodFuncs[PATCH].exist {
		ho.methodFuncs[PATCH].params = ho.params
		ho.methodFuncs[PATCH].paramsName = ho.paramsName
	}
	return ho.methodFuncs[PATCH], ho.methodFuncs[PATCH].exist
}

func (ho *HandlerObject) traceFunc() (FuncObject, bool) {
	if ho.methodFuncs[TRACE].exist {
		ho.methodFuncs[TRACE].params = ho.params
		ho.methodFuncs[TRACE].paramsName = ho.paramsName
	}
	return ho.methodFuncs[TRACE], ho.methodFuncs[TRACE].exist
}
