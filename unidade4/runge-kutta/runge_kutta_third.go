package rungekutta

import (
	"context"
	"log/slog"

	eulermethod "github.com/ArtroxGabriel/numeric-methods-2/unidade4/euler-method"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

var _ RungeKuttaInterface = (*RungeKuttaThirdOrder)(nil)

type RungeKuttaThirdOrder struct {
	em eulermethod.ExplicitEuler
}

func NewRungeKuttaThird() *RungeKuttaThirdOrder {
	return &RungeKuttaThirdOrder{
		em: *eulermethod.NewExplicitEuler(),
	}
}

func (rk *RungeKuttaThirdOrder) RungeKutta() {}

func (rk *RungeKuttaThirdOrder) getGuess(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *RungeKuttaResult {
	emResult := rk.em.Execute(ctx, fc, initialCondition, initialTime, h)
	return NewRungeKuttaResult(emResult.Time, emResult.State)
}

func (rk *RungeKuttaThirdOrder) Execute(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *RungeKuttaResult {
	slog.InfoContext(ctx, "Executing runge kutta",
		slog.Uint64("order", uint64(3)),
		slog.Any("initialCondition", initialCondition.RawVector().Data),
		slog.Float64("initialTime", initialTime),
		slog.Float64("h", h),
	)
	r, c := 1, initialCondition.Len()

	nextStateHat := mat.NewDense(r+2, c, nil)
	// S_i
	nextStateHat.SetRow(0, initialCondition.RawVector().Data)
	guessResult := NewRungeKuttaResult(0, initialCondition)
	time := initialTime

	for i := 1; i <= 2; i++ {
		// calculate the guess by previous guess
		guessResult = rk.getGuess(ctx, fc, guessResult.State, time, h/2)
		nextStateHat.SetRow(i, guessResult.State.RawVector().Data)
		time += h / 2
	}

	// F(S_i,t_i)
	tmpValCurrent := fc(ctx, nextStateHat, 0, initialTime)

	// F(S_{i+1/2},t_{i+1/2})
	tmpValHalfStep := fc(ctx, nextStateHat, 1, initialTime+h/2)
	tmpValHalfStep.ScaleVec(4, tmpValHalfStep)

	// F(S_{i+1},t_{i+1})
	tmpValNext := fc(ctx, nextStateHat, 2, initialTime+h)

	// F(S_i,t_i) + 4*F(S_{i+1/2},t_{i+1/2})
	tmpValHalfStep.AddVec(tmpValCurrent, tmpValHalfStep)

	// F(S_i,t_i) + 4*F(S_{i+1/2},t_{i+1/2}) + F(S_{i+1},t_{i+1})
	tmpValNext.AddVec(tmpValHalfStep, tmpValNext)

	// Δt/6 *( F(S_i,t_i) + 4*F(S_{i+1/2},t_{i+1/2}) + F(S_{i+1},t_{i+1}) )
	tmpValNext.ScaleVec(h/6, tmpValNext)

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
