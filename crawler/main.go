package main

import (
	"Grammar/crawler/engine"
	"Grammar/crawler/zhenai/parser"
)

func main()  {
<<<<<<< HEAD
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		fmt.Println("something error")
		return
	}
	defer func() {
		e := resp.Body.Close()
		if e != nil {
			fmt.Println("Body close error")
		}
	}()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp.StatusCode is ", resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error")
		return
	}
	printCityList(all)
}

func printCityList(contents []byte) {
	// 正则表达式
	// `.+` 中 `.` 表示任意字符 `+` 表示有一个或者多个
	// `.*` 中 `.` 表示任意字符 `*` 表示有零个或者多个
	// `[0-9a-zA-Z]+` 中 `[0-9a-zA-Z]` 表示字符的规则
	// `[^>]` 中 `[^>]` 表示不为 `>` 字符的所有字符
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)" [^>]*>([^<]+)</a>`)
	all := re.FindAllSubmatch(contents, -1)
	for _, a := range all {
		fmt.Printf("City: %s, URL: %s\n", a[2], a[1])
	}
=======
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
>>>>>>> 829e3538edba3869675c27c63d1d9144f4d94e78
}