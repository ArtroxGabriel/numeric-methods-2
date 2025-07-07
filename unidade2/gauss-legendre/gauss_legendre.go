// Package gausslegendre implements the Gauss-Legendre quadrature method for numerical integration.
package gausslegendre

import (
	"math"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade2/result"
)

func getXFunc(a, b float64) func(float64) float64 {
	return func(s float64) float64 {
		return (a + b + (a-b)*s) / 2.0
	}
}

type GaussLegendreCalculator interface {
	Calculate(f func(float64) float64, a, b float64) float64
}

func Integrate(
	method GaussLegendreCalculator,
	f func(float64) float64,
	a, b, e float64,
) *result.IntegrateResult {
	val, iterations := integrateRecursive(method, f, a, b, e)

	return result.NewIntegrateResult(val, iterations)
}

func integrateRecursive(
	method GaussLegendreCalculator,
	f func(float64) float64,
	a, b,
	tolerance float64,
) (float64, int) {
	integralWhole := method.Calculate(f, a, b)
	mid := (a + b) / 2.0
	integralPart1 := method.Calculate(f, a, mid)
	integralPart2 := method.Calculate(f, mid, b)
	sumOfParts := integralPart1 + integralPart2

	err := math.Abs(integralWhole - sumOfParts)

	if err < tolerance {
		return sumOfParts, 1
	}

	newTolerance := tolerance / 2.0
	leftResult, leftIters := integrateRecursive(method, f, a, mid, newTolerance)
	rightResult, rightIters := integrateRecursive(method, f, mid, b, newTolerance)

	totalResult := leftResult + rightResult
	totalIterations := 1 + leftIters + rightIters

	return totalResult, totalIterations
}

var (
	_ GaussLegendreCalculator = (*TwoPoints)(nil)
	_ GaussLegendreCalculator = (*ThreePoints)(nil)
	_ GaussLegendreCalculator = (*FourPoints)(nil)
)

type TwoPoints struct {
	s [2]float64
}

func NewTwoPoints() *TwoPoints {
	return &TwoPoints{
		s: [2]float64{
			-math.Sqrt(1.0 / 3.0),
			math.Sqrt(1.0 / 3.0),
		},
	}
}

func (gl *TwoPoints) Calculate(f func(float64) float64, a, b float64) float64 {
	h := (b - a) / 2.0
	x := getXFunc(a, b)
	acc := f(x(gl.s[0])) + f(x(gl.s[1]))
	return h * acc
}

type ThreePoints struct {
	s [3]float64
	w [3]float64
}

func NewThreePoints() *ThreePoints {
	return &ThreePoints{
		s: [3]float64{
			-math.Sqrt(3.0 / 5.0),
			0.0,
			math.Sqrt(3.0 / 5.0),
		},
		w: [3]float64{
			5.0 / 9.0,
			8.0 / 9.0,
			5.0 / 9.0,
		},
	}
}

func (gl *ThreePoints) Calculate(f func(float64) float64, a, b float64) float64 {
	h := (b - a) / 2.0
	x := getXFunc(a, b)
	acc := 0.0
	for i, wi := range gl.w {
		acc += wi * f(x(gl.s[i]))
	}
	return h * acc
}

type FourPoints struct {
	s [4]float64
	w [4]float64
}

func NewFourPoints() *FourPoints {
	return &FourPoints{
		s: [4]float64{
			-math.Sqrt(3.0/7.0 + 2.0/7.0*math.Sqrt(6.0/5.0)),
			-math.Sqrt(3.0/7.0 - 2.0/7.0*math.Sqrt(6.0/5.0)),
			math.Sqrt(3.0/7.0 - 2.0/7.0*math.Sqrt(6.0/5.0)),
			math.Sqrt(3.0/7.0 + 2.0/7.0*math.Sqrt(6.0/5.0)),
		},
		w: [4]float64{
			(18 - math.Sqrt(30.0)) / 36,
			(18 + math.Sqrt(30.0)) / 36,
			(18 + math.Sqrt(30.0)) / 36,
			(18 - math.Sqrt(30.0)) / 36,
		},
	}
}

func (gl *FourPoints) Calculate(f func(float64) float64, a, b float64) float64 {
	h := (b - a) / 2.0
	x := getXFunc(a, b)
	acc := 0.0
	for i, wi := range gl.w {
		acc += wi * f(x(gl.s[i]))
	}
	return h * acc
}
