package gausshermite

import "math"

type GaussHermiteCalculator interface {
	Calculate(f func(float64) float64) float64
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

type ThreePoints struct{}

func NewThreePoints() *ThreePoints {
	return &ThreePoints{}
}

func (gh *ThreePoints) Calculate(f func(float64) float64) float64 {
	x1 := -math.Sqrt(3.0) / 2.0
	x2 := 0.0
	x3 := math.Sqrt(3.0) / 2.0
	return (f(x1)+f(x3))*(math.SqrtPi/6.0) + f(x2)*(2.0*math.SqrtPi)/3.0
}

type FourPoints struct{}

func NewFourPoints() *FourPoints {
	return &FourPoints{}
}

func (gh *FourPoints) Calculate(f func(float64) float64) float64 {
	x1 := -math.Sqrt(3/2 + math.Sqrt(3/2))
	x2 := -math.Sqrt(3/2 - math.Sqrt(3/2))
	x3 := math.Sqrt(3/2 - math.Sqrt(3/2))
	x4 := math.Sqrt(3/2 + math.Sqrt(3/2))

	getW := func(xi float64) float64 {
		numerador := 8.0 * 24.0 * math.SqrtPi
		divisor := 16.0 * math.Pow(8.0*xi*xi*xi-12.0*xi, 2.0)
		return numerador / divisor
	}

	return (f(x1)+f(x4))*getW(x1) + (f(x2)+f(x3))*getW(x2)
}
