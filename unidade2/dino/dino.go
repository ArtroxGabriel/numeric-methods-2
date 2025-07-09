// Package dino implements the Dino method for numerical integration using Gauss-Hermite quadrature.
package dino

import (
	"math"

	gausshermite "github.com/ArtroxGabriel/numeric-methods-2/unidade2/gauss-hermite"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade2/result"
)

type interval struct {
	a, b float64
}

// DinoCalculator é a interface para os métodos de integração de exponenciais
type DinoCalculator interface {
	Calculate(func(float64) float64, float64, float64) float64
}

// IntegrateDino é a função que realiza a integração numérica usando um dos métodos Dino
func IntegrateDino(calculator DinoCalculator, f func(float64) float64, a, b float64) *result.IntegrateResult {
	iterations := 0
	tolerance := 1e-5

	var integrateRecursive func(float64, float64) (float64, int)

	integrateRecursive = func(currentA, currentB float64) (float64, int) {
		iterations++

		areaFull := calculator.Calculate(f, currentA, currentB)

		mid := (currentA + currentB) / 2.0

		areaLeft := calculator.Calculate(f, currentA, mid)
		areaRight := calculator.Calculate(f, mid, currentB)

		areaHalves := areaLeft + areaRight

		errorEstimate := math.Abs(areaFull - areaHalves)

		if errorEstimate < tolerance || (currentB-currentA) < 1e-9 {
			return areaHalves, 1
		}

		// Se não convergiu, chama recursivamente para as duas metades
		areaLeftRecursive, iterLeft := integrateRecursive(currentA, mid)
		areaRightRecursive, iterRight := integrateRecursive(mid, currentB)

		return areaLeftRecursive + areaRightRecursive, iterLeft + iterRight
	}

	finalArea, _ := integrateRecursive(a, b)

	return result.NewIntegrateResult(finalArea, iterations)
}

var (
	_ DinoCalculator = (*DinoSimples)(nil)
	_ DinoCalculator = (*DinoDuo)(nil)
)

// DinoSimples is exponencial simples
type DinoSimples struct {
	hermite gausshermite.GaussHermiteCalculator
}

func NewDinoSimples() *DinoSimples {
	return &DinoSimples{
		hermite: gausshermite.NewFourPoints(),
	}
}

func (d *DinoSimples) Calculate(f func(float64) float64, a, b float64) float64 {
	// mudanças de variaveis para dino simples
	xS := func(s float64) float64 {
		return ((a + b) + (b-a)*math.Tanh(s)) / 2.0
	}
	dxS := func(s float64) float64 {
		return (b - a) / (2.0 * math.Pow(math.Cosh(s), 2))
	}

	fHat := func(s float64) float64 {
		return f(xS(s)) * dxS(s)
	}

	fHat2 := func(s float64) float64 {
		return math.Exp(s*s) * fHat(s)
	}

	return d.hermite.Calculate(fHat2)
}

// DinoDuo is exponencial dupla
type DinoDuo struct {
	hermite gausshermite.GaussHermiteCalculator
}

func NewDinoDuo() *DinoDuo {
	return &DinoDuo{
		hermite: gausshermite.NewFourPoints(),
	}
}

func (d *DinoDuo) Calculate(f func(float64) float64, a, b float64) float64 {
	// mudanças de variaveis para dino duplo
	piOverTwo := math.Pi / 2.0
	xS := func(s float64) float64 {
		return ((a + b) + (b-a)*math.Tanh(math.Sinh(s)*piOverTwo)) / 2.0
	}
	dxS := func(s float64) float64 {
		return piOverTwo * ((b - a) * math.Cosh(s)) / (2.0 * math.Pow(math.Cosh(piOverTwo*math.Sinh(s)), 2))
	}

	fHat := func(s float64) float64 {
		return f(xS(s)) * dxS(s)
	}

	fhat2 := func(s float64) float64 {
		return math.Exp(s*s) * fHat(s)
	}

	return d.hermite.Calculate(fhat2)
}
