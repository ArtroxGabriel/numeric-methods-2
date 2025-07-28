package pvcprocessor_test

import (
	"context"
	"testing"

	pvcprocessor "github.com/ArtroxGabriel/numeric-methods-2/unidade5/pvc-processor"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade5/result"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

const errorTolerance = 75e-5

func TestPVC_Perfom(t *testing.T) {
	tests := []struct {
		name    string
		input   *result.PVCInput
		want    *result.PVCResult
		wantErr bool
	}{
		{
			name: "PVC-1",
			input: &result.PVCInput{
				MaskValues: func() []float64 {
					maskFunctions := []func(float64) float64{
						func(dx float64) float64 {
							return 1.0 / (dx * dx)
						},
						func(dx float64) float64 {
							return (-2.0 / (dx * dx)) - 1.0
						},
						func(dx float64) float64 {
							return 1.0 / (dx * dx)
						},
					}
					maskValues := make([]float64, len(maskFunctions))
					for i, maskFunction := range maskFunctions {
						maskValues[i] = maskFunction(0.25) // Using a step size of 0.25
					}
					return maskValues
				}(),
				A:            0,
				B:            1,
				InitialCond:  []float64{0, 1},
				StepSize:     0.25,
				DefaultValue: 0,
			},
			want: &result.PVCResult{
				Solution: mat.NewDense(3, 1, []float64{
					0.21495513,
					0.44340944,
					0.69972421,
				}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pvcprocessor.NewPVC()
			got, gotErr := p.Perfom(context.Background(), tt.input)
			assert.NoError(t, gotErr, "Perfom() should not return an error")

			checkSolution(t, tt.want.Solution, got.Solution)
		})
	}
}

func checkSolution(t *testing.T, expected, got *mat.Dense) {
	expectedRows, _ := expected.Dims()
	gotRows, _ := got.Dims()
	assert.Equal(t, expectedRows, gotRows, "The number of rows in the expected and got matrices should be equal")

	for i := range expectedRows {
		expectedValue := expected.At(i, 0)
		gotValue := got.At(i, 0)
		relativeError := (expectedValue - gotValue) / expectedValue
		assert.InDelta(t, 0, relativeError, errorTolerance, "The relative error at row %d should be within the acceptable range", i)
	}
}
