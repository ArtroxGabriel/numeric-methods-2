// Package gausshermite provides implementations of Gauss-Hermite quadrature for numerical integration.
package gausshermite

import "math"

type GaussHermiteCalculator interface {
	Calculate(f func(float64) float64) float64
}

func Integrate(calculator GaussHermiteCalculator, f func(float64) float64) float64 {
	return calculator.Calculate(f)
}

var (
	_ GaussHermiteCalculator = (*TwoPoints)(nil)
	_ GaussHermiteCalculator = (*ThreePoints)(nil)
	_ GaussHermiteCalculator = (*FourPoints)(nil)
)

type TwoPoints struct{}

func NewTwoPoints() *TwoPoints {
	return &TwoPoints{}
}

func (gh *TwoPoints) Calculate(f func(float64) float64) float64 {
	x1, x2 := -math.Sqrt2/2.0, math.Sqrt2/2.0
	return (f(x1) + f(x2)) * (math.SqrtPi / 2.0)
}

type ThreePoints struct {
	s [3]float64
	w [3]float64
}

func NewThreePoints() *ThreePoints {
	return &ThreePoints{
		s: [3]float64{
			-math.Sqrt(6.0) / 2.0,
			0.0,
			math.Sqrt(6.0) / 2.0,
		},
		w: [3]float64{
			math.SqrtPi / 6.0,
			2 * math.SqrtPi / 3.0,
			math.SqrtPi / 6.0,
		},
	}
}

func (gh *ThreePoints) Calculate(f func(float64) float64) float64 {
	var acc float64

	for i := range gh.w {
		acc += gh.w[i] * f(gh.s[i])
	}

	return acc
}

type FourPoints struct {
	s [4]float64
	w [4]float64
}

func NewFourPoints() *FourPoints {
	return &FourPoints{
		s: [4]float64{
			-math.Sqrt((3.0 + math.Sqrt(6)) / 2.0),
			-math.Sqrt((3.0 - math.Sqrt(6)) / 2.0),
			math.Sqrt((3.0 - math.Sqrt(6)) / 2.0),
			math.Sqrt((3.0 + math.Sqrt(6)) / 2.0),
		},
		w: [4]float64{
			math.SqrtPi / (12 + 4.0*math.Sqrt(6)),
			math.SqrtPi / (12 - 4.0*math.Sqrt(6)),
			math.SqrtPi / (12 - 4.0*math.Sqrt(6)),
			math.SqrtPi / (12 + 4.0*math.Sqrt(6)),
		},
	}
}

func (gh *FourPoints) Calculate(f func(float64) float64) float64 {
	var acc float64

	for i := range gh.w {
		acc += gh.w[i] * f(gh.s[i])
	}

	return acc
}
