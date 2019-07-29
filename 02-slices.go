package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// Define 2 dimension slice
	var slice [][]uint8

	// For dx param, populate dimension 1 slice with uin8 slice
	for i := 0; i <= dx; i++ {
		// Define the dimension 2 slice
		var sliceInside []uint8

		// For dy param, populate dimension 2 slice with uint8 value
		for j := 0; j <= dy; j++ {
			// Append a uint8 value to the dimension 2 slice
			sliceInside = append(sliceInside, uint8((i+j)/2))
		}

		// Append the builded dimension 2 slice to the dimension 1 slice
		slice = append(slice, sliceInside)
	}

	// Return the builded 2 dimension slice
	return slice
}

func main() {
	pic.Show(Pic)
}
