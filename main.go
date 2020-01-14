package main

import (
	"crawler/fetcher"
	"fmt"
	"regexp"
)

const host = "http://newcar.xcar.com.cn"

func main() {
	/*engine.Run(engine.Request {
		Url: host + "/car/0-0-0-0-0-0-0-0-0-0-0-1",
		ParserFunc: xcar.ParseCarList,
	})*/

	body, _ := fetcher.Fetch("http://newcar.xcar.com.cn/m44278/")
	printCarModelList(body)
}

func printCarModelList(contents []byte) {
	// 级别
	var classRe = regexp.MustCompile(`<td class="w237">(.*)</td>`)
	matches := classRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", string(m[1]))
	}
}
