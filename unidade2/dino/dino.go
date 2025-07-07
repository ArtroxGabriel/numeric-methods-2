// Package dino implements the Dino method for numerical integration using Gauss-Hermite quadrature.
package dino

import (
	"math"

	gausshermite "github.com/ArtroxGabriel/numeric-methods-2/unidade2/gauss-hermite"
)

type DinoCalculator interface {
	calculate(func(float64) float64, float64, float64) float64
}

var (
	_ DinoCalculator = (*DinoSimples)(nil)
	_ DinoCalculator = (*DinoDuo)(nil)
)

// DinoSimples is exponencial simples
type DinoSimples struct {
	hermite gausshermite.GaussHermiteCalculator
}

func (d *DinoSimples) calculate(f func(float64) float64, a, b float64) float64 {
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

func (d *DinoDuo) calculate(f func(float64) float64, a, b float64) float64 {
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
