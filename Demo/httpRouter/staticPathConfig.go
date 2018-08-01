package odserver

import (
	"strings"
)

type regexpMap []string

type StaticPathConfig struct {
	regexpMap
}

func newStaticPathConfig() *StaticPathConfig {
	return &StaticPathConfig{
		regexpMap: make]string)
	}
}

func (s StaticPathConfig) setStaticPath(url string) {
	append(s.regexpMap, url)
}