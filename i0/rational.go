package main

import (
	"math/big"
)

//Rational ;
type RationalFunc struct {
	top, bot Vector
}

//RationalFromContinued returns the rational function equivalent to a fininte contiuned fraction.
func RationalFromContinued(a Vector) (RationalFunc, RationalFunc) {
	r := RationalFunc{}
	// p/q
	pm1 := Vector{big.NewRat(1, 1)}
	pm2 := Vector{big.NewRat(0, 1)}
	qm1 := Vector{big.NewRat(0, 1)}
	qm2 := Vector{big.NewRat(1, 1)}
	for _, v := range a {
		r.top = pm1.Mul(v).ShiftRight().Add(pm2)
		r.bot = qm1.Mul(v).ShiftRight().Add(qm2)
		pm1, pm2 = r.top, pm1
		qm1, qm2 = r.bot, qm1
		// fmt.Println(r.top.toString(), r.bot.toString())
	}
	return r, RationalFunc{top: pm2, bot: qm2}
}

func (r RationalFunc) equals(a RationalFunc) bool {
	return r.top.equals(a.top) && r.bot.equals(a.bot)
}

func RationalFromContinuedVector(a []Vector) (RationalFunc, RationalFunc) {
	r := RationalFunc{}
	// p/q
	pm1 := Vector{big.NewRat(1, 1)}
	pm2 := Vector{big.NewRat(0, 1)}
	qm1 := Vector{big.NewRat(0, 1)}
	qm2 := Vector{big.NewRat(1, 1)}
	for i := range a[0] {
		r.top = pm1.PolynomialMul(Vector{a[0][i], a[1][i]}).ShiftRight().Add(pm2)
		r.bot = qm1.PolynomialMul(Vector{a[0][i], a[1][i]}).ShiftRight().Add(qm2)
		pm1, pm2 = r.top, pm1
		qm1, qm2 = r.bot, qm1
		// fmt.Println(r.top.toString(), r.bot.toString())
	}
	return r, RationalFunc{top: pm2, bot: qm2}
}
