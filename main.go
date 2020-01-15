package main

import (
	"crawler/engine"
	"crawler/parse/xcar"
	"fmt"
	"regexp"
	"strings"
)

const host = "http://newcar.xcar.com.cn"

func main() {
	engine.Run(engine.Request{
		Url:        host + "/car/0-0-0-0-0-0-0-0-0-0-0-1",
		ParserFunc: xcar.ParseCarList,
	})

	//body, _ := fetcher.Fetch("http://newcar.xcar.com.cn/m48061/config.htm")
	//printCarModelList(body)
}

func printCarModelList(contents []byte) {
	// 级别
	var nameRe = regexp.MustCompile(`<td id="m_dynamic_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)
	matches := nameRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s%v\n", string(m[1]), strings.Fields(string(m[1]))[0])
	}
}
