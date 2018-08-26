package odserver

import (
	"regexp"
	"strings"
)

const (
	DefaultPathSeparator = "/"
)

type PathMatcher interface {
	IsPattern(path string) bool
	//将包含占位符的路径转换成正则表达式
	ToPattern(path string) (string, []string, bool)
	Match(pattern, path string) (bool, error)
}

type AntPathMatcher struct {
	pathSeparator string
}

func NewAntPathMatcher() AntPathMatcher {
	return AntPathMatcher{pathSeparator: DefaultPathSeparator}
}

func (matcher AntPathMatcher) IsPattern(path string) bool {
	return strings.IndexAny(path, "*") != -1 || strings.IndexAny(path, "?") != -1 || strings.IndexAny(path, "$") != -1
}
func (matcher AntPathMatcher) ToPattern(path string) (string, []string, bool) {
	re := regexp.MustCompile("\\{\\w*}")
	findParamsNames := re.FindAllString(path, -1)
	s := re.ReplaceAllString(path, "\\w*")
	if s == path {
		return s, nil, false
	}
	s += "$"
	return s, findParamsNames, true
}

func (matcher AntPathMatcher) Match(pattern, path string) (bool, error) {
	return regexp.MatchString(pattern, path)
}
func (matcher AntPathMatcher) tokenizePath(path string) []string {
	tokenized := strings.Split(path, matcher.pathSeparator)
	for i, _ := range tokenized {
		tokenized[i] = strings.TrimSpace(tokenized[i])
	}
	return tokenized

}
