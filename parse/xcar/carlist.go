package xcar

import (
	"crawler/engine"
	"regexp"
)

const host = "http://newcar.xcar.com.cn"

// 汽车模型
var carModelRe = regexp.MustCompile(`<a href="(/\d+/)" target="_blank" class="list_img">`)

// 每一页的汽车模型列表
var carListRe = regexp.MustCompile(`<a href="(//newcar.xcar.com.cn/car/[\d+-]+\d+/)" onclick="" rel="nofollow" class="page_down"`)

func ParseCarList(contents []byte) engine.ParseResult {
	matches := carModelRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        host + string(m[1]),
				ParserFunc: ParseCarModel,
			})

	}

	matches = carListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        "http:" + string(m[1]),
				ParserFunc: ParseCarList,
			})
	}

	return result
}
