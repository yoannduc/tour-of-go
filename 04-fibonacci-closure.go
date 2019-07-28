package main

import (
  "fmt"
)

var displayedNumbers int = 10

// fibonacci is a function that returns a function that returns an int
func fibonacci() func() int {
	lastFib := 0
	currentFib := 1

  // Returns an anonymous function which returns an int
	return func () int {
    // Check current value by adding the 2 last values
    // We need that temporary variable as we need to update old variables
    // with new values
		newFib := lastFib + currentFib

    // Update n and n-1 values
		lastFib = currentFib
		currentFib = newFib

    // Return current value
		return newFib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < displayedNumbers; i++ {
		fmt.Println(f())
	}
}
