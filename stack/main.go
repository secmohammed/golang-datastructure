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

type Queue struct {
    items []int
}

func (q *Queue) Enqueue(i int) {
    q.items = append(q.items, i)
}
func (q *Queue) Dequeue() int {
    item := q.items[0]
    q.items = q.items[1:]
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

    queue := Queue{}
    fmt.Println(queue)
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    fmt.Println(queue)
    queue.Dequeue()
    fmt.Println(queue)

}
