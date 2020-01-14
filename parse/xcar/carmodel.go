package xcar

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

var carDetailRe = regexp.MustCompile(`<a href="(/m\d+/)" target="_blank"`)

func ParseCarModel(contents []byte) engine.ParseResult {
	matches := carDetailRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		fmt.Println(string(m[1]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        host + string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result
}
