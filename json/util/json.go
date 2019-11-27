package util

import (
	"encoding/json"
)

// map 转成 json 格式
func MapOrStructToJson(m interface{}) string {
	bytes, err := json.Marshal(m)
	if err != nil {
		err.Error()
	}
	return string(bytes)
}

// json 转化成 map
func JsonToMap(j string) (v interface{}) {
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		err.Error()
	}
	return
}
