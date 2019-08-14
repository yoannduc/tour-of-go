package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Implement a recursive implicit walk function which will be called by Walk
// because Walk need to close when done
func walkImpl(t *tree.Tree, ch, quit chan int) {
	// If t was nil, return, closing the channel in Walk
	if t == nil {
		return
	}

	// Walk left
	walkImpl(t.Left, ch, quit)

	// Switch for channels
	select {
	// Send value to channel
	case ch <- t.Value:
		// Value successfully sent.
	// If quit has value, return, closing the channel in Walk
	case <-quit:
		return
	}

	// Walk right
	walkImpl(t.Right, ch, quit)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// It calls implicit walk function helper & then close the channel when done
func Walk(t *tree.Tree, ch, quit chan int) {
	// Use walk implicit helper
	walkImpl(t, ch, quit)
	// Close the channel when done
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2, quit := make(chan int), make(chan int), make(chan int)
	// Defer close of quit channel after return
	defer close(quit)

	// Use a goroutine to walk both tree 1 & 2
	go Walk(t1, w1, quit)
	go Walk(t2, w2, quit)

	for {
		// Get progressive values & channel status of both channels
		v1, ok1 := <-w1
		v2, ok2 := <-w2

		// If any one channel is closed, return equality channel status
		// If both are closed, that means they were the same because their value
		// did not defer (see other if)
		// If they close at same time and values did not defer, they are same
		// If values did not defer but doesn't close at same time, not same
		// as more values can be in one another tree
		if !ok1 || !ok2 {
			return ok1 == ok2
		}

		// If values are not equal, trees are not equal, return false
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
