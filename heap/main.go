package main

import "fmt"

//MaxHeap struct has a slice that holds the array
type MaxHeap struct {
    array []int
}

// Insert ddds an element to the heap.
func (h *MaxHeap) Insert(key int) {
    // append to last index
    h.array = append(h.array, key)
    //heapify
    h.maxHeapifyUp(len(h.array) - 1)
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
    for h.array[parent(index)] < h.array[index] {
        h.swap(parent(index), index)
        index = parent(index)
    }
}
func (h *MaxHeap) swap(i1, i2 int) {
    h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// get the parent index
func parent(i int) int {
    return (i - 1) / 2
}

// left node of the tree
func left(i int) int {
    return 2*i + 1
}

// right node of the tree
func right(i int) int {
    return 2*i + 2
}

//maxHeapifyDown will hepaify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
    lastIndex := len(h.array) - 1
    l, r := left(index), right(index)
    childToCompare := 0
    // loop while index has at least one child
    for l <= lastIndex {
        // when left child  is the only child
        if l == lastIndex {
            childToCompare = l
        } else if h.array[l] > h.array[r] {
            // when left child is larger
            childToCompare = l
        } else {
            // when right child is larger
            childToCompare = r
        }
        // compare array value of current index to larger child and swap if sammler
        if h.array[index] < h.array[childToCompare] {
            h.swap(index, childToCompare)
            index = childToCompare
            l, r = left(index), right(index)
        } else {
            // means that the number found its right place and we have to stop the loop.
            return
        }
    }
}

// Extract returns the la rgest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
    extracted := h.array[0]
    if len(h.array) == 0 {
        fmt.Println("cannot extract because array length is 0")
        return -1
    }
    // take out the last index and put it on the root
    lastIndex := len(h.array) - 1
    h.array[0] = h.array[lastIndex]
    h.array = h.array[:lastIndex]
    h.maxHeapifyDown(0)
    return extracted
}
func main() {
    m := &MaxHeap{}
    buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
    for _, v := range buildHeap {
        m.Insert(v)
        fmt.Println(m)
    }
    for i := 0; i < 5; i++ {
        m.Extract()
        fmt.Println(m)
    }
}
