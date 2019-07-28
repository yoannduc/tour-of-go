package main

import (
  "golang.org/x/tour/reader"
)

type MyReader struct{}

func (m *MyReader) Read(b []byte) (int, error) {
  // For each element in b slice, replace it by byte 'A'
  for i := range b {
    // Make use of rune with single quote AKA string litteral
    b[i] = 'A'
  }

  // Return len of b with nil error
  // This is standard reader implementation
  return len(b), nil
}

func main() {
	reader.Validate(&MyReader{})
}
