package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	f := MakeEulerEquation(big.NewRat(1, 1), Vector{big.NewRat(2, 1), big.NewRat(0, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(-1, 1)}, 51)

	// f := MakeEulerEquation(big.NewRat(1, 1), Vector{big.NewRat(3, 1), big.NewRat(1, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(2, 1)}, 10)

	// f := MakeEulerEquation(big.NewRat(0, 1), Vector{big.NewRat(2, 1), big.NewRat(0, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(-1, 1)}, 120)

	point := Complex{big.NewRat(1, 10), big.NewRat(0, 1)}
	// point := Complex{big.NewRat(1538, 1000), big.NewRat(0, 1)}
	fmt.Println("point:", point[0].FloatString(3), point[1].FloatString(3))
	aboutzero := f(point)
	fmt.Println("val:", aboutzero[0].FloatString(3), aboutzero[1].FloatString(3))

	p := Path(func(i, j int) Complex {
		a, b, theta := 0.5, 10.0, math.Pi/20
		acircle, bcircle, line := a*(math.Pi-2*theta), b*(math.Pi-2*theta), (b - a)
		len := acircle + bcircle + 2*line
		position := float64(i) / float64(j) * len
		if position < acircle { //inner circle
			position = acircle - position
			offset := -math.Pi/2 + theta
			real, imag := a*math.Cos(position/a+offset), a*math.Sin(position/a+offset)
			return Complex{new(big.Rat).SetFloat64(real), new(big.Rat).SetFloat64(imag)}
		}
		position -= acircle
		if position < line { //lower line
			t := position / line
			angle := math.Pi/2 - theta
			real := t*b*math.Cos(angle) + (1-t)*a*math.Cos(angle)
			imag := t*-b*math.Sin(angle) + (1-t)*-a*math.Sin(angle)
			return Complex{new(big.Rat).SetFloat64(real), new(big.Rat).SetFloat64(imag)}
		}
		position -= line
		if position < bcircle { //outer circle
			offset := -math.Pi/2 + theta
			real, imag := b*math.Cos(position/b+offset), b*math.Sin(position/b+offset)
			return Complex{new(big.Rat).SetFloat64(real), new(big.Rat).SetFloat64(imag)}
		}
		position -= bcircle
		//upper line
		t := 1 - position/line
		angle := math.Pi/2 - theta
		real := t*b*math.Cos(angle) + (1-t)*a*math.Cos(angle)
		imag := t*b*math.Sin(angle) + (1-t)*a*math.Sin(angle)
		return Complex{new(big.Rat).SetFloat64(real), new(big.Rat).SetFloat64(imag)}
	})

	samples := 2000
	p.Draw(f, samples, "_result.png", true)
	p.Draw(func(x Complex) Complex {
		return x
	}, samples, "_contour.png", false)

	/*
		w := big.NewRat(15, 100)
		x := big.NewRat(70, 100)
		y := big.NewRat(160, 100)
		nx := big.NewRat(0, 1)
		ny := big.NewRat(0, 1)
		ax := new(big.Rat).Add(x, new(big.Rat).Mul(w, nx))
		ay := new(big.Rat).Add(y, new(big.Rat).Mul(w, ny))
		fmt.Printf("<%s, %s> ⇒ <%s, %s> \n", ax.FloatString(3), ay.FloatString(3), new(big.Rat).Add(ax, w).FloatString(3), new(big.Rat).Sub(ay, w).FloatString(3))

		path := Path(func(i, j int) Complex {
			c := square(i, j, w, ax, ay)
			// c := square(i, j, w, ax, new(big.Rat).Mul(ay, big.NewRat(-1, 1)))
			return c
		})
		α := big.NewRat(1, 1)
		p := Vector{big.NewRat(2, 1), big.NewRat(0, 1)}
		q := Vector{big.NewRat(-1, 1), big.NewRat(-1, 1)}

		DrawDebths(α, p, q, path, 80, 80)
	*/

}

func DrawDebths(α *big.Rat, p, q Vector, path Path, start, end int) {
	title := α.FloatString(3) + "_" + p.toString() + "_" + q.toString()
	samples := 800
	for i := start; i <= end; i++ {
		f := MakeEulerEquation(α, p, q, int64(i))
		path.Draw(f, samples, "output/output_"+title+fmt.Sprintf("%02d", i)+".png", true)
	}
	path.Draw(func(x Complex) Complex {
		return x
	}, samples, "output/input_"+title+".png", false)
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
