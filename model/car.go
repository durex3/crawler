package model

type Car struct {
	Name         string  // 车名
	Price        float64 // 价格
	Manufacturer string  // 厂商
	Class        string  // 级别
	//Engine       string  // 发动机
	PowerType    string // 动力类型
	Transmission string // 变速箱
	Length       int    // 长(mm)
	Width        int    // 宽(mm)
	Height       int    // 高(mm)
	Structure    string // 车身结构
	Year         int    // 上市年份
	MaxSpeed     int    // 最高车速(km/h)
}
