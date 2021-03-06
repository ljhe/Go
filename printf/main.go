package main

/**
 * 常见输出格式
 */

import (
	"fmt"
)

type point struct {
	x, y int
}

func main()  {
	p := point{
		x: 1,
		y: 2,
	}

	// 打印了一个结构体的实例
	// 输出结果: {1 2}
	fmt.Printf("%v\n", p)

	// 如果值是一个结构体 %+v 的格式化输出内容将包括结构体的字段名
	// 输出结果: {x:1 y:2}
	fmt.Printf("%+v\n", p)

	// %#v 形式则输出这个值的 Go 语法表示 例如,值的运行源代码片段
	// 输出结果: main.point{x:1, y:2}
	fmt.Printf("%#v\n", p)

	// 需要打印值的类型, 使用 %T
	// 输出结果: main.point
	fmt.Printf("%T\n", p)

	// 格式化布尔值是简单的
	// 输出结果: true
	fmt.Printf("%t\n", true)

	// 格式化整形数有多种方式, 使用 %d进行标准的十进制格式化
	// 输出结果: 123
	fmt.Printf("%d\n", 123)

	// 当输出数字的时候, 你将经常想要控制输出结果的宽度和精度, 可以使用在 % 后面使用数字来控制输出宽度 .
	// 默认结果使用右对齐并且通过空格来填充空白部分。
	// 左对齐使用 "-"
	// 输出结果: |    12|   345|
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%-6d|%-6d|\n", 12, 345)

	// 输出二进制表示形式
	// 输出结果: 1110
	fmt.Printf("%b\n", 14)

	// 输出给定整数的对应字符
	// 输出结果: !
	fmt.Printf("%c\n", 33)

	// 输出十六进制编码对应字符
	// 输出结果: 1c8
	fmt.Printf("%x\n", 456)

	// 对于浮点型同样有很多的格式化选项 使用 %f 进行最基本的十进制格式化
	// 输出结果: 78.900000
	fmt.Printf("%f\n", 78.9)

	// 指定浮点型的输出宽度, 同时也可以通过 宽度 精度 的语法来指定输出的精度
	// 输出结果: |  1.20| 3.450|
	// 6 表示宽度 .2 和 .3 小数点后面的数字表示精度
	fmt.Printf("|%6.2f|%6.3f|\n", 1.2, 3.45)

	// %e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式
	// 输出结果: 1.234000e+08
	// 输出结果: 1.234000E+08
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// 使用 %s 进行基本的字符串输出
	// 字符串也可以像整数那样 控制输出时的宽度 以及对齐方式
	// 输出结果: "string"
	// 输出结果: |   foo|b     |
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("|%6s|%-6s|\n", "foo", "b")

	// 像 Go 源代码中那样带有双引号的输出
	// 输出结果: "\"string\""
	fmt.Printf("%q\n", "\"string\"")

	// 输出一个指针的值
	// 输出结果: 0xc000062090
	fmt.Printf("%p\n", &p)
}
