package predictorcorrector

import (
	"context"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

type PredictorCorrectorInterface interface {
	PredictorCorrector()
	Execute(
		ctx context.Context,
		fc types.DerivativeFunc,
		initialCondition *mat.VecDense,
		initialTime, h, tolerance float64,
	) *PredictorCorrectorResult
}

type PredictorCorrectorResult struct {
	Time  float64
	State *mat.VecDense
}

func NewPredictorCorrectorResult(time float64, state *mat.VecDense) *PredictorCorrectorResult {
	return &PredictorCorrectorResult{
		Time:  time,
		State: state,
	}
}
