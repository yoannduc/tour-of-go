package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	var a, z byte

	// If byte is between 'a' & 'z' set limits as 'a' for lower & 'z' for higher
	// If byte is between 'A' & 'Z' set limits as 'A' for lower & 'Z' for higher
	// Else simply return b without shifting it
	switch {
	case b >= 'a' && b <= 'z':
		a, z = 'a', 'z'
	case b >= 'A' && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}

	// Return the mathematical calculus to shift by 13
	// with reset if high or low limit have been reached
	return (b-a+13)%(z-a+1) + a
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	// Get the len to read by r reader, r being an Reader instance
	read, err := r.r.Read(b)

	// Return error if any
	if err != nil {
		return 0, err
	}

	// For len to read from b, apply rot13 function to byte being processed
	for i := 0; i < read; i++ {
		b[i] = rot13(b[i])
	}

	// Return the len read (and transformed)
	return read, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
