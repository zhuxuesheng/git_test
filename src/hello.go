// Package path implements utility routines for
// manipulating slash-separated filename paths.
package main

import "fmt"

// NewType records the name and value of a object
// and it is important to know
type NewType struct {
	name  string // name of the object
	value int    // its value
}

func fi16(x interface{}, y []int16, s uint) {
	s = 1 << s
	c := float32(s)
	switch xt := x.(type) {
	case []float32:
		for i := range xt {
			y[i] = int16(xt[i] * c)
		}
	case []complex64:
		for i := range xt {
			y[2*i] = int16(real(xt[i]) * c)
			y[2*i+1] = int16(imag(xt[i]) * c)
		}
	}
}

func fi32(x interface{}, y []int32, s uint) {
	s = 1 << s
	c := float32(s)
	switch xt := x.(type) {
	case []float32:
		for i := range xt {
			y[i] = int32(xt[i] * c)
		}
	case []complex64:
		for i := range xt {
			y[2*i] = int32(real(xt[i]) * c)
			y[2*i+1] = int32(imag(xt[i]) * c)
		}
	}
}

// NewType records the name and value of a object
// and it is important to know
func main() {
	x := []complex64{0.5 + 0.4i, 0.3 + 0.2i, 0.1 + 0.0i}
	t := []float32{0.5, 0.4, 0.3, 0.2, 0.1, 0.0}

	y := make([]int16, 2*len(x))
	fi16(x, y, 10)
	fmt.Println(y)

	z := make([]int32, len(t))
	fi32(t, z, 31)
	fmt.Println(z)
}
