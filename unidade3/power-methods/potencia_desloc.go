package powermethods

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// PotenciaDeslocamento calcula o autovalor e autovetor de uma matriz A usando o método de potência deslocada.
func PotenciaDeslocamento(
	A *mat.Dense,
	v0 *mat.VecDense,
	tolerance float64,
	mu float64,
	maxIterations int,
) (*PowerMethodResult, error) {
	r, c := A.Dims()
	if r == 0 || r != c {
		return nil, fmt.Errorf("the matrix A must be square and non-empty")
	}

	AHat := mat.NewDense(r, c, nil)
	AHat.Copy(A)

	for i := range r {
		// desloca a matriz A subtraindo mu da diagonal
		currentVal := AHat.At(i, i)
		AHat.Set(i, i, currentVal-mu)
	}

	// menor autovalor e autovetor correspondente da matriz deslocada
	inverseResult, invErr := PotenciaInversa(AHat, v0, tolerance, maxIterations)
	if invErr != nil {
		return nil, fmt.Errorf("the shifted matrix A_hat failed in the inverse power method: %v", invErr)
	}

	// ajusta o autovalor para o autovalor da matriz original
	lambdaI := inverseResult.Eigenvalue + mu

	xI := inverseResult.Eigenvector

	return &PowerMethodResult{
		Eigenvalue:  lambdaI,
		Eigenvector: xI,
	}, nil
}
