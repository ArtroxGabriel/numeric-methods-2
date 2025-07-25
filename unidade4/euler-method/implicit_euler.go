package eulermethod

import (
	"context"
	"log/slog"
	"math"

	"gonum.org/v1/gonum/mat"
)

const (
	MaxIterations = 100
	eTolerance    = 1e-6
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
	fc DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *EulerResult {
	return em.explicitEuler.Execute(ctx, fc, initialCondition, initialTime, h)
}

func (em *ImplicitEuler) Execute(
	ctx context.Context,
	fc DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *EulerResult {
	slog.InfoContext(ctx, "Executing Implicit Euler Method",
		slog.Any("initialCondition", initialCondition.RawVector().Data),
		slog.Float64("initialTime", initialTime),
		slog.Float64("h", h),
	)
	r, c := 1, initialCondition.Len()

	// return the guess nextStateHat by this formula: S_i + Δt*F(S_i,t_i)
	guessResult := em.getGuess(ctx, fc, initialCondition, initialTime, h)

	nextStateHat := mat.NewDense(r, c, nil)
	nextStateHat.SetRow(0, guessResult.State.RawVector().Data)
	slog.InfoContext(ctx, "Next State Hat copied")

	nextState := mat.NewDense(r+1, c, nil)
	var err float64

	for iteration := 0; iteration < MaxIterations; {
		// F(S_{i+1},t_{i+1})
		tempState := fc(ctx, nextStateHat, 0)

		// S_i + Δt*F(S_{i+1},t_{i+1})
		tempState.AddScaledVec(initialCondition, h, tempState)
		slog.DebugContext(ctx, "Computed next state refined",
			slog.Any("nextState", tempState.RawVector().Data))

		err = calculateRelativeError(nextStateHat.RawRowView(0), tempState)
		if err < eTolerance {
			nextState.SetRow(1, tempState.RawVector().Data)
			break
		}
		nextStateHat.SetRow(0, tempState.RawVector().Data)
	}

	// x_{i+1}
	nextStateTime := initialTime + h

	slog.InfoContext(ctx, "Computed next state refined",
		slog.Float64("nextTime", nextStateTime),
		slog.Any("nextState", nextState.RawRowView(1)),
		slog.Float64("relative error", err),
	)
	return NewEulerResult(nextStateTime,
		mat.NewVecDense(c, nextState.RawRowView(1)),
	)
}

func calculateRelativeError(previousStateSlice []float64, nextState *mat.VecDense) float64 {
	previousState := mat.NewVecDense(len(previousStateSlice), previousStateSlice)
	return math.Abs(previousState.Norm(2)-nextState.Norm(2)) / previousState.Norm(2)
}
