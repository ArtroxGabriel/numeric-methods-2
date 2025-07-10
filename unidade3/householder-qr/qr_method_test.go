package householderqr_test

import (
	"testing"

	householderqr "github.com/ArtroxGabriel/numeric-methods-2/unidade3/householder-qr"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestQRMethod(t *testing.T) {
	tests := []struct {
		name      string
		A         *mat.Dense
		tolerance float64
	}{
		{
			name: "Matriz simétrica 5x5",
			A: mat.NewDense(5, 5, []float64{
				40, 8, 4, 2, 1,
				8, 30, 12, 6, 2,
				4, 12, 20, 1, 2,
				2, 6, 1, 25, 4,
				1, 2, 2, 4, 5,
			}),
			tolerance: 1e-9,
		},
		{
			name: "Matriz simétrica 3x3",
			A: mat.NewDense(3, 3, []float64{
				6, 4, 1,
				4, 6, 1,
				1, 1, 5,
			}),
			tolerance: 1e-9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Executa o fluxo completo: Householder -> QR
			householderResult := householderqr.HouseholderMethod(tt.A)
			qrResult := householderqr.QRMethod(householderResult.T, householderResult.H, tt.tolerance)
			Lambda := qrResult.Lambda
			X := qrResult.X
			n, _ := tt.A.Dims()

			// 2. VERIFICAÇÃO DA RECONSTRUÇÃO (A ≈ X * Λ * X^T)
			temp := new(mat.Dense)
			AReconstructed := new(mat.Dense)

			temp.Mul(X, Lambda)
			AReconstructed.Mul(temp, X.T())

			assert.True(t,
				mat.EqualApprox(AReconstructed, tt.A, tt.tolerance),
				"A matriz reconstruída A = X*Λ*X^T deve ser igual à original",
			)

			// 3. VERIFICAÇÃO DA DEFINIÇÃO (A*v ≈ λ*v para cada autovetor)
			for i := range n {
				lambda := Lambda.At(i, i) // Autovalor
				v := X.ColView(i)         // Autovetor (coluna i da matriz X)

				var Av mat.VecDense
				Av.MulVec(tt.A, v) // Lado esquerdo: A*v

				lambdaV := new(mat.VecDense)
				lambdaV.ScaleVec(lambda, v) // Lado direito: λ*v

				// Compara os dois vetores resultantes
				assert.True(t,
					mat.EqualApprox(&Av, lambdaV, tt.tolerance),
					"A*v deve ser igual a λ*v para o autovalor/vetor %d",
					i,
				)
			}
		})
	}
}
