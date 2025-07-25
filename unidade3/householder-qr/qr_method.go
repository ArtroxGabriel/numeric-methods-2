// Package householderqr implements the QR method for finding eigenvalues and eigenvectors of a matrix using Householder transformations.
package householderqr

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type QRResult struct {
	Lambda *mat.Dense
	X      *mat.Dense
}

// QRDecomp perfomar a decomposição QR de uma matriz A usando transformações de Householder.
func QRDecomp(A *mat.Dense) (*mat.Dense, *mat.Dense) {
	n, _ := A.Dims()
	R := mat.DenseCopyOf(A)
	Q := NewIdentityMatrix(n)

	for j := range n {
		for i := n - 1; i > j; i-- {
			a := R.At(i-1, j)
			b := R.At(i, j)

			r := math.Sqrt(a*a + b*b)
			if r < 1e-15 {
				continue
			}

			cos := a / r
			sen := -b / r

			G := NewIdentityMatrix(n)
			G.Set(i-1, i-1, cos)
			G.Set(i, i, cos)
			G.Set(i-1, i, -sen)
			G.Set(i, i-1, sen)

			// R = G * R
			R.Mul(G, R)

			// Q = Q * G^T
			Q.Mul(Q, G.T())
		}
	}

	return Q, R
}

// QRMethod aplica o método QR para encontrar os autovalores e autovetores de uma matriz T.
func QRMethod(T, H *mat.Dense, epsilon float64) QRResult {
	X := mat.DenseCopyOf(H)
	A := mat.DenseCopyOf(T)
	n, _ := A.Dims()

	error := 1.0

	for error > epsilon {
		// QRDecomp em A
		Q, R := QRDecomp(A)

		// A_k+1 = R_k * Q_k
		A.Mul(R, Q)

		// Acumula os autovetores
		X.Mul(X, Q)

		error = 0.0
		for j := range n - 1 {
			for i := j + 1; i < n; i++ {
				error += math.Abs(A.At(i, j))
			}
		}
	}

	return QRResult{Lambda: A, X: X}
}
