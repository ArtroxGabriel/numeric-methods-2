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
func NewIdentityMatrix(r int) *mat.Dense {
	h := mat.NewDense(r, r, nil)
	for i := range r {
		h.Set(i, i, 1.0)
	}
	return h
}

// householderMatrix cria uma matriz de Householder para a coluna i de A
func householderMatrix(A *mat.Dense, i int) *mat.Dense {
	n, _ := A.Dims()
	I := NewIdentityMatrix(n)

	w := mat.NewVecDense(n, nil)
	col := A.ColView(i)
	for j := i + 1; j < n; j++ {
		w.SetVec(j, col.AtVec(j))
	}

	Lw := w.Norm(2)
	wHat := mat.NewVecDense(n, nil)
	if i+1 < n {
		wHat.SetVec(i+1, Lw)
	}

	N := mat.NewVecDense(n, nil)
	// N = w - wHat
	N.SubVec(w, wHat)

	normN := N.Norm(2)
	nVec := mat.NewVecDense(n, nil)
	if normN >= 1e-12 {
		nVec.ScaleVec(1/normN, N)
	}

	outerProd := new(mat.Dense)
	outerProd.Outer(2, nVec, nVec)

	matrixH := new(mat.Dense)
	matrixH.Sub(I, outerProd)

	return matrixH
}

// HouseholderMethod aplica o método de Householder para triangularizar a matriz A
func HouseholderMethod(A *mat.Dense) HouseholderResult {
	n, _ := A.Dims()
	H := NewIdentityMatrix(n)

	AOld := mat.DenseCopyOf(A)
	ANew := mat.NewDense(n, n, nil)

	for i := range n - 2 {
		// householderMatrix para a coluna i de A
		Hi := householderMatrix(AOld, i)

		// A_k+1 = H_k * A_k
		temp := new(mat.Dense)
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
