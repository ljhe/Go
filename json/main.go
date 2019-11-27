package main

import (
	"Grammar/json/util"
	"encoding/json"
	"fmt"
)

// 字段名小写转不了 json 输出结果为: {}
type Person struct {
	name string
	age int
}

// 字段名大写可以转 json
type Person2 struct {
	Name string
	Age int
}

// 给转成的 json 的 key 设置别名 输出结果: {"name":"li","age":18}
type Person3 struct {
	Name string `json:"name"`
	Age int `json:"age"`
} 

// json 转 struct
func jsonToStruct(j string) (v Person2) {
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		err.Error()
	}
	return
}

func main()  {
	a := map[string]interface{}{
		"a" : "b",
		"m" : "n",
		"x" : "y",
	}

	b := util.MapOrStructToJson(a)
	fmt.Println(b)

	fmt.Println(util.JsonToMap(b))

	// 字段名小写转不了 json 输出结果为: {}
	p := Person{
		name: "li",
		age:  18,
	}
	c := util.MapOrStructToJson(p)
	fmt.Println(c)

	p2 := Person2{
		Name: "li",
		Age:  18,
	}
	c2 := util.MapOrStructToJson(p2)
	fmt.Println(c2)

	p3 := Person3{
		Name: "li",
		Age:  18,
	}
	fmt.Printf("%+v\n", p3)
	c3 := util.MapOrStructToJson(p3)
	fmt.Println(c3)
	fmt.Printf("%+v\n", jsonToStruct(c3))
}
