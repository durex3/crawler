package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const host = "http://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-9/"

func main() {
	resp, err := http.Get(host)
	if err != nil {
		panic(err)
	}
	if resp == nil || resp.Body == nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error: status code", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	printCarModelList(all)
}

func printCarModelList(contents []byte) {
	var carListRe = regexp.MustCompile(`<a href="(//newcar.xcar.com.cn/car/[\d+-]+\d+/)" onclick="" rel="nofollow" class="page_down"`)
	matches := carListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", string(m[1]))
	}
}
