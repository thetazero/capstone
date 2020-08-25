package main

import (
	"fmt"
	"math"
	"strings"
)

//Vector ;
type Vector []float64

func (v Vector) size() float64 {
	return math.Pow(v.SizeSquared(), 0.5)
}

// Mul multiplies the vector by a scalar
func (v Vector) Mul(k float64) Vector {
	r := make(Vector, len(v))
	for i := range v {
		r[i] = v[i] * k
	}
	return r
}

//ShiftRight inserts a 0 at the begining of the vector and shifts the values to the right
func (v Vector) ShiftRight() Vector {
	r := make(Vector, len(v)+1)
	for i := range v {
		r[i+1] = v[i]
	}
	return r
}

// Add returns the summation of the two vectors
func (v Vector) Add(a Vector) Vector {
	r := make(Vector, max(len(v), len(a)))
	for i := 0; i < max(len(v), len(a)); i++ {
		if i < len(a) && i < len(v) {
			r[i] = v[i] + a[i]
		} else if i < len(a) {
			r[i] = a[i]
		} else {
			r[i] = v[i]
		}
	}
	return r
}

//SizeSquared ;
func (v Vector) SizeSquared() float64 {
	x := 0.0
	for _, n := range v {
		x += n * n
	}
	return x
}

func (v Vector) equals(w Vector) bool {
	if len(v) != len(w) {
		return false
	}
	for i := range v {
		if v[i] != w[i] {
			return false
		}
	}
	return true
}

func (v Vector) toString() string {
	valuesText := []string{}

	for i := range v {
		number := v[i]
		text := fmt.Sprintf("%f", number)
		valuesText = append(valuesText, text)
	}
	return "[" + strings.Join(valuesText, ",") + "]"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//PolynomialMul Multiplies the polynomials represented by v1 and v2.
func (v Vector) PolynomialMul(w Vector) Vector {
	res := make(Vector, len(v)+len(w)-1)
	//if this is slow use DFT to do NlogN polynomial multiplication
	for i := range v {
		for j := range w {
			res[i+j] += v[i] * w[j]
		}
	}
	return res
}
