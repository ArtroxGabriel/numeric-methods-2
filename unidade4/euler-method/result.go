package eulermethod

import "gonum.org/v1/gonum/mat"

type EulerResult struct {
	Time  float64
	State *mat.VecDense
}

func NewEulerResult(time float64, data *mat.VecDense) *EulerResult {
	return &EulerResult{
		Time:  time,
		State: data,
	}
}
