package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func main() {
	stack := Stack{}
	fmt.Println(stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
