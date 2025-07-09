// Package powermethods implements the power method for finding the smallest eigenvalue and corresponding eigenvector of a matrix using the inverse power method.
package powermethods

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// PotenciaInversa calcula o menor autovalor e o autovetor correspondente de uma matriz A usando o método de potência inversa.
func PotenciaInversa(
	a *mat.Dense,
	x0 *mat.VecDense,
	tolerance float64,
	maxIterations int,
) (*PowerMethodResult, error) {
	r, c := a.Dims()
	matrixInverse := mat.NewDense(r, c, nil)
	// inverter a matriz A
	if err := matrixInverse.Inverse(a); err != nil {
		return nil, fmt.Errorf("could not compute inverse matrix: %v", err)
	}

	// calcula o menor autovalor e autovetor correspondente da matriz inversa
	regularResult, err := PotenciaRegular(matrixInverse, x0, tolerance, maxIterations)
	if err != nil {
		return nil, fmt.Errorf("power iteration on inverse matrix failed: %v", err)
	}

	// ajusta o autovalor para o autovalor da matriz original
	eigenvalue := 1.0 / regularResult.Eigenvalue

	return &PowerMethodResult{
		Eigenvalue:  eigenvalue,
		Eigenvector: regularResult.Eigenvector,
	}, nil
}
