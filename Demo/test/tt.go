package main

import (
	"fmt"
)

type staticUrlMap []string

type StaticPathConfig struct {
	staticUrlMap
}

func newStaticPathConfig() *StaticPathConfig {
	return &StaticPathConfig{
		staticUrlMap: make([]string, 0),
	}
}

func (s *StaticPathConfig) setStaticPath(url string) {
	s.staticUrlMap = append(s.staticUrlMap, url)
}

func main() {
	s1 := newStaticPathConfig()
	fmt.Println(s1)
	fmt.Println(len(s1.staticUrlMap))
	s1.setStaticPath("abc/abcd")
	fmt.Println(s1)
	fmt.Println(len(s1.staticUrlMap))
}
