package main

import (
	"fmt"
)

// Stack ...
type Stack struct {
	el  []interface{}
	top int
}

// NewSequceStack 新建栈
func NewSequceStack() *Stack {
	return new(Stack)
}

// Push 入栈
func (s *Stack) Push(el interface{}) {
	s.top++
	s.el = append(s.el, el)
}

// Pop 出栈
func (s *Stack) Pop() (ok bool, el interface{}) {
	if s.top < 1 {
		return
	}
	s.top--
	el = s.el[s.top]
	s.el = s.el[:s.top]
	return true, el
}

//PrintStack 打印栈中的所有元素
func (s *Stack) PrintStack() {
	fmt.Println(s.el)
}

func main() {
	stack := NewSequceStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.PrintStack()
	stack.Pop()
	stack.PrintStack()
}