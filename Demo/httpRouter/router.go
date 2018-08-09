package odserver

import (
	"regexp"
	"strings"
	"time"
)

/**
提供基本的路由功能，添加路由，查找路由
*/
const (
	GET = iota
	POST
	PUT
	DELETE
	CONNECTIBNG
	HEAD
	OPTIONS
	PATCH
	TRACE
	paramsSize = 6
)

var matcher PathMatcher = NewAntPathMatcher()

type handler map[string]*HandlerObject
type regexpMap map[*regexp.Regexp]*HandlerObject

func NewRouter() *Router {
	return &Router{
		handler:   make(map[string]*HandlerObject),
		regexpMap: make(map[*regexp.Regexp]*HandlerObject),
	}
}

type Router struct {
	handler
	regexpMap
}

func (r *Router) Target(url string) *HandlerObject {
	HandlerObject := NewHandlerObject(r, AddSlash("/"))
	HandlerObject.Target(url)
	return HandlerObject
}

func (r *Router) Start(url string) *HandlerObject {
	return NewHandlerObject(r, AddSlash(url))
}

func (ho *HandlerObject) And() *HandlerObject {
	if ho.Router == nil || ho.startPath == "" {
		panic("ho.Router is nil or startPath is unknown，maybe u should use Start()")
	}
	return NewHandlerObject(ho.Router, ho.startPath)
}

func (ho *HandlerObject) Target(url string) *HandlerObject {
	//设置完整的路径
	if ho.startPath == "/" {
		ho.path = ho.startPath + DeleteSlash(url)
	} else {
		if strings.HasSuffix(ho.startPath, "/") {
			url = DeleteSlash(url)
		} else {
			url = AddSlash(url)
		}
		ho.path = ho.startPath + url
	}
	//尝试将url转换成正则表达式，如果没有占位符，则转换不成功
	pattern, ok := matcher.ToPattern(ho.path)
	if ok {
		ho.path = pattern
		re, err := regexp.Compile(pattern)
		if err != nil {
			panic("error compile pattern:" + pattern)
		}
		ho.Router.regexpMap[re] = ho
	} else {
		ho.handler[ho.path] = ho
	}
	return ho
}
func AddSlash(s string) string {
	if !strings.HasPrefix(s, "/") {
		s = "/" + s
	}
	return s
}

func DeleteSlash(s string) string {
	if strings.HasPrefix(s, "/") {
		array := strings.SplitN(s, "/", 2)
		s = array[1]
	}
	return s
}

//匹配路径
func (r *Router) doUrlMapping(url string, method int) (FuncObject, bool) {
	ch := make(chan *HandlerObject)
	//精准匹配
	go func() {
		if ho, ok := r.handler[url]; ok {
			ch <- ho
		}
	}()
	//正则匹配
	go func() {
		for k, v := range r.regexpMap {
			if k.MatchString(url) {
				pathArray := strings.Split(url, "/")[1:]
				regexpArray := strings.Split(k.String(), "/")[1:]
				if len(pathArray) == len(regexpArray) {
					//设置参数
					paramsNum := 0
					for i := 0; i < len(pathArray); i++ {
						if matcher.IsPattern(regexpArray[i]) {
							v.params[paramsNum] = pathArray[i]
							paramsNum++
						}
					}
					v.params = v.params[:paramsNum]
				}
				ch <- v
			}
		}
	}()
	select {
	case ho := <-ch:
		{
			return ho.Func(method)
		}
	case <-time.After(2e6):
		{
			return FuncObject{}, false
		}
	}

}
