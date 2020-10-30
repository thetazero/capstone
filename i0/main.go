package main

import (
	"fmt"
	"math/big"
	"syscall/js"
)

func main() {
	// lambda, f, g, _, _, p0, c := solve(big.NewRat(1, 1), Vector{big.NewRat(3, 1), big.NewRat(1, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(2, 1)}, 10)
	// fmt.Println(f, g, p0, c)
	// fmt.Println(lambda)

	c := make(chan struct{}, 0)
	js.Global().Set("solve", js.FuncOf(solveJS))
	js.Global().Set("solve_ns", js.FuncOf(solve_nsJS))
	js.Global().Set("get_case", js.FuncOf(getCaseJS))
	<-c
}

//Continued ;
type Continued []float64

func (r RationalFunc) compute() float64 {
	return -1
}

func solve(alpha *big.Rat, p, q Vector, debth int64) (float64, RationalFunc, RationalFunc, Vector, Vector, *big.Rat, string) {
	c := getCase(p, q)
	if c == "i0" || c == "ii" {
		debth++
		pn := computePN(alpha, p, q, debth)
		fmt.Println(pn)
		positiveP := pn[debth+1:]
		negativeP := pn[:debth]
		for i := 0; i < len(negativeP)/2; i++ {
			negativeP[i], negativeP[len(negativeP)-1-i] = negativeP[len(negativeP)-1-i], negativeP[i]
		}
		for i := range positiveP {
			positiveP[i].Inv(positiveP[i])
			negativeP[i].Inv(negativeP[i])
		}
		positiveP = append(Vector{big.NewRat(0, 1)}, positiveP...)
		negativeP = append(Vector{big.NewRat(0, 1)}, negativeP...)
		p0 := pn[debth]
		fp1, f := RationalFromContinued(positiveP)
		gp1, g := RationalFromContinued(negativeP)
		return -1, f, g, fp1.bot, gp1.bot, p0, c
	} else if c == "i+" {
		debth++
		pn := computePN(alpha, p, q, debth)
		negativeP := pn[:debth]
		for i := 0; i < len(negativeP)/2; i++ {
			negativeP[i], negativeP[len(negativeP)-1-i] = negativeP[len(negativeP)-1-i], negativeP[i]
		}
		for i := range negativeP {
			negativeP[i].Inv(negativeP[i])
		}
		negativeP = append(Vector{big.NewRat(0, 1)}, negativeP...)
		p0 := pn[debth]
		gp1, g := RationalFromContinued(negativeP)
		return -1, RationalFunc{}, g, Vector{}, gp1.bot, p0, c
	} else if c == "i-" {
		debth++
		pn := computePN(alpha, p, q, debth)
		positiveP := pn[debth+1:]
		for i := range positiveP {
			positiveP[i].Inv(positiveP[i])
		}
		positiveP = append(Vector{big.NewRat(0, 1)}, positiveP...)
		p0 := pn[debth]
		fp1, f := RationalFromContinued(positiveP)
		return -1, f, RationalFunc{}, fp1.bot, Vector{}, p0, c
	}
	return 0, RationalFunc{}, RationalFunc{}, Vector{}, Vector{}, new(big.Rat), c
}

func getCase(p, q Vector) string {
	r2 := p.SizeSquared()
	qpp := p.Add(q).SizeSquared()
	qmp := p.Add(q.Mul(big.NewRat(-1, 1))).SizeSquared()
	if qpp.Cmp(r2) == 1 && qmp.Cmp(r2) == 1 {
		return "i0"
	} else if qpp.Cmp(r2) == -1 && qmp.Cmp(r2) == 1 {
		return "ii"
	} else if qpp.Cmp(r2) == 0 && qmp.Cmp(r2) == 1 {
		return "i+"
	} else if qpp.Cmp(r2) == 1 && qmp.Cmp(r2) == 0 {
		return "i-"
	}
	return ""
}
