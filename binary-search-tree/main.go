package main

import "fmt"

var count int

// Node represents the components of a binary search tree
type Node struct {
    Key   int
    Left  *Node
    Right *Node
}

// Insert will add a node to the tree
// the key to addd shouldn't be already in the tree
func (n *Node) Insert(k int) {
    if n.Key < k {
        // move right
        if n.Right == nil {
            n.Right = &Node{Key: k}
        } else {
            n.Right.Insert(k)
        }
    } else if n.Key > k {
        // move left
        if n.Left == nil {
            n.Left = &Node{Key: k}
        } else {
            n.Left.Insert(k)
        }

    }
}

// Search will take in a key value
// and return true if there is a node with that value
func (n *Node) Search(k int) bool {
    count++
    // the end of the tree and couldn't find the match.
    if n == nil {
        return false
    }
    if n.Key < k {
        // move right
        return n.Right.Search(k)
    } else if n.Key > k {
        // move left
        n.Left.Search(k)
    }
    return true
}
func main() {
    tree := &Node{Key: 100}
    tree.Insert(200)
    tree.Insert(300)
    tree.Insert(130)
    tree.Insert(45)
    tree.Insert(39)
    tree.Insert(294)
    fmt.Println(tree)
    fmt.Println(tree.Search(200), tree.Search(1000))
}
