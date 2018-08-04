package odserver

import (
	"net/http"
	"strings"
)

type staticUrlMap map[string]string

type StaticPathConfig struct {
	staticUrlMap
}

func NewStaticPathConfig() *StaticPathConfig {
	return &StaticPathConfig{
		staticUrlMap: make(map[string]string),
	}
}

func (s *StaticPathConfig) SetStaticPath(url string, path string) {
	s.staticUrlMap[url] = path
}

func (s *StaticPathConfig) MapStaticPath(url string) (bool, string, string) {
	staticMap := s.staticUrlMap
	for rangeUrl, staticPath := range staticMap {
		if strings.Index(url, rangeUrl) == 0 {
			return true, rangeUrl, staticPath
		}
	}
	return false, "", ""
}

func (s *StaticPathConfig) doStaticPath(w http.ResponseWriter, req *http.Request, url string, path string) {
	fs := http.FileServer(http.Dir(path))
	staticFile := http.StripPrefix(url, fs)
	staticFile.ServeHTTP(w, req)
}
