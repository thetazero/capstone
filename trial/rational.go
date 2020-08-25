package main

//Rational ;
type Rational struct {
	top, bot Vector
}

//RationalFromContinued returns the rational function equivalent to a fininte contiuned fraction.
func RationalFromContinued(a Vector) Rational {
	r := Rational{}
	// p/q
	pm1 := Vector{1}
	pm2 := Vector{0}
	qm1 := Vector{0}
	qm2 := Vector{1}
	for _, v := range a {
		r.top = pm1.Mul(v).ShiftRight().Add(pm2)
		r.bot = qm1.Mul(v).ShiftRight().Add(qm2)
		pm1, pm2 = r.top, pm1
		qm1, qm2 = r.bot, qm1
		// fmt.Println(r.top, r.bot)
	}
	return r
}

func (r Rational) equals(a Rational) bool {
	return r.top.equals(a.top) && r.bot.equals(a.bot)
}
