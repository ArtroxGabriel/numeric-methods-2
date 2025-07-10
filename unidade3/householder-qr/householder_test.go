package householderqr_test

import (
	"testing"

	householderqr "github.com/ArtroxGabriel/numeric-methods-2/unidade3/householder-qr"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestHouseholderMethod(t *testing.T) {
	tests := []struct {
		name      string
		A         *mat.Dense
		tolerance float64
	}{
		{
			name: "Matriz simétrica 3x3",
			A: mat.NewDense(3, 3, []float64{
				6, 4, 1,
				4, 6, 1,
				1, 1, 5,
			}),
			tolerance: 1e-9,
		},
		{
			name: "matrix 5x5",
			A: mat.NewDense(5, 5, []float64{
				40, 8, 4, 2, 1,
				8, 30, 12, 6, 2,
				4, 12, 20, 1, 2,
				2, 6, 1, 25, 4,
				1, 2, 2, 4, 5,
			}),
			tolerance: 1e-5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := householderqr.HouseholderMethod(tt.A)
			T := result.T
			H := result.H
			n, _ := tt.A.Dims()

			temp := new(mat.Dense)
			AReconstructed := new(mat.Dense)
			temp.Mul(H, T)
			AReconstructed.Mul(temp, H.T())

			assert.True(
				t,
				mat.EqualApprox(AReconstructed, tt.A, tt.tolerance),
				"A matriz reconstruída A = H*T*H^T deve ser igual à original",
			)

			HTH := new(mat.Dense)
			HTH.Mul(H.T(), H)
			I := householderqr.NewIdentityMatrix(n)

			assert.True(t,
				mat.EqualApprox(HTH, I, tt.tolerance),
				"A matriz H deve ser ortogonal (H^T * H = I)",
			)
		})
	}
}
