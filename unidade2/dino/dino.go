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

// exponencial simples
type DinoSimples struct {
	hermite gausshermite.GaussHermiteCalculator
}

func (d *DinoSimples) calculate(f func(float64) float64, a, b float64) float64 {
	x_s := func(s float64) float64 {
		return ((a + b) + (b-a)*math.Tanh(s)) / 2.0
	}
	dx_s := func(s float64) float64 {
		return (b - a) / (2.0 * math.Pow(math.Cosh(s), 2))
	}

	f_hat := func(s float64) float64 {
		return f(x_s(s)) * dx_s(s)
	}

	f_hat_2 := func(s float64) float64 {
		return math.Exp(s*s) * f_hat(s)
	}

	return d.hermite.Calculate(f_hat_2)
}

// exponencial dupla
type DinoDuo struct {
	hermite gausshermite.GaussHermiteCalculator
}

func (d *DinoDuo) calculate(f func(float64) float64, a, b float64) float64 {
	piOverTwo := math.Pi / 2.0
	x_s := func(s float64) float64 {
		return ((a + b) + (b-a)*math.Tanh(math.Sinh(s)*piOverTwo)) / 2.0
	}
	dx_s := func(s float64) float64 {
		return piOverTwo * ((b - a) * math.Cosh(s)) / (2.0 * math.Pow(math.Cosh(piOverTwo*math.Sinh(s)), 2))
	}

	f_hat := func(s float64) float64 {
		return f(x_s(s)) * dx_s(s)
	}

	f_hat_2 := func(s float64) float64 {
		return math.Exp(s*s) * f_hat(s)
	}

	return d.hermite.Calculate(f_hat_2)
}
