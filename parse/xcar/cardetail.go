package xcar

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 车名
var nameRe = regexp.MustCompile(`<h1>(.*) 参数配置</h1>`)

// 车的价格
var priceRe = regexp.MustCompile(`<td id="price_\d+" class="bg4">[\s]*(.*)万[\s]*</td>`)

// 厂商
var manufacturerRe = regexp.MustCompile(`<td id="bname_\d+" class="bg4">[\s]*<a target="_blank" href="/price/.*">(.*)</a>[\s]*</td>`)

// 级别
var classRe = regexp.MustCompile(`<td id="type_name_\d+" class="bg4">[\s]*<a target="_blank" href="/car/.*">(.*)</a>[\s]*</td>`)

// 发动机
var engineRe = regexp.MustCompile(`<td id="m_motortype_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 动力类型
var powerType = regexp.MustCompile(`<td id="m_dynamic_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 变速箱
var transmission = regexp.MustCompile(`<td id="m_speed_transtype_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 长(mm)
var length = regexp.MustCompile(`<td id="m_length_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 宽(mm)
var width = regexp.MustCompile(`<td id="m_width_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 高(mm)
var height = regexp.MustCompile(`<td id="m_height_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 车身结构
var structure = regexp.MustCompile(`<td id="m_door_seat_frame_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 上市年份
var year = regexp.MustCompile(`<td id="syear_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 最高车速(km/h)
var maxSpeed = regexp.MustCompile(`<td id="m_mspeed_\d+" class="bg4">[\s]*(.*)[\s]*</td>`)

// 提取id
var urlRe = regexp.MustCompile(`http://newcar.xcar.com.cn/(m\d+)/`)

func ParseCarDetail(contents []byte) engine.ParseResult {
	// id := extractString([]byte(url), urlRe)
	car := model.Car{
		Name:         extractString(contents, nameRe),
		Price:        extractFloat(contents, priceRe),
		Manufacturer: extractString(contents, manufacturerRe),
		Class:        extractString(contents, classRe),
		PowerType:    extractString(contents, powerType),
		Transmission: extractString(contents, transmission),
		Length:       extractInt(contents, length),
		Width:        extractInt(contents, width),
		Height:       extractInt(contents, height),
		Structure:    extractString(contents, structure),
		Year:         extractInt(contents, year),
		MaxSpeed:     extractInt(contents, maxSpeed),
	}
	fmt.Printf("Got item: %v\n", car)
	result := engine.ParseResult{
		Items: []interface{}{car},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return strings.Fields(string(match[1]))[0]
	} else {
		return ""
	}
}

func extractFloat(contents []byte, re *regexp.Regexp) float64 {
	f, err := strconv.ParseFloat(extractString(contents, re), 64)
	if err != nil {
		return 0
	}
	return f
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	i, err := strconv.Atoi(extractString(contents, re))
	if err != nil {
		return 0
	}
	return i
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
