package main

import (
	"fmt"
)

// Define two error types
type ErrNegativeSqrt float64
type ErrZeroSqrt float64

// Define a struct which contains both approx square root value &
// iteration number to find this approx value
type SqrtStruct struct {
	Value float64
	Iterator int
}

// Define precision. This is used to set the precision for our approximation
// Each zero means 1 digit after coma precise
var precision float64 = float64(10000)

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func (e ErrZeroSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt zero !")
}

func Sqrt(x float64) (*SqrtStruct, error) {
  // If x is not a positive integer, return a formated error
  // If x is zero, return a formated error
  switch {
  case x < 0:
    return nil, ErrNegativeSqrt(x)

  case x == 0:
    return nil, ErrZeroSqrt(x)
  }

  // Declare arbitrary start value for our square root
	z := float64(1)

  // Declare tmp outside of if function as we need to use it in for condition
  // We need to initialize it because zero value will trigger for exit condition
	tmp := float64(1)

  // Declare iterator out of for scope to return it at end of function
	i := 1

  // While loop
  // Exit condition is that last iteration is close enough of current iteration
  // Close enough means for digit after coma precise
	for ; int(z * precision) != int((z - tmp) * precision) ; i++ {
    // Use Newton's method as provided in excercise definition
		tmp = (z*z - x) / (2*z)

    // Decrement z with tmp current value. This is part of Newton's method
		z -= tmp
	}

  // Return both approx square root & iterations in struct
	return &SqrtStruct{z, i}, nil
}

func main() {
	fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(-2))
}
