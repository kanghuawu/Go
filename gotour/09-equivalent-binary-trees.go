package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    defer close(ch)
    WalkHelper(t, ch)
}
func WalkHelper(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    
    WalkHelper(t.Left, ch)
    ch <- t.Value
    WalkHelper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    
    for num1 := range ch1 {
        if num1 == <- ch2 {
            continue
        } else {
            return false
        }
    }
    _, ok := <- ch2
    return ok == false
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}