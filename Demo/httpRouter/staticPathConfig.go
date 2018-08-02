package odserver

type staticUrlMap []string

type StaticPathConfig struct {
	staticUrlMap
}

func NewStaticPathConfig() *StaticPathConfig {
	return &StaticPathConfig{
		staticUrlMap: make([]string, 0),
	}
}

func (s *StaticPathConfig) SetStaticPath(url string) {
	s.staticUrlMap = append(s.staticUrlMap, url)
}
