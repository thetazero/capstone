package main

import (
	"math/big"
	"syscall/js"
)

func main() {
	// lambda, f, g, _, _, p0 := solve(big.NewRat(1, 1), Vector{big.NewRat(3, 1), big.NewRat(1, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(2, 1)}, 10)
	// fmt.Println(f, g, p0)
	// fmt.Println(lambda)

	c := make(chan struct{}, 0)
	js.Global().Set("solve", js.FuncOf(solveJS))
	<-c
}

func solveJS(this js.Value, args []js.Value) interface{} {
	alpha := big.NewRat(int64(args[0].Int()), 1)
	p := Vector{}
	q := Vector{}
	debth := int64(args[3].Int())
	for i := 0; i < args[1].Length(); i++ {
		p = append(p, big.NewRat(int64(args[1].Index(i).Int()), 1))
	}
	for i := 0; i < args[2].Length(); i++ {
		q = append(q, big.NewRat(int64(args[2].Index(i).Int()), 1))
	}
	lambda, f, g, fbotplus1, gbotplus1, p0 := solve(alpha, p, q, debth)
	res := make(map[string]interface{})
	res["lambda"] = lambda
	res["ftop"] = f.top.toFloatArr()
	res["fbot"] = f.bot.toFloatArr()
	res["gtop"] = g.top.toFloatArr()
	res["gbot"] = g.bot.toFloatArr()
	res["fbotplus1"] = fbotplus1.toFloatArr()
	res["gbotplus1"] = gbotplus1.toFloatArr()
	roh0, _ := p0.Float64()
	res["p0"] = roh0
	return js.ValueOf(res)
}

//Continued ;
type Continued []float64

func (r Rational) compute() float64 {
	return -1
}

func solve(alpha *big.Rat, p, q Vector, debth int64) (float64, Rational, Rational, Vector, Vector, *big.Rat) {
	debth++
	pn := computePN(alpha, p, q, debth)
	positiveP := pn[debth+1:]
	negativeP := pn[:debth]
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
	p0 := pn[debth]
	fp1, f := RationalFromContinued(positiveP)
	// fmt.Println(f)
	gp1, g := RationalFromContinued(negativeP)
	// equation := f.top.PolynomialMul(g.bot).Add(g.top.PolynomialMul(f.bot)).Add(Vector{big.NewRat(0, 1), new(big.Rat).Inv(p0)}.PolynomialMul(g.bot).PolynomialMul(f.bot))
	// var result = func(x float64) float64 {
	// 	res := 0.0
	// 	for i := range equation {
	// 		f, _ := equation[i].Float64()
	// 		res += math.Pow(x, float64(i)) * f

	// 	}
	// 	return res
	// }
	// fmt.Println(p0)
	// fmt.Println(equation.toString())
	return -1, f, g, fp1.top, gp1.top, p0
}
