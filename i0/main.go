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

func solveJS(this js.Value, args []js.Value) interface{} {
	alpha := new(big.Rat).SetFloat64(args[0].Float())
	fmt.Println(alpha)
	p := Vector{}
	q := Vector{}
	debth := int64(args[3].Int())
	for i := 0; i < args[1].Length(); i++ {
		p = append(p, big.NewRat(int64(args[1].Index(i).Int()), 1))
	}
	for i := 0; i < args[2].Length(); i++ {
		q = append(q, big.NewRat(int64(args[2].Index(i).Int()), 1))
	}
	lambda, f, g, fbotplus1, gbotplus1, p0, pCase := solve(alpha, p, q, debth)
	fmt.Println(pCase)
	res := make(map[string]interface{})
	res["case"] = pCase
	res["lambda"] = lambda
	rho0, _ := p0.Float64()
	res["p0"] = rho0
	switch {
	case pCase == "i0" || pCase == "ii":
		res["ftop"] = f.top.toFloatArr()
		res["fbot"] = f.bot.toFloatArr()
		res["gtop"] = g.top.toFloatArr()
		res["gbot"] = g.bot.toFloatArr()
		res["fbotplus1"] = fbotplus1.toFloatArr()
		res["gbotplus1"] = gbotplus1.toFloatArr()
	case pCase == "i+":
		res["gtop"] = g.top.toFloatArr()
		res["gbot"] = g.bot.toFloatArr()
		res["gbotplus1"] = gbotplus1.toFloatArr()
	case pCase == "i-":
		res["ftop"] = f.top.toFloatArr()
		res["fbot"] = f.bot.toFloatArr()
		res["fbotplus1"] = fbotplus1.toFloatArr()
	}
	return js.ValueOf(res)
}

func solve_nsJS(this js.Value, args []js.Value) interface{} {
	ν := new(big.Rat).SetFloat64(args[0].Float())
	p := Vector{}
	q := Vector{}
	debth := int64(args[3].Int())
	for i := 0; i < args[1].Length(); i++ {
		p = append(p, big.NewRat(int64(args[1].Index(i).Int()), 1))
	}
	for i := 0; i < args[2].Length(); i++ {
		q = append(q, big.NewRat(int64(args[2].Index(i).Int()), 1))
	}
	f, g, fbotplus1, gbotplus1, p0, pCase := solve_ns(ν, p, q, debth)
	res := make(map[string]interface{})
	res["case"] = pCase
	res["ftop"] = f.top.toFloatArr()
	res["fbot"] = f.bot.toFloatArr()
	res["gtop"] = g.top.toFloatArr()
	res["gbot"] = g.bot.toFloatArr()
	res["p0"] = p0.toFloatArr()
	res["fbotplus1"] = fbotplus1.toFloatArr()
	res["gbotplus1"] = gbotplus1.toFloatArr()
	return js.ValueOf(res)
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

func getCaseJS(this js.Value, args []js.Value) interface{} {
	p := Vector{}
	q := Vector{}
	for i := 0; i < args[0].Length(); i++ {
		p = append(p, big.NewRat(int64(args[0].Index(i).Int()), 1))
	}
	for i := 0; i < args[1].Length(); i++ {
		q = append(q, big.NewRat(int64(args[1].Index(i).Int()), 1))
	}
	return getCase(p, q)
}
