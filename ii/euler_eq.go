package main

import (
	"log"
	"math/big"
)

// EulerEquation is the complex function described in the paper.
type EulerEquation func(x Complex) Complex

//MakeEulerEquation ;
func MakeEulerEquation(alpha *big.Rat, p Vector, q Vector, depth int64) EulerEquation {
	c := getCase(p, q)
	if c == "i0" || c == "ii" {
		pn := computePN(alpha, p, q, depth)
		positiveP := pn[depth+1:]
		negativeP := pn[:depth]
		for i := 0; i < len(negativeP)/2; i++ {
			negativeP[i], negativeP[len(negativeP)-1-i] = negativeP[len(negativeP)-1-i], negativeP[i]
		}
		for i := range positiveP {
			positiveP[i].Inv(positiveP[i])
			negativeP[i].Inv(negativeP[i])
		}
		positiveP = append(Vector{big.NewRat(0, 1)}, positiveP...)
		negativeP = append(Vector{big.NewRat(0, 1)}, negativeP...)
		p0 := Complex{pn[depth], big.NewRat(0, 1)}
		_, f := RationalFromContinued(positiveP)
		_, g := RationalFromContinued(negativeP)
		return func(x Complex) Complex {
			ans := *new(Complex).Add(f.ComplexCompute(x), g.ComplexCompute(x))
			ans.Add(ans, new(Complex).Quo(x, p0))
			return ans
		}
	}
	log.Fatal("i+, i- cases not supported")
	return func(x Complex) Complex {
		return x
	}
}
