package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var added bool

	if t.Left != nil {
		Walk(t.Left, ch)
	} else {
		ch <- t.Value
		added = true
	}

	if t.Right != nil {
		if !added {
			ch <- t.Value
		}
		Walk(t.Right, ch)
	} else {
		if !added {
			ch <- t.Value
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(t1)
	fmt.Println(t2)

	z := Same(t1, t2)
	fmt.Println(z)
}
