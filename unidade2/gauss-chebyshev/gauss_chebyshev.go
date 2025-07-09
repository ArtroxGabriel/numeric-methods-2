// Package gausschebyshev provides implementations of Gauss-Chebyshev quadrature methods.
package gausschebyshev

import "math"

type GaussChebyshevCalculator interface {
	chebyshev()
	Calculate(f func(float64) float64) float64
}

func Integrate(calculator GaussChebyshevCalculator, f func(float64) float64) float64 {
	return calculator.Calculate(f)
}

var (
	_ GaussChebyshevCalculator = (*TwoPoints)(nil)
	_ GaussChebyshevCalculator = (*ThreePoints)(nil)
	_ GaussChebyshevCalculator = (*FourPoints)(nil)
)

type TwoPoints struct{}

func (t *TwoPoints) chebyshev() {}

func NewTwoPoints() *TwoPoints {
	return &TwoPoints{}
}

func (t *TwoPoints) Calculate(f func(float64) float64) float64 {
	x1 := -0.5 * math.Sqrt2
	x2 := 0.5 * math.Sqrt2

	return (math.Pi / 2.0) * (f(x1) + f(x2))
}

type ThreePoints struct{}

func (t *ThreePoints) chebyshev() {}

func NewThreePoints() *ThreePoints {
	return &ThreePoints{}
}

func (t *ThreePoints) Calculate(f func(float64) float64) float64 {
	x1 := -math.Sqrt(3) / 2.0
	x2 := math.Sqrt(3) / 2.0

	return (math.Pi / 3.0) * (f(x1) + f(0) + f(x2))
}

type FourPoints struct{}

func (f *FourPoints) chebyshev() {}

func NewFourPoints() *FourPoints {
	return &FourPoints{}
}

func (*FourPoints) Calculate(f func(float64) float64) float64 {
	x := [4]float64{
		-math.Sqrt(2+math.Sqrt2) / 2.0,
		-math.Sqrt(2-math.Sqrt2) / 2.0,
		math.Sqrt(2-math.Sqrt2) / 2.0,
		math.Sqrt(2+math.Sqrt2) / 2.0,
	}

	return (math.Pi / 4.0) * (f(x[0]) + f(x[1]) + f(x[2]) + f(x[3]))
}
