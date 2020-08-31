package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/DzananGanic/numericalgo/root"
)

func main() {
	lambda := solve(big.NewRat(1, 1), Vector{big.NewRat(3, 1), big.NewRat(1, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(2, 1)})
	// fmt.Println(RationalFromContinued(Vector{1, 1, 1}))
	fmt.Println(lambda)
}

//Continued ;
type Continued []float64

func (r Rational) compute() float64 {
	return -1
}

func solve(alpha *big.Rat, p, q Vector) float64 {
	cap := int64(10)
	pn := computePN(alpha, p, q, cap)
	positiveP := pn[cap+1:]
	negativeP := pn[:cap]
	for i := 0; i < len(negativeP)/2; i++ {
		negativeP[i], negativeP[len(negativeP)-1-i] = negativeP[len(negativeP)-1-i], negativeP[i]
	}
	for i := range positiveP {
		// positiveP[i] = 1 / positiveP[i]
		positiveP[i].Inv(positiveP[i])
		// negativeP[i] = 1 / negativeP[i]
		negativeP[i].Inv(negativeP[i])
	}
	positiveP = append(Vector{big.NewRat(0, 1)}, positiveP...)
	negativeP = append(Vector{big.NewRat(0, 1)}, negativeP...)
	p0 := pn[cap]
	f := RationalFromContinued(positiveP)
	// fmt.Println(f)
	g := RationalFromContinued(negativeP)
	equation := f.top.PolynomialMul(g.bot).Add(g.top.PolynomialMul(f.bot)).Add(Vector{big.NewRat(0, 1), new(big.Rat).Inv(p0)}.PolynomialMul(g.bot).PolynomialMul(f.bot))
	var result = func(x float64) float64 {
		res := 0.0
		for i := range equation {
			f, _ := equation[i].Float64()
			res += math.Pow(x, float64(i)) * f

		}
		return res
	}
	// fmt.Println(p0)
	// fmt.Println(equation.toString())
	return root.Bisection(result, 0.0001, 0, 1000)
}
