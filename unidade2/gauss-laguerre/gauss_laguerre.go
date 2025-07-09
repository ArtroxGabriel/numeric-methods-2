// Package gausslaguerre provides implementations of the Gauss-Laguerre quadrature method
package gausslaguerre

import "math"

// GaussLaguerreCalculator é a interface para os métodos de Gauss-Laguerre
type GaussLaguerreCalculator interface {
	laguerre()
	Calculate(f func(float64) float64) float64
}

func Integrate(calculator GaussLaguerreCalculator, f func(float64) float64) float64 {
	return calculator.Calculate(f)
}

var (
	_ GaussLaguerreCalculator = (*TwoPoints)(nil)
	_ GaussLaguerreCalculator = (*ThreePoints)(nil)
	_ GaussLaguerreCalculator = (*FourPoints)(nil)
)

// TwoPoints é o método de Gauss-Laguerre N = 2
type TwoPoints struct {
	s [2]float64
	w [2]float64
}

func (gl *TwoPoints) laguerre() {}

func NewTwoPoints() *TwoPoints {
	// Abscissas e correspondentes pesos do gauss-laguerre N = 2
	return &TwoPoints{
		s: [2]float64{
			2.0 - math.Sqrt2,
			2.0 + math.Sqrt2,
		},
		w: [2]float64{
			(2.0 + math.Sqrt2) / 4.0,
			(2.0 - math.Sqrt2) / 4.0,
		},
	}
}

func (gl *TwoPoints) Calculate(f func(float64) float64) float64 {
	acc := 0.0
	for i := range gl.s {
		acc += gl.w[i] * f(gl.s[i])
	}

	return acc
}

type ThreePoints struct {
	s [3]float64
	w [3]float64
}

func (gl *ThreePoints) laguerre() {}

func NewThreePoints() *ThreePoints {
	// Abscissas do gauss-laguerre N = 3, calculadas com wolfram alpha
	return &ThreePoints{
		s: [3]float64{
			0.415774556783479,
			2.29428036027904,
			6.28994508293748,
		},
	}
}

func (gl *ThreePoints) Calculate(f func(float64) float64) float64 {
	// função auxiliar para calcular o peso em tempo de execução, resultando em uma melhor precisão
	w := func(xi float64) float64 {
		L5 := (math.Pow(xi, 4) / 24.0) +
			-(2.0 * xi * xi * xi / 3.0) +
			(3.0 * xi * xi) +
			-(4.0 * xi) + 1

		return xi / (16.0 * L5 * L5)
	}

	acc := 0.0
	for i := range gl.s {
		acc += w(gl.s[i]) * f(gl.s[i])
	}

	return acc
}

type FourPoints struct {
	s [4]float64
	w [4]float64
}

func (gl *FourPoints) laguerre() {}

func NewFourPoints() *FourPoints {
	return &FourPoints{
		// calculado com wolfram
		s: [4]float64{
			0.322547689619392,
			1.74576110115835,
			4.53662029692113,
			9.39507091230113,
		},
	}
}

func (gl *FourPoints) Calculate(f func(float64) float64) float64 {
	// função auxiliar para calcular o peso em tempo de execução, resultando em uma melhor precisão
	w := func(xi float64) float64 {
		L5 := -(math.Pow(xi, 5) / 120.0) +
			(5.0 * math.Pow(xi, 4) / 24.0) +
			-(5.0 * xi * xi * xi / 3.0) +
			(5.0 * xi * xi) +
			-(5.0 * xi) + 1

		return xi / (25.0 * L5 * L5)
	}

	var acc float64
	for i := range gl.s {
		acc += w(gl.s[i]) * f(gl.s[i])
	}

	return acc
}
