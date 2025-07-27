// Package pvcprocessor(problema de valor de contorno) implements the solution for a boundary value problem using the finite difference method.
package pvcprocessor

import (
	"context"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade5/result"
	"gonum.org/v1/gonum/mat"
)

type PVC struct{}

func NewPVC() *PVC {
	return &PVC{}
}

// Perfom solves a one-dimensional boundary value problem using the finite difference method.
// The input must be a 1D array of mask values representing the problem domain.
func (p *PVC) Perfom(ctx context.Context, input *result.PVCInput) (*result.PVCResult, error) {
	slog.Info("Performing PVC calculation",
		slog.Any("maskValues", input.MaskValues),
		slog.Float64("a", input.A),
		slog.Float64("b", input.B),
		slog.Float64("stepSize", input.StepSize),
		slog.Any("initialCond", input.InitialCond),
		slog.Float64("defaultValue", input.DefaultValue),
	)
	if input.StepSize <= 0 {
		slog.ErrorContext(ctx, "step size must be greater than 0",
			slog.Float64("stepSize", input.StepSize),
		)
		return nil, result.ErrInvalidStepSize
	}

	n := int((input.B-input.A)/input.StepSize) - 1
	slog.DebugContext(ctx, "calculated the number of points to be used",
		slog.Int("n", n),
	)

	matrix := mat.NewDense(n, n, nil)
	vectorB := mat.NewVecDense(n, p.initBValue(input.DefaultValue, n))

	for i := range n {
		row := make([]float64, n)

		switch i {
		case 0:
			vectorB.SetVec(i, vectorB.AtVec(i)-input.InitialCond[0]*input.MaskValues[0])
		case n - 1:
			vectorB.SetVec(i, vectorB.AtVec(i)-input.InitialCond[1]*input.MaskValues[len(input.MaskValues)-1])
		}

		for j := range n {
			switch j {
			case i - 1: // previous
				row[j] = input.MaskValues[0]
			case i: // current
				row[j] = input.MaskValues[1]
			case i + 1: // next
				row[j] = input.MaskValues[2]
			default: // other values are 0
				row[j] = 0

			}
		}
		matrix.SetRow(i, row)
	}
	slog.DebugContext(ctx, "system of equations created")

	var solution mat.Dense

	if err := solution.Solve(matrix, vectorB); err != nil {
		slog.ErrorContext(ctx, "failed to solve the system of equations",
			slog.Any("error", err))
		return nil, err
	}

	slog.InfoContext(ctx, "system of equations solved successfully")

	res := result.NewPVCResult(matrix, vectorB, &solution)
	return res, nil
}

func (*PVC) initBValue(defaultValue float64, size int) []float64 {
	bValues := make([]float64, size)
	for i := range bValues {
		bValues[i] = defaultValue
	}
	return bValues
}
