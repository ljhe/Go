package main

import "fmt"

type stack struct {
	val []rune
	t   map[string]rune
}

// 这里如果不是 * 的话, 在外面修改不了 s.val 的值
func (s *stack) pop() rune {
	r := s.val[len(s.val) - 1]
	s.val = s.val[0 : len(s.val) - 1]
	return r
}

func isValid(s string) bool {
	flag := true
	var stack stack
	stack.t = map[string]rune{
		")": '(',
		"]": '[',
		"}": '{',
	}
	length := len(s)
	if length == 0 {
		return flag
	}
	for _, i := range []rune(s) {
		if len(stack.val) == 0 || string(i) == "(" || string(i) == "[" || string(i) == "{" {
			stack.val = append(stack.val, i)
		} else {
			top := stack.pop()
			if top == stack.t[string(i)] {
				flag = true
			} else {
				flag = false
				return flag
			}
		}
	}
	if len(stack.val) != 0 {
		flag = false
	}
	return flag
}

func main() {
	s := "{{)}"
	fmt.Println(isValid(s))
}
