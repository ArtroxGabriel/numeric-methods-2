package rungekutta

import (
	"context"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

type RungeKuttaInterface interface {
	RungeKutta()
	Execute(ctx context.Context, fc types.DerivativeFunc, initialCondition *mat.VecDense, initialTime, h float64) *RungeKuttaResult
}

type RungeKuttaResult struct {
	Time  float64
	State *mat.VecDense
}

func NewRungeKuttaResult(time float64, state *mat.VecDense) *RungeKuttaResult {
	return &RungeKuttaResult{
		Time:  time,
		State: state,
	}
}
