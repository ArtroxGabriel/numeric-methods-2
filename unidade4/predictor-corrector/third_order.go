// Package predictorcorrector implements the predictor-corrector method for solving ordinary differential equations.
package predictorcorrector

import (
	"context"
	"log/slog"

	rungekutta "github.com/ArtroxGabriel/numeric-methods-2/unidade4/runge-kutta"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

var _ PredictorCorrectorInterface = (*ThirdOrder)(nil)

type ThirdOrder struct {
	rk rungekutta.RungeKuttaInterface
}

func NewThirdOrder() *ThirdOrder {
	return &ThirdOrder{
		rk: rungekutta.NewRungeKuttaThird(),
	}
}

func (ab *ThirdOrder) PredictorCorrector() {}

func (ab *ThirdOrder) Execute(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialVec *mat.VecDense,
	startTime, interval float64,
	errorTolerance float64,
) *PredictorCorrectorResult {
	slog.InfoContext(ctx, "Executing PredictorCorrector method",
		slog.Uint64("order", uint64(3)),
		slog.Any("initialCondition", initialVec.RawVector().Data),
		slog.Float64("initialTime", startTime),
		slog.Float64("h", interval),
		slog.Float64("tolerance", errorTolerance),
	)
	interval /= 3

	vecLen := initialVec.Len()

	// Initialization phase

	// S_i
	states := mat.NewDense(4, vecLen, nil)
	states.SetRow(0, initialVec.RawVector().Data)

	// calculate S_{i+1}
	rkResult := ab.rk.Execute(ctx, fc, initialVec, startTime, interval)
	states.SetRow(1, rkResult.State.RawVector().Data)

	// calculate S_{i+2}
	tmpStates := mat.NewVecDense(initialVec.Len(), states.RawRowView(1))
	rkResult = ab.rk.Execute(ctx, fc, tmpStates, startTime+interval, interval)
	states.SetRow(2, rkResult.State.RawVector().Data)

	// Prediction Phase
	// S_{i+3}_hat = S_{i+2} + ∆t/12 * (5F_i - 16F_{i+1} + 23F_{i+2})
	slog.DebugContext(ctx, "Starting prediction phase")
	// calculate F_{i}
	derivPrev := fc(ctx, states, 0, startTime)

	// calculate F_{i+1}
	derivCurr := fc(ctx, states, 1, startTime+interval)

	// calculate F_{i+2}
	deriv2Curr := fc(ctx, states, 2, startTime+interval*2)

	// calculate S_{i+2}_hat
	predictedNext := mat.NewVecDense(vecLen, nil)

	predictedNext.ScaleVec(5, derivPrev)
	predictedNext.AddScaledVec(predictedNext, -16, derivCurr)
	predictedNext.AddScaledVec(predictedNext, 23, deriv2Curr)
	predictedNext.AddScaledVec(states.RowView(2), interval/12, predictedNext)
	slog.DebugContext(ctx, "Predicted S_{i+3}_hat state",
		slog.Any("S_{i+3}_hat", predictedNext.RawVector().Data),
	)

	states.SetRow(3, predictedNext.RawVector().Data)

	// Correction Phase
	// S_{i+3} = S_{i+2} + ∆t/12 * (-F_{i+1} + 8F_{i+2} + 5F_{i+3})
	slog.DebugContext(ctx, "Starting correction phase")
	for iter, errVal := 0, 1.0; iter < MaxIterations && errVal > errorTolerance; iter++ {
		// F_{i+3}
		derivNext := fc(ctx, states, 3, startTime+interval*3)

		// -F_{i+1} + 8F_{i+2} + 5F_{i+3}
		derivNext.ScaleVec(5, derivNext)
		derivNext.AddScaledVec(derivNext, 8, deriv2Curr)
		derivNext.AddScaledVec(derivNext, -1, derivCurr)

		derivNext.AddScaledVec(states.RowView(2), interval/12, derivNext)

		prevState := states.RawRowView(3)

		errVal = calculateError(prevState, derivNext)

		// update the state with the new correction
		states.SetRow(3, derivNext.RawVector().Data)
	}

	nextTime := startTime + interval*3

	slog.InfoContext(ctx, "Computed the S_{i+3} state",
		slog.Float64("nextTime", nextTime),
		slog.Any("nextState", states.RawRowView(3)),
	)
	return NewPredictorCorrectorResult(nextTime, mat.NewVecDense(vecLen, states.RawRowView(3)))
}
