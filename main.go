package main

import (
	"crawler/engine"
	"crawler/parse/xcar"
	"crawler/scheduler"
	"fmt"
	"regexp"
)

const host = "http://newcar.xcar.com.cn"

func main() {
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        host + "/car/0-0-0-0-0-0-0-0-0-0-0-1",
		ParserFunc: xcar.ParseCarList,
	})

	//body, _ := fetcher.Fetch("http://newcar.xcar.com.cn/m48061/config.htm")
	//printCarModelList(body)
}

func printCarModelList(contents []byte) {
	// 级别
	var nameRe = regexp.MustCompile(`<br/>[\s]*\((.*)\)[\s]*</td>`)
	matches := nameRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s", string(m[1]))
	}
}
