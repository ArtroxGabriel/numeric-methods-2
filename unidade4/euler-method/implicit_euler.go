package eulermethod

import (
	"context"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

type ImplicitEuler struct {
	explicitEuler *ExplicitEuler
}

func NewImplicitEuler() *ImplicitEuler {
	return &ImplicitEuler{
		explicitEuler: NewExplicitEuler(),
	}
}

func (em *ImplicitEuler) getGuess(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *EulerResult {
	return em.explicitEuler.Execute(ctx, fc, initialCondition, initialTime, h)
}

func (em *ImplicitEuler) Execute(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *EulerResult {
	slog.InfoContext(ctx, "Executing Implicit Euler Method",
		slog.Any("initialCondition", initialCondition.RawVector().Data),
		slog.Float64("initialTime", initialTime),
		slog.Float64("h", h),
	)
	r, c := 1, initialCondition.Len()

	// return the guess nextStateHat(S_{i+1}) by this formula: S_i + Δt*F(S_i,t_i)
	guessResult := em.getGuess(ctx, fc, initialCondition, initialTime, h)

	nextStateHat := mat.NewDense(r, c, nil)
	nextStateHat.SetRow(0, guessResult.State.RawVector().Data)
	slog.DebugContext(ctx, "Next State S_{i+1}_Hat copied")

	nextState := mat.NewDense(r+1, c, nil)

	// F(S_{i+1},t_{i+1})
	tempState := fc(ctx, nextStateHat, 0)

	// S_i + Δt*F(S_{i+1},t_{i+1})
	tempState.AddScaledVec(initialCondition, h, tempState)
	slog.DebugContext(ctx, "Computed next state refined",
		slog.Any("nextState", tempState.RawVector().Data))

	nextState.SetRow(1, tempState.RawVector().Data)

	// x_{i+1}
	nextStateTime := initialTime + h

	slog.InfoContext(ctx, "Computed next state refined",
		slog.Float64("nextTime", nextStateTime),
		slog.Any("nextState", nextState.RawRowView(1)),
	)
	return NewEulerResult(nextStateTime,
		mat.NewVecDense(c, nextState.RawRowView(1)),
	)
}
