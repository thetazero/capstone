package main

import (
	"image/color"
	"log"
	"math/big"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

//Path defines a complex path. First argument is the step and Second is the total number of steps for a cycle. Prefer (relativly) equidistant samples.
type Path func(int, int) Complex

//Draw result of going along the path
func (p Path) Draw(f EulerEquation) {
	plt, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(1)
	plt.X.Label.Text = "Real"
	plt.Y.Label.Text = "Imaginary"
	// plt.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
	// 	{Value: 0, Label: "0"}, {Value: 0.25, Label: ""}, {Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: ""}, {Value: 1, Label: "1"},
	// })
	// plt.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
	// 	{Value: 0, Label: "0"}, {Value: 0.25, Label: ""}, {Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: ""}, {Value: 1, Label: "1"},
	// })
	j := 1000
	pts := make(plotter.XYs, j)
	for i := 0; i < j; i++ {
		val := f(p(i, 100))
		// val := p(i, j)
		real, _ := val[0].Float64()
		img, _ := val[1].Float64()
		pts[i] = plotter.XY{X: real, Y: img}
		// pts = append(pts, plotter.XY{X: 1, Y: 1})
	}
	// fmt.Println(pts)
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}
	plt.Add(scatter)

	plotter.DefaultGlyphStyle.Radius = 2
	plotter.DefaultGlyphStyle.Color = color.RGBA{0, 100, 100, 0xff}
	scatter, err = plotter.NewScatter(plotter.XYs{{X: 0, Y: 0}})
	if err != nil {
		log.Panic(err)
	}
	plt.Add(scatter)

	err = plt.Save(500, 500, "test.png")
	if err != nil {
		log.Panic(err)
	}
}

func circle(i, j int, r, x, y *big.Rat) Complex {
	one := big.NewRat(1, 1)
	t := big.NewRat(int64(i), int64(j))
	t.Mul(t, big.NewRat(2, 1))
	t.Sub(t, one)
	t2 := new(big.Rat).Mul(t, t)
	t3 := new(big.Rat).Mul(t2, t)
	t4 := new(big.Rat).Mul(t3, t)
	denom := new(big.Rat).Add(t4, new(big.Rat).Mul(t2, big.NewRat(2, 1)))
	denom.Add(denom, one)

	realnumerator := new(big.Rat).Sub(t4, new(big.Rat).Mul(t2, big.NewRat(6, 1)))
	realnumerator.Add(realnumerator, one)

	imgaginarynumerator := new(big.Rat).Sub(new(big.Rat).Mul(t, big.NewRat(4, 1)), new(big.Rat).Mul(t3, big.NewRat(4, 1)))
	// fmt.Println(t)
	// fmt.Println(realnumerator, denom, imgaginarynumerator)
	return Complex{new(big.Rat).Quo(realnumerator, denom), new(big.Rat).Quo(imgaginarynumerator, denom)}
}
