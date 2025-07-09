// Package newtoncotes implements the Newton-Cotes numerical integration methods.
package newtoncotes

import (
	"math"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade2/result"
)

// NewtonCotesCalculator é a interface para os metodos de newton-cotes
type NewtonCotesCalculator interface {
	Calculate(f func(float64) float64, a, b float64) float64
}

// Integrate realize a integração numérica usando o método de Newton-Cotes
func Integrate(
	method NewtonCotesCalculator,
	f func(float64) float64,
	a, b, e float64,
) *result.IntegrateResult {
	val, iterations := integrateRecursive(method, f, a, b, e)

	return result.NewIntegrateResult(val, iterations)
}

// integrateRecursive é uma função recursiva que divide o intervalo [a, b] em duas partes
func integrateRecursive(
	method NewtonCotesCalculator,
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

// check in compile time
var (
	// metodos fechados
	_ NewtonCotesCalculator = (*ClosedOrder2)(nil)
	_ NewtonCotesCalculator = (*ClosedOrder3)(nil)
	_ NewtonCotesCalculator = (*ClosedOrder4)(nil)

	// metodos abertos
	_ NewtonCotesCalculator = (*OpenOrder2)(nil)
	_ NewtonCotesCalculator = (*OpenOrder3)(nil)
	_ NewtonCotesCalculator = (*OpenOrder4)(nil)
)

// ClosedOrder2 é o metodo de newton-cotes fechado de ordem 2
type ClosedOrder2 struct{}

func NewClosedOrder2() *ClosedOrder2 {
	return &ClosedOrder2{}
}

func (nc *ClosedOrder2) Calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := b - a
	return (f(a) + f(b)) * (h / 2)
}

// ClosedOrder3 é o metodo de newton-cotes fechado de ordem 3
type ClosedOrder3 struct{}

func NewClosedOrder3() *ClosedOrder3 {
	return &ClosedOrder3{}
}

func (nc *ClosedOrder3) Calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 2.0
	return (f(a) + 4.0*f(a+h) + f(b)) * (h / 3.0)
}

// ClosedOrder4 é o metodo de newton-cotes fechado de ordem 4
type ClosedOrder4 struct{}

func NewClosedOrder4() *ClosedOrder4 {
	return &ClosedOrder4{}
}

func (nc *ClosedOrder4) Calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 3.0
	return (f(a) + 3.0*f(a+h) + 3.0*f(a+2.0*h) + f(b)) * (h * 3.0 / 8.0)
}

// OpenOrder2 é o metodo de newton-cotes aberto de ordem 2
type OpenOrder2 struct{}

func NewOpenOrder2() *OpenOrder2 {
	return &OpenOrder2{}
}

func (nc *OpenOrder2) Calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 3.0
	return (f(a+h) + f(a+2.0*h)) * (3.0 * h / 2.0)
}

// OpenOrder3 é o metodo de newton-cotes aberto de ordem 3
type OpenOrder3 struct{}

func NewOpenOrder3() *OpenOrder3 {
	return &OpenOrder3{}
}

func (nc *OpenOrder3) Calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 4.0
	return (2.0*f(a+h) - f(a+2.0*h) + 2.0*f(a+3.0*h)) * (h * 4.0 / 3.0)
}

// OpenOrder4 é o metodo de newton-cotes aberto de ordem 4
type OpenOrder4 struct{}

func NewOpenOrder4() *OpenOrder4 {
	return &OpenOrder4{}
}

func (nc OpenOrder4) Calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 5.0
	return (11.0*f(a+h) + f(a+2.0*h) + f(a+3.0*h) + 11.0*f(a+4*h)) * (5.0 * h / 24.0)
}
