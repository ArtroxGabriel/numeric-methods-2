// Package gausshermite provides implementations of Gauss-Hermite quadrature for numerical integration.
package gausshermite

import "math"

// GaussHermiteCalculator é a interface para os métodos de Gauss-Hermite
type GaussHermiteCalculator interface {
	hermite()
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

// TwoPoints é o método de Gauss-Hermite N = 2
type TwoPoints struct{}

func (gh *TwoPoints) hermite() {}

func NewTwoPoints() *TwoPoints {
	return &TwoPoints{}
}

func (gh *TwoPoints) Calculate(f func(float64) float64) float64 {
	// abscissas dos pontos de Gauss-Hermite N = 2
	x1, x2 := -math.Sqrt2/2.0, math.Sqrt2/2.0

	// pesos correspondentes w1 = w2 = sqrt(pi) / 2
	return (f(x1) + f(x2)) * (math.SqrtPi / 2.0)
}

type ThreePoints struct {
	s [3]float64
	w [3]float64
}

func (gh *ThreePoints) hermite() {}

func NewThreePoints() *ThreePoints {
	// absicssas e pesos do gauss-hermite N=3, formula analitica encontrada nas notas de aula
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

func (gh *FourPoints) hermite() {}

func NewFourPoints() *FourPoints {
	// absicssas e pesos do gauss-hermite N=4, formula analitica encontrada nas notas de aula
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
