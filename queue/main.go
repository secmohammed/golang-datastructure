package main

import "fmt"

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
    queue := Queue{}
    fmt.Println(queue)
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    fmt.Println(queue)
    queue.Dequeue()
    fmt.Println(queue)

}
