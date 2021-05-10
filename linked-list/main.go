package main

import "fmt"

type node struct {
    data int
    next *node
}

type linkedList struct {
    head   *node
    length int
}

func (l *linkedList) prepend(n *node) {
    second := l.head
    l.head = n
    l.head.next = second
    l.length++
}
func (l linkedList) printListData() {
    toPrint := l.head
    for l.length != 0 {
        fmt.Printf("%d ", toPrint.data)
        toPrint = toPrint.next
        l.length--
    }
    fmt.Printf("\n")
}
func (l *linkedList) deleteWithValue(value int) {
    // empty linkedlist
    if l.length == 0 {
        return
    }
    //deleting header of the linkedlist.
    if l.head.data == value {
        l.head = l.head.next
        l.length--
        return
    }
    toDelete := l.head
    for toDelete.next.data != value {
        if toDelete.next.next == nil {
            return
        }
        toDelete = toDelete.next
    }
    toDelete.next = toDelete.next.next
    l.length--
}
func main() {
    list := linkedList{}
    n1 := &node{data: 48}
    n2 := &node{data: 38}
    n3 := &node{data: 68}
    n4 := &node{data: 368}
    n5 := &node{data: 618}
    n6 := &node{data: 680}
    n7 := &node{data: 18}
    list.prepend(n1)
    list.prepend(n2)
    list.prepend(n3)
    list.prepend(n4)
    list.prepend(n5)
    list.prepend(n6)
    list.prepend(n7)
    list.printListData()
    list.deleteWithValue(18)
    list.printListData()

}
