package xcar

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

// 车名
var nameRe = regexp.MustCompile(`title="(.*)参数配置"`)

// 车的图片
var carImageRe = regexp.MustCompile(`<img class="color_car_img_new" src="([^"]+)"`)

// 车的基本参数和车身参数
var carInfoRe = regexp.MustCompile(`<td class="w237">(.*)</td>`)

func ParseCarDetail(contents []byte, url string) engine.ParseResult {
	matches := carInfoRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", string(m[1]))
	}
}
