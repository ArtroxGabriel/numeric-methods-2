package powermethods_test

import (
	powermethods "potencia/power-methods"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestPotenciaRegular(t *testing.T) {
	t.Parallel()

	maxIterations := uint64(100)

	tests := []struct {
		name            string
		matrix          *mat.Dense
		x0              *mat.VecDense
		tolerance       float64
		wantEigenvalue  float64
		wantEigenvector *mat.VecDense
	}{
		{
			name:            "matrix 2x2",
			matrix:          mat.NewDense(2, 2, []float64{2, 3, 5, 4}),
			x0:              mat.NewVecDense(2, []float64{1, 1}),
			tolerance:       1e-5,
			wantEigenvalue:  7,
			wantEigenvector: mat.NewVecDense(2, []float64{3.0 / 5.0, 1.0}),
		},
		{
			name: "matrix 3x3",
			matrix: mat.NewDense(3, 3, []float64{
				0, 2, 4,
				1, 1, -2,
				-2, 0, 5,
			}),
			x0:              mat.NewVecDense(3, []float64{1, 1, 1}),
			tolerance:       1e-5,
			wantEigenvalue:  3,
			wantEigenvector: mat.NewVecDense(3, []float64{-1, 0.5, -1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.tolerance / 10.0
			got, gotErr := powermethods.PotenciaRegular(tt.matrix, tt.x0, e, maxIterations)

			assert.NoError(t, gotErr, "PotenciaRegular() did not expect an error but got: %v", gotErr)

			assert.InDeltaf(t, tt.wantEigenvalue, got.Eigenvalue, tt.tolerance,
				"PotenciaRegular() for %s: got eigenvalue = %v, want %v", tt.name, got.Eigenvalue, tt.wantEigenvalue)

			assert.Equal(t, tt.wantEigenvector.Len(), got.Eigenvector.Len(),
				"PotenciaRegular() for %s: eigenvector length mismatch, got %d, want %d", tt.name, got.Eigenvector.Len(), tt.wantEigenvector.Len())

			for i := 0; i < tt.wantEigenvector.Len(); i++ {
				assert.InDeltaf(t, tt.wantEigenvector.AtVec(i), got.Eigenvector.AtVec(i), tt.tolerance,
					"PotenciaRegular() for %s: eigenvector component at index %d mismatch, got %v, want %v",
					tt.name, i, got.Eigenvector.AtVec(i), tt.wantEigenvector.AtVec(i))
			}
		})
	}
}
