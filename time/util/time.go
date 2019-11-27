package util

import "time"

const layout = "2006-01-02 15:04:05"

// 时间戳转化为日期
func TimeToDate(times int64) string {
	return time.Unix(times, 0).Format(layout)
}

// 日期转化为时间戳
func DateToTime(date string) (int64, error) {
	// 获取 location
	location, e := time.LoadLocation("Local")
	if e != nil {
		return 0, e
	}

	// 使用模板在对应时区转化为 time.time 类型
	times, e := time.ParseInLocation(layout, date, location)
	if e != nil {
		return 0, e
	}
	return times.Unix(), nil
}