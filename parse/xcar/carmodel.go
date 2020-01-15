package xcar

import (
	"crawler/engine"
	"regexp"
)

var carDetailRe = regexp.MustCompile(`<a href="(/m\d+/)" target="_blank"`)

func ParseCarModel(contents []byte) engine.ParseResult {
	matches := carDetailRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        host + string(m[1]) + "/config.htm",
				ParserFunc: ParseCarDetail,
			})
	}
	return result
}
