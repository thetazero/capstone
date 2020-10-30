package main

import (
	"math/big"
)

//Rational ;
type RationalFunc struct {
	top, bot Vector
}

//RationalFromContinued returns the rational function equivalent to a fininte contiuned fraction.
// func RationalFromContinued(a Vector) (RationalFunc, RationalFunc) {
// 	r := RationalFunc{}
// 	// p/q
// 	pm1 := Vector{big.NewRat(1, 1)}
// 	pm2 := Vector{big.NewRat(0, 1)}
// 	qm1 := Vector{big.NewRat(0, 1)}
// 	qm2 := Vector{big.NewRat(1, 1)}
// 	for _, v := range a {
// 		r.top = pm1.Mul(v).ShiftRight().Add(pm2)
// 		r.bot = qm1.Mul(v).ShiftRight().Add(qm2)
// 		pm1, pm2 = r.top, pm1
// 		qm1, qm2 = r.bot, qm1
// 	}
// 	return r, RationalFunc{top: pm2, bot: qm2}
// }

func RationalFromContinued(a Vector) (RationalFunc, RationalFunc) {
	// p/q
	// [1,2,3]= 1x^2+2x+3
	// pm1 := Vector{big.NewRat(1, 1)}
	// pm2 := Vector{big.NewRat(0, 1)}
	// qm1 := Vector{big.NewRat(0, 1)}
	// qm2 := Vector{big.NewRat(1, 1)}
	pm1 := make(Vector, len(a)+1)
	pm2 := make(Vector, len(a)+1)
	qm1 := make(Vector, len(a))
	qm2 := make(Vector, len(a))
	for i := range pm1 {
		pm1[i] = big.NewRat(0, 1)
		pm2[i] = big.NewRat(0, 1)
	}
	for i := range qm1 {
		qm1[i] = big.NewRat(0, 1)
		qm2[i] = big.NewRat(0, 1)
	}
	pm1[0].Set(big.NewRat(1, 1))
	qm2[0].Set(big.NewRat(1, 1))
	for j, v := range a {
		for i := len(a); i >= 0; i-- {
			if i == 0 || i == 1 {
				pm2[i].Mul(pm1[i], v)
			} else {
				pm2[i].Mul(pm1[i], v)
				pm2[i].Add(pm2[i], pm2[i-2])
			}

		}
		if j == 0 {
			qm2[0].Set(qm2[0])
		} else {
			for i := len(a) - 1; i >= 0; i-- {
				qm2[i].Mul(qm1[i], v)
				if !(i == 0 || i == 1) {
					qm2[i].Add(qm2[i], qm2[i-2])
				}

			}
		}
		pm1, pm2 = pm2, pm1
		qm1, qm2 = qm2, qm1
		// pm1, pm2 = r.top, pm1
	}
	pm1.reverse()
	pm2.reverse()
	qm1.reverse()
	qm2.reverse()
	return RationalFunc{top: pm1, bot: qm1}, RationalFunc{top: pm2, bot: qm2}
}

func (r RationalFunc) equals(a RationalFunc) bool {
	return r.top.equals(a.top) && r.bot.equals(a.bot)
}

func (r RationalFunc) Compute(x *big.Rat) *big.Rat {
	return new(big.Rat).Quo(r.top.Compute(x), r.bot.Compute(x))
}

func RationalFromContinuedVector(a []Vector) (RationalFunc, RationalFunc) {
	r := RationalFunc{}
	// p/q
	pm1 := Vector{big.NewRat(1, 1)}
	pm2 := Vector{big.NewRat(0, 1)}
	qm1 := Vector{big.NewRat(0, 1)}
	qm2 := Vector{big.NewRat(1, 1)}
	for i := range a[0] {
		x := make(Vector, len(a))
		for j := range a {
			x[j] = new(big.Rat).Set(a[j][i])
		}
		r.top = pm1.PolynomialMul(x).Add(pm2)
		r.bot = qm1.PolynomialMul(x).Add(qm2)
		pm1, pm2 = r.top, pm1
		qm1, qm2 = r.bot, qm1
	}
	return r, RationalFunc{top: pm2, bot: qm2}
}
