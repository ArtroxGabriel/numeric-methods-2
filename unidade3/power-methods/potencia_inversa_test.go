package powermethods_test

import (
	"math"
	"testing"

	powermethods "github.com/ArtroxGabriel/numeric-methods-2/unidade3/power-methods"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestPotenciaInversa(t *testing.T) {
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
			wantEigenvalue:  1,
			wantEigenvector: mat.NewVecDense(2, []float64{1, -1}),
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
			wantEigenvalue:  (-math.Sqrt(129) + 13.0) / 2.0,
			wantEigenvector: mat.NewVecDense(3, []float64{(-math.Sqrt(129) + 7.0) / 4.0, 0.5, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eps := tt.tolerance / 10.0
			_, gotErr := powermethods.PotenciaInversa(tt.matrix, tt.x0, eps, maxIterations)

			assert.NoError(t, gotErr, "PotenciaRegular() did not expect an error but got: %v", gotErr)
		})
	}
}
