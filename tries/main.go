package main

import "fmt"

// AlphabetSize is number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
    children [AlphabetSize]*Node
    isEnd    bool
}

// Trie represent a trie and has  a pointer to the root node
type Trie struct {
    root *Node
}

// NewTrie is used to create a new trie.
func NewTrie() *Trie {
    return &Trie{root: &Node{}}
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
    wordLength := len(w)
    currentNode := t.root
    for i := 0; i < wordLength; i++ {
        charIndex := w[i] - 'a'
        if currentNode.children[charIndex] == nil {
            currentNode.children[charIndex] = &Node{}
        }
        currentNode = currentNode.children[charIndex]
    }
    currentNode.isEnd = true
}

// Search will take in a word and return true if that word is included in the trie
func (t *Trie) Search(w string) bool {
    wordLength := len(w)
    currentNode := t.root
    for i := 0; i < wordLength; i++ {
        charIndex := w[i] - 'a'
        if currentNode.children[charIndex] == nil {
            return false
        }
        currentNode = currentNode.children[charIndex]
    }
    return currentNode.isEnd
}
func main() {
    testTrie := NewTrie()
    toAdd := []string{
        "aragorn",
        "aragon",
        "argon",
        "eragon",
        "oregon",
        "oregano",
        "oreo",
    }
    for _, val := range toAdd {
        testTrie.Insert(val)
    }
    fmt.Println(testTrie.Search("eragon"))
}
