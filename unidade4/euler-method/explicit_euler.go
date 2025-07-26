// Package eulermethod implements the explicit and implicit Euler method for solving ordinary differential equations (ODEs).
package eulermethod

import (
	"context"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

// y[x+1] = y[x] + h*f(x, y[x])


type ExplicitEuler struct{}

func NewExplicitEuler() *ExplicitEuler {
	return &ExplicitEuler{}
}

func (em *ExplicitEuler) Execute(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *EulerResult {
	slog.InfoContext(ctx, "Executing Explicit Euler Method",
		slog.Any("initialCondition", initialCondition.RawVector().Data),
		slog.Float64("initialTime", initialTime),
		slog.Float64("h", h),
	)
	r, c := 1, initialCondition.Len()

	previousState := mat.NewDense(r, c, nil)
	previousState.SetRow(0, initialCondition.RawVector().Data)
	slog.InfoContext(ctx, "Initial state copied")

	// copy previous state to next, except new dimension for next state

	nextState := mat.NewDense(r+1, c, nil)
	// Δt*F(S_i,t_i)
	tempState := fc(ctx, previousState, 0, initialTime)
	tempState.AddScaledVec(previousState.RowView(0), h, tempState)
	slog.DebugContext(ctx, "Computed next state", slog.Any("nextState", tempState.RawVector().Data))

	nextState.SetRow(1, tempState.RawVector().Data)

	// x_{i+1}
	nextTime := initialTime + h

	slog.InfoContext(ctx, "Computed next state",
		slog.Float64("nextTime", nextTime),
		slog.Any("nextState", nextState.RawRowView(1)),
	)
	return NewEulerResult(nextTime,
		mat.NewVecDense(c, nextState.RawRowView(1)),
	)
}
