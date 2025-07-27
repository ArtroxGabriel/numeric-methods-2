// Package result defines the PVCResult type
package result

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

var ErrInvalidStepSize = errors.New("step size must be greater than 1")

type PVCInput struct {
	MaskValues   []float64
	A, B         float64   // interval limits
	StepSize     float64   // step size
	InitialCond  []float64 // initial conditions for the boundary values
	DefaultValue float64   // default value for b vector
}

type PVCResult struct {
	Matrix   *mat.Dense
	Vector   *mat.VecDense
	Solution *mat.Dense
}

func NewPVCResult(
	matrix *mat.Dense,
	vector *mat.VecDense,
	result *mat.Dense,
) *PVCResult {
	return &PVCResult{
		Matrix:   matrix,
		Vector:   vector,
		Solution: result,
	}
}
