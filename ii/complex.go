package main

import (
	"math/big"
)

//Complex represents complex numbers through [a, b] = a+bi
type Complex [2]*big.Rat

//Add ;
func (c *Complex) Add(a, b Complex) *Complex {
	(*c)[0] = new(big.Rat).Add(a[0], b[0])
	(*c)[1] = new(big.Rat).Add(a[1], b[1])
	return c
}

//Sub ;
func (c *Complex) Sub(a, b Complex) Complex {
	(*c)[0] = new(big.Rat).Sub(a[0], b[0])
	(*c)[1] = new(big.Rat).Sub(a[1], b[1])
	return *c
}

//Mul ;
func (c *Complex) Mul(a, b Complex) Complex {
	(*c)[0] = new(big.Rat).Mul(a[0], b[0])
	(*c)[0].Sub(c[0], new(big.Rat).Mul(a[1], b[1]))
	(*c)[1] = new(big.Rat).Mul(a[0], b[1])
	(*c)[1].Add(c[1], new(big.Rat).Mul(a[1], b[0]))
	return *c
}

//Quo ;
func (c *Complex) Quo(a, b Complex) Complex {
	(*c)[0] = new(big.Rat).Mul(a[0], b[0])
	(*c)[0].Add((*c)[0], new(big.Rat).Mul(a[1], b[1]))
	(*c)[0].Quo((*c)[0], new(big.Rat).Add(new(big.Rat).Mul(b[0], b[0]), new(big.Rat).Mul(b[1], b[1])))

	(*c)[1] = new(big.Rat).Mul(a[1], b[0])
	(*c)[1].Sub((*c)[1], new(big.Rat).Mul(a[0], b[1]))
	(*c)[1].Quo((*c)[1], new(big.Rat).Add(new(big.Rat).Mul(b[0], b[0]), new(big.Rat).Mul(b[1], b[1])))
	return *c
}

//Equals ;
func (c *Complex) Equals(b Complex) bool {
	for i := range c {
		if c[i].Cmp(b[i]) != 0 {
			return false
		}
	}
	return true
}
