// Package rungekutta implements the Runge-Kutta method for solving ordinary differential equations (ODEs).
package rungekutta

import (
	"context"
	"log/slog"

	eulermethod "github.com/ArtroxGabriel/numeric-methods-2/unidade4/euler-method"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

var _ RungeKuttaInterface = (*RungeKuttaSecondOrder)(nil)

type RungeKuttaSecondOrder struct {
	em eulermethod.ImplicitEuler
}

func NewRungeKuttaSecond() *RungeKuttaSecondOrder {
	return &RungeKuttaSecondOrder{
		em: *eulermethod.NewImplicitEuler(),
	}
}

func (rk *RungeKuttaSecondOrder) RungeKutta() {}

func (rk *RungeKuttaSecondOrder) getGuess(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *RungeKuttaResult {
	emResult := rk.em.Execute(ctx, fc, initialCondition, initialTime, h)
	return NewRungeKuttaResult(emResult.Time, emResult.State)
}

func (rk *RungeKuttaSecondOrder) Execute(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *RungeKuttaResult {
	slog.InfoContext(ctx, "Executing runge kutta",
		slog.Uint64("order", uint64(2)),
		slog.Any("initialCondition", initialCondition.RawVector().Data),
		slog.Float64("initialTime", initialTime),
		slog.Float64("h", h),
	)
	r, c := 1, initialCondition.Len()

	guessResult := rk.getGuess(ctx, fc, initialCondition, initialTime, h)

	nextStateHat := mat.NewDense(r+1, c, nil)
	// S_i
	nextStateHat.SetRow(0, initialCondition.RawVector().Data)
	// S_{i+1}
	nextStateHat.SetRow(1, guessResult.State.RawVector().Data)
	slog.DebugContext(ctx, "state i+1_hat copied")

	// F(S_i,t_i)
	tmpValCurrent := fc(ctx, nextStateHat, 0, initialTime)
	// F(S_{i+1},t_{i+1})
	tmpValNext := fc(ctx, nextStateHat, 1, initialTime+h)

	// F(S_i,t_i) + F(S_{i+1},t_{i+1})
	tmpValNext.AddVec(tmpValCurrent, tmpValNext)

	// Δt/2 *( F(S_i,t_i) + F(S_{i+1},t_{i+1}) )
	tmpValNext.ScaleVec(h/2, tmpValNext)

	// S_i + Δt/2 *( F(S_i,t_i) + F(S_{i+1},t_{i+1}) )
	tmpValNext.AddVec(initialCondition, tmpValNext)

	// x_{i+1}
	nextStateTime := initialTime + h

	slog.InfoContext(ctx, "Computed next state refined",
		slog.Float64("nextTime", nextStateTime),
		slog.Any("nextState", tmpValNext.RawVector().Data),
	)
	return NewRungeKuttaResult(nextStateTime, tmpValNext)
}
