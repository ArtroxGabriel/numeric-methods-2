package gausslaguerre

import "math"

type GaussLaguerreCalculator interface {
	Calculate(f func(float64) float64) float64
}

var (
	_ GaussLaguerreCalculator = (*TwoPoints)(nil)
	_ GaussLaguerreCalculator = (*ThreePoints)(nil)
	_ GaussLaguerreCalculator = (*FourPoints)(nil)
)

type TwoPoints struct {
}

func (gl *TwoPoints) Calculate(f func(float64) float64) float64 {
	x1 := 2 - math.Sqrt2
	x2 := 2 + math.Sqrt2
	return (x2*f(x1) + x1*f(x2)) * 0.25
}

type ThreePoints struct{}

// Calculate implements GaussLaguerreCalculator.
func (gl *ThreePoints) Calculate(f func(float64) float64) float64 {
	x1 := 0.4157745568
	x2 := 2.2942803603
	x3 := 6.2899450829

	w1 := 0.7110930099
	w2 := 0.2785177336
	w3 := 0.0103892565

	return w1*f(x1) + w2*f(x2) + w3*f(x3)
}

type FourPoints struct{}

func (gl *FourPoints) Calculate(f func(float64) float64) float64 {
	x := [...]float64{
		0.32255,
		1.7558,
		4.5366,
		9.3951,
	}

	w := func(xi float64) float64 {
		L_5 := -(math.Pow(xi, 5) / 120.0) +
			(5 * math.Pow(xi, 4) / 24.0) +
			-(5 * math.Pow(xi, 3) / 3.0) +
			(5 * xi * xi) +
			-(5 * xi) + 1

		return xi / (25.0 * L_5 * L_5)
	}

	return w(x[0])*f(x[0]) + w(x[1])*f(x[1]) + w(x[2])*f(x[2]) + w(x[3])*f(x[3])
}
