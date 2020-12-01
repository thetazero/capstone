package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("hey world!")
	f := MakeEulerEquation(big.NewRat(1, 1), Vector{big.NewRat(3, 1), big.NewRat(1, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(2, 1)}, 10)
	aboutzero := f(Complex{big.NewRat(1539, 10), big.NewRat(0, 1)})
	fmt.Println(aboutzero[0].FloatString(5), aboutzero[1].FloatString(3))
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
