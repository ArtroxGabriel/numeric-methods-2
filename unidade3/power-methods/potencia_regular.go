package powermethods

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// PotenciaRegular calcula o maior autovalor e o autovetor correspondente de uma matriz A usando o método de potência regular.
func PotenciaRegular(
	matrixA *mat.Dense,
	initialVector *mat.VecDense,
	tolerance float64,
	maxIterations uint64,
) (*PowerMethodResult, error) {
	r, c := matrixA.Dims()

	if r != c {
		return nil, fmt.Errorf("matrix must be square")
	}

	if r == 0 {
		return nil, fmt.Errorf("matrix cannot be empty")
	}

	if c != initialVector.Len() {
		return nil, fmt.Errorf("the dimensions of the matrix and the vector are incompatible")
	}

	x := mat.NewVecDense(initialVector.Len(), nil)
	x.CopyVec(initialVector)

	y := mat.NewVecDense(initialVector.Len(), nil)

	var oldEigenvalue float64
	var eigenvalue float64

	for i := range maxIterations {
		y.MulVec(matrixA, x)

		// calcula o autovalor como a norma infinita do vetor y
		eigenvalue = mat.Norm(y, math.Inf(0))

		if eigenvalue == 0 {
			return nil, fmt.Errorf("the eigenvalue is zero, it's not possible to normalize the eigenvector")
		}

		// normaliza o vetor y
		x.ScaleVec(1/eigenvalue, y)

		// verifica a convergência
		if i > 0 && math.Abs((eigenvalue-oldEigenvalue)/eigenvalue) < tolerance {
			return &PowerMethodResult{
				Eigenvalue:  eigenvalue,
				Eigenvector: x,
			}, nil
		}

		oldEigenvalue = eigenvalue
	}

	return nil, fmt.Errorf("power iteration did not converge within %d iterations", maxIterations)
}
