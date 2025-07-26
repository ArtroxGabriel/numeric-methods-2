package predictorcorrector

import (
	"context"
	"log/slog"

	rungekutta "github.com/ArtroxGabriel/numeric-methods-2/unidade4/runge-kutta"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

type FourthOrder struct {
	rk rungekutta.RungeKuttaInterface
}

func NewFourthOrder() *FourthOrder {
	return &FourthOrder{
		rk: rungekutta.NewRungeKuttaFour(),
	}
}

func (ab *FourthOrder) PredictorCorrector() {}

func (ab *FourthOrder) Execute(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialVec *mat.VecDense,
	startTime float64,
	interval float64,
	errorTolerance float64,
) *PredictorCorrectorResult {
	slog.InfoContext(ctx, "Executing PredictorCorrector method",
		slog.Uint64("order", uint64(4)),
		slog.Any("initialCondition", initialVec.RawVector().Data),
		slog.Float64("initialTime", startTime),
		slog.Float64("h", interval),
		slog.Float64("tolerance", errorTolerance),
	)
	interval /= 4

	vecLen := initialVec.Len()

	// Initialization phase

	// S_i
	states := mat.NewDense(5, vecLen, nil)
	states.SetRow(0, initialVec.RawVector().Data)

	// calculate S_{i+1}
	rkResult := ab.rk.Execute(ctx, fc, initialVec, startTime, interval)
	states.SetRow(1, rkResult.State.RawVector().Data)

	// calculate S_{i+2}
	tmpStates := mat.NewVecDense(initialVec.Len(), states.RawRowView(1))
	rkResult = ab.rk.Execute(ctx, fc, tmpStates, startTime+interval, interval)
	states.SetRow(2, rkResult.State.RawVector().Data)

	// calculate S_{i+3}
	tmpStates.CopyVec(states.RowView(2))
	rkResult = ab.rk.Execute(ctx, fc, tmpStates, startTime+interval*2, interval)
	states.SetRow(3, rkResult.State.RawVector().Data)

	// Prediction Phase
	// S_{i+4}_hat = S_{i+3} + ∆t/24 * (-9F_i + 37F_{i+1} - 59F_{i+2} + 55F_{i+3})
	slog.DebugContext(ctx, "Starting prediction phase")
	// calculate F_{i}
	derivPrev := fc(ctx, states, 0, startTime)

	// calculate F_{i+1}
	derivCurr := fc(ctx, states, 1, startTime+interval)

	// calculate F_{i+2}
	deriv2Curr := fc(ctx, states, 2, startTime+interval*2)

	deriv3Curr := fc(ctx, states, 3, startTime+interval*3)

	// calculate S_{i+4}_hat
	predictedNext := mat.NewVecDense(vecLen, nil)

	predictedNext.ScaleVec(-9, derivPrev)
	predictedNext.AddScaledVec(predictedNext, 37, derivCurr)
	predictedNext.AddScaledVec(predictedNext, -59, deriv2Curr)
	predictedNext.AddScaledVec(predictedNext, 55, deriv3Curr)
	predictedNext.AddScaledVec(states.RowView(3), interval/24, predictedNext)
	slog.DebugContext(ctx, "Predicted S_{i+4}_hat state",
		slog.Any("S_{i+4}_hat", predictedNext.RawVector().Data),
	)

	states.SetRow(4, predictedNext.RawVector().Data)

	// Correction Phase
	// S_{i+4} = S_{i+3} + ∆t/24 * (9 F_{i+4} + 19 F_{i+3} - 5 F_{i+2} + F_{i+1})
	slog.DebugContext(ctx, "Starting correction phase")
	for iter, errVal := 0, 1.0; iter < MaxIterations && errVal > errorTolerance; iter++ {
		// F_{i+4}
		derivNext := fc(ctx, states, 4, startTime+interval*4)

		// 9 F_{i+4} + 19 F_{i+3} - 5 F_{i+2} + F_{i+1}
		derivNext.ScaleVec(9, derivNext)
		derivNext.AddScaledVec(derivNext, 19, deriv3Curr)
		derivNext.AddScaledVec(derivNext, -5, deriv2Curr)
		derivNext.AddVec(derivNext, derivCurr)

		derivNext.AddScaledVec(states.RowView(3), interval/24, derivNext)

		prevState := states.RawRowView(4)

		errVal = calculateError(prevState, derivNext)

		// update the state with the new correction
		states.SetRow(4, derivNext.RawVector().Data)
	}

	nextTime := startTime + interval*4

	slog.InfoContext(ctx, "Computed the S_{i+4} state",
		slog.Float64("nextTime", nextTime),
		slog.Any("nextState", states.RawRowView(4)),
	)
	return NewPredictorCorrectorResult(nextTime, mat.NewVecDense(vecLen, states.RawRowView(4)))
}
