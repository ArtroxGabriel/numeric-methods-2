package gausschebyshev

import "math"

type GaussChebyshevCalculator interface {
	Calculate(f func(float64) float64) float64
}

var (
	_ GaussChebyshevCalculator = (*TwoPoints)(nil)
	_ GaussChebyshevCalculator = (*ThreePoints)(nil)
	_ GaussChebyshevCalculator = (*FourPoints)(nil)
)

type TwoPoints struct{}

func (t *TwoPoints) Calculate(f func(float64) float64) float64 {
	x1 := -1.0 / math.Sqrt2
	x2 := 1.0 / math.Sqrt2

	return (math.Pi / 2.0) * (f(x1) + f(x2))
}

type ThreePoints struct{}

func (t *ThreePoints) Calculate(f func(float64) float64) float64 {
	x1 := -math.Sqrt(3) / 2.0
	x2 := math.Sqrt(3) / 2.0

	return (math.Pi / 3.0) * (f(x1) + f(0) + f(x2))
}

type FourPoints struct{}

func (*FourPoints) Calculate(f func(float64) float64) float64 {
	x := [...]float64{
		-math.Sqrt(2+math.Sqrt2) / 2.0,
		-math.Sqrt(2-math.Sqrt2) / 2.0,
		math.Sqrt(2-math.Sqrt2) / 2.0,
		math.Sqrt(2+math.Sqrt2) / 2.0,
	}

	return (math.Pi / 4.0) * (f(x[0]) + f(x[1]) + f(x[2]) + f(x[3]))
}
