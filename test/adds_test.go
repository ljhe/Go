package main

import (
	"testing"
)

/**
 * 单元测试 方法名前缀必须为 Test
 * 命令行运行测试方法 Grammar/test 目录下: go test .
 */
func TestAdd(t *testing.T)  {
	tests := []struct{a, b, c int}{
		{1, 2, 3},
		{4, 5, 9},
		{10, 20, 30},
		{100000, 200000, 300000},
	}

	for _, test := range tests {
		if actual := add(test.a, test.b); actual != test.c {
			t.Errorf("add(%d, %d); got %d; expected %d", test.a, test.b, actual, test.c)
		}
	}
}

/**
 * 性能测试 方法名前缀必须为 Benchmark
 * 命令行运行测试方法 Grammar/test 目录下: go test -bench .
 */
func BenchmarkAdd(b *testing.B)  {
	x, y, z := 1, 2, 3
	actual := add(x, y)
	for i := 0; i < b.N; i++ {
		if actual != z {
			b.Errorf("add(%d, %d); got %d; expected %d", x, y, actual, z)
		}
	}
}
