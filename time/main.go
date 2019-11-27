package main

import (
	"Grammar/time/util"
	"fmt"
	"time"
)

func main()  {
	// 获取当前时间戳
	a := time.Now().Unix()
	fmt.Println(a)

	// 时间戳转化为日期
	b := util.TimeToDate(a)
	fmt.Println(b)

	// 日期转化为时间戳
	c, err := util.DateToTime(b)
	fmt.Println(c, err)
}
