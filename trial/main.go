package main

import (
	"fmt"
	"math"

	"github.com/DzananGanic/numericalgo/root"
)

func main() {
	lambda := solve(1, Vector{3, 1}, Vector{-1, 2})
	// fmt.Println(RationalFromContinued(Vector{1, 1, 1}))
	fmt.Println(lambda)
}

//Continued ;
type Continued []float64

func (r Rational) compute() float64 {
	return -1
}

func solve(alpha float64, p, q Vector) float64 {
	cap := 11
	pn := computePN(alpha, p, q, cap)
	positiveP := pn[cap+1:]
	negativeP := pn[:cap]
	for i := 0; i < len(negativeP)/2; i++ {
		negativeP[i], negativeP[len(negativeP)-1-i] = negativeP[len(negativeP)-1-i], negativeP[i]
	}
	for i := range positiveP {
		positiveP[i] = 1 / positiveP[i]
		negativeP[i] = 1 / negativeP[i]
	}
	positiveP = append([]float64{0}, positiveP...)
	negativeP = append([]float64{0}, negativeP...)
	p0 := pn[cap]
	f := RationalFromContinued(positiveP)
	fmt.Println(f)
	g := RationalFromContinued(negativeP)
	equation := f.top.PolynomialMul(g.bot).Add(g.top.PolynomialMul(f.bot)).Add(Vector{0, 1.0 / p0}.PolynomialMul(g.bot).PolynomialMul(f.bot))
	var result = func(x float64) float64 {
		res := 0.0
		for i := range equation {
			res += math.Pow(x, float64(i)) * equation[i]
		}
		return res
	}
	fmt.Println(p0)
	return root.Bisection(result, 0.0001, 0, 1000)
}
