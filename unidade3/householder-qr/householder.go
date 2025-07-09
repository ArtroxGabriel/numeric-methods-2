package householderqr

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type HouseholderResult struct {
	T *mat.Dense
	H *mat.Dense
}

// NewIdentityMatrix cria uma matrix identdade r x c, where r = c
func NewIdentityMatrix(r, c int) *mat.Dense {
	h := mat.NewDense(r, c, nil)
	for i := range r {
		h.Set(i, i, 1.0)
	}
	return h
}

// householderMatrix cria uma matriz de Householder para a coluna i de A
func householderMatrix(A *mat.Dense, i int) *mat.Dense {
	n, _ := A.Dims()
	I := NewIdentityMatrix(n, n)

	w := mat.NewVecDense(n, nil)
	wHat := mat.NewVecDense(n, nil)

	col := A.ColView(i)
	for j := i + 1; j < n; j++ {
		w.SetVec(j, col.AtVec(j))
	}

	Lw := mat.Norm(w, 2)

	if i+1 < n {
		wHat.SetVec(i+1, Lw)
	}

	N := mat.NewVecDense(n, nil)
	N.SubVec(w, wHat)

	normN := mat.Norm(N, 2)
	nVec := mat.NewVecDense(n, nil)
	if normN != 0 {
		nVec.ScaleVec(1/normN, N)
	}

	nVecT := nVec.T()
	outerProd := mat.NewDense(n, n, nil)
	outerProd.Mul(nVec, nVecT)

	hMatrix := mat.NewDense(n, n, nil)
	hMatrix.Scale(-2, outerProd)
	hMatrix.Add(I, hMatrix)

	return hMatrix
}

// HouseholderMethod aplica o método de Householder para triangularizar a matriz A
func HouseholderMethod(A *mat.Dense) HouseholderResult {
	n, _ := A.Dims()
	H := NewIdentityMatrix(n, n)

	AOld := mat.DenseCopyOf(A)
	ANew := mat.NewDense(n, n, nil)

	for i := range n - 2 {
		// householderMatrix para a coluna i de A
		Hi := householderMatrix(AOld, i)

		// A_k+1 = H_k * A_k
		temp := mat.NewDense(n, n, nil)
		temp.Mul(AOld, Hi)
		ANew.Mul(Hi.T(), temp)

		// AOld = ANew
		AOld.Copy(ANew)

		// Atualiza a matriz de Householder
		H.Mul(H, Hi)
	}

	return HouseholderResult{T: AOld, H: H}
}

func PrintMatrix(m mat.Matrix) {
	fa := mat.Formatted(
		m,
		mat.Prefix(""),
		mat.Squeeze(),
	)
	fmt.Printf("%.6f\n", fa)
}
