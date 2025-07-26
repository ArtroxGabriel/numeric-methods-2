// Package predictorcorrector implements the predictor-corrector method for solving ordinary differential equations.
package predictorcorrector

import (
	"context"
	"log/slog"
	"math"

	rungekutta "github.com/ArtroxGabriel/numeric-methods-2/unidade4/runge-kutta"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

const MaxIterations = 100

var _ PredictorCorrectorInterface = (*AdamsBashforth)(nil)

type AdamsBashforth struct {
	rk rungekutta.RungeKuttaInterface
}

func NewAdamsBashforth() *AdamsBashforth {
	return &AdamsBashforth{
		rk: rungekutta.NewRungeKuttaSecond(),
	}
}

func (ab *AdamsBashforth) PredictorCorrector() {}

func (ab *AdamsBashforth) Execute(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialVec *mat.VecDense,
	startTime, stepSize float64,
	errorTolerance float64,
) *PredictorCorrectorResult {
	slog.InfoContext(ctx, "Executing Adams-Bashforth method",
		slog.Uint64("order", uint64(2)),
		slog.Any("initialCondition", initialVec.RawVector().Data),
		slog.Float64("initialTime", startTime),
		slog.Float64("h", stepSize),
		slog.Float64("tolerance", errorTolerance),
	)
	stepSize /= 2

	vecLen := initialVec.Len()

	// S_i
	states := mat.NewDense(3, vecLen, nil)
	states.SetRow(0, initialVec.RawVector().Data)

	// calculate S_{i+1}
	rkResult := ab.rk.Execute(ctx, fc, initialVec, startTime, stepSize)
	states.SetRow(1, rkResult.State.RawVector().Data)

	// Prediction Phase
	// S_{i+2}_hat = S_{i+1} + ∆t/2 * (-F_i + 3F_{i+1})
	slog.DebugContext(ctx, "Starting prediction phase")
	// calculate F_{i}
	derivPrev := fc(ctx, states, 0, startTime)

	// calculate F_{i+1}
	derivCurr := fc(ctx, states, 1, startTime+stepSize)

	// calculate S_{i+2}_hat
	predictedNext := mat.NewVecDense(vecLen, nil)
	predictedNext.ScaleVec(-1, derivPrev)
	predictedNext.AddScaledVec(predictedNext, 3, derivCurr)
	predictedNext.AddScaledVec(states.RowView(1), stepSize/2, predictedNext)
	slog.DebugContext(ctx, "Predicted S_{i+2}_hat state",
		slog.Any("S_{i+2}_hat", predictedNext.RawVector().Data),
	)

	states.SetRow(2, predictedNext.RawVector().Data)

	// Correction Phase
	// S_{i+2} = S_{i+1} + ∆t/2 * (F_{i+1} + F_{i+2})
	slog.DebugContext(ctx, "Starting correction phase")
	for iter, errVal := 0, 1.0; iter < MaxIterations; iter++ {
		// F_{i+2}
		derivNext := fc(ctx, states, 2, startTime+stepSize*2)

		// F_{i+1} + F_{i+2}
		derivNext.AddVec(derivCurr, derivNext)

		derivNext.AddScaledVec(states.RowView(1), stepSize/2, derivNext)

		prevState := states.RawRowView(2)

		errVal = calculateError(prevState, derivNext)

		// update the state with the new correction
		states.SetRow(2, derivNext.RawVector().Data)
		if errVal < errorTolerance {
			break
		}
	}

	nextTime := startTime + stepSize*2

	slog.InfoContext(ctx, "Computed the S_{i+2} state",
		slog.Float64("nextTime", nextTime),
		slog.Any("nextState", states.RawRowView(2)),
	)
	return NewPredictorCorrectorResult(nextTime, mat.NewVecDense(vecLen, states.RawRowView(2)))
}

func calculateError(prevState []float64, goalVec *mat.VecDense) float64 {
	tmpVec := mat.NewVecDense(len(prevState), prevState)
	prevNorm := tmpVec.Norm(2)
	goalNorm := goalVec.Norm(2)

	return math.Abs(goalNorm-prevNorm) / goalNorm
}
