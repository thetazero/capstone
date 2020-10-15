package main

import (
	"math/big"
	"strings"
)

//Vector ;
type Vector []*big.Rat

// Mul multiplies the vector by a scalar
func (v Vector) Mul(k *big.Rat) Vector {
	r := make(Vector, len(v))
	for i := range r {
		r[i] = new(big.Rat)
	}
	for i := range v {
		r[i].Set(new(big.Rat).Mul(v[i], k))
	}
	return r
}

//ShiftRight inserts a 0 at the begining of the vector and shifts the values to the right
func (v Vector) ShiftRight() Vector {
	r := make(Vector, len(v)+1)
	for i := range r {
		r[i] = new(big.Rat)
	}
	for i := range v {
		r[i+1].Set(v[i])
	}
	return r
}

// Add returns the summation of the two vectors
func (v Vector) Add(a Vector) Vector {
	r := make(Vector, max(len(v), len(a)))
	for i := range r {
		r[i] = new(big.Rat)
	}
	for i := 0; i < max(len(v), len(a)); i++ {
		if i < len(a) && i < len(v) {
			// r[i] = v[i] + a[i]
			r[i].Add(v[i], a[i])
		} else if i < len(a) {
			r[i].Set(a[i])
		} else {
			r[i].Set(v[i])
		}
	}
	return r
}

//Compute computes the value of the vector at x as if it is a polynomial, ie: [1,2,3]=1+2x+3x^2
func (v Vector) Compute(x *big.Rat) *big.Rat {
	c := big.NewRat(1, 1)
	ans := big.NewRat(0, 1)
	for i := range v {
		ans.Add(new(big.Rat).Mul(c, v[i]), ans)
		c.Mul(c, x)
	}
	return ans
}

//SizeSquared ;
func (v Vector) SizeSquared() *big.Rat {
	x := big.NewRat(0, 1)
	for _, n := range v {
		x.Add(x, new(big.Rat).Mul(n, n))
	}
	return x
}

func (v Vector) equals(w Vector) bool {
	if len(v) != len(w) {
		return false
	}
	for i := range v {
		if v[i].Cmp(w[i]) != 0 {
			return false
		}
	}
	return true
}

func (v Vector) toString() string {
	valuesText := []string{}

	for i := range v {
		number := v[i]
		text := number.FloatString(5)
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
	for i := range res {
		res[i] = big.NewRat(0, 1)
	}
	//if this is slow use DFT to do NlogN polynomial multiplication
	for i := range v {
		for j := range w {
			// res[i+j] += v[i] * w[j]
			res[i+j].Add(res[i+j], new(big.Rat).Mul(v[i], w[j]))
		}
	}
	return res
}

func (v Vector) toFloatArr() []interface{} {
	x := make([]interface{}, len(v))
	for i := range x {
		f, _ := v[i].Float64()
		x[i] = f
	}
	return x
}

func (v *Vector) reverse() {
	for i := 0; i < len(*v)/2; i++ {
		(*v)[i], (*v)[len(*v)-1-i] = (*v)[len(*v)-1-i], (*v)[i]
	}
}
