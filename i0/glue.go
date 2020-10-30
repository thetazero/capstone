package main

import (
	"fmt"
	"math/big"
	"syscall/js"
)

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

//
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
