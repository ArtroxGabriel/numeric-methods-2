package main

import (
	"fmt"
	householderqr "potencia/householder-qr"

	"gonum.org/v1/gonum/mat"
)

func main() {
	data := []float64{
		21, 29, 21, 32, 40,
		29, 94, 62, 87, 94,
		21, 62, 131, 90, 73,
		32, 87, 90, 94, 95,
		40, 94, 73, 95, 105,
	}
	// v := []float64{1, 1, 1, 1, 1}
	n := 5
	tolerance := 1e-6
	// maxIterations := 100
	// mu := 7.5
	A := mat.NewDense(n, n, data)
	// initialGuess := mat.NewVecDense(n, v)

	/*
		regularResult, err := powermethods.PotenciaRegular(A, initialGuess, tolerance, maxIterations)
		if err != nil {
			slog.Error("Error in Regular Power Method: ", slog.Any("Error", err))
		} else {
			slog.Info("Regular Power Method completed successfully")
			fmt.Println("eigenvalue:", regularResult.Eigenvalue)
			fmt.Printf("eigenvector: \n\t%v\t\n", mat.Formatted(regularResult.Eigenvector, mat.Prefix("\t")))

		}

		inverseResult, err := powermethods.PotenciaInversa(A, initialGuess, tolerance, maxIterations)
		if err != nil {
			slog.Error("Error in Inverse Power Method: ", slog.Any("Error", err))
		} else {
			slog.Info("Inverse Power Method completed successfully")
			fmt.Println("eigenvalue:", inverseResult.Eigenvalue)
			fmt.Printf("eigenvector: \n\t%v\t\n", mat.Formatted(inverseResult.Eigenvector, mat.Prefix("\t")))
		}

		deslocResult, err := powermethods.PotenciaDeslocamento(A, initialGuess, tolerance, mu, maxIterations)
		if err != nil {
			slog.Error("Error in Shifted Power Method: ", slog.Any("Error", err))
		} else {
			slog.Info("Shifted Power Method completed successfully")
			fmt.Println("eigenvalue:", deslocResult.Eigenvalue)
			fmt.Printf("eigenvector: \n\t%v\t\n",
				mat.Formatted(deslocResult.Eigenvector, mat.Prefix("\t")),
			)
		}
	*/

	fmt.Println("Matriz Original (A):")
	householderqr.PrintMatrix(A)

	fmt.Println("\n--- Método Householder QR ---")
	resultHouseholder := householderqr.HouseholderMethod(A)

	fmt.Println("\nMatriz Tridiagonal (T):")
	householderqr.PrintMatrix(resultHouseholder.T)

	fmt.Println("\nMatriz Accumulada (H):")
	householderqr.PrintMatrix(resultHouseholder.H)

	fmt.Println("\nAplicando o método QR iterativo...")
	resultQR := householderqr.QRMethod(resultHouseholder.T, resultHouseholder.H, tolerance)

	fmt.Println("\nMatriz de autovalores (Lambda):")
	householderqr.PrintMatrix(resultQR.Lambda)

	fmt.Println("\nMatriz de Autovetores (X):")
	householderqr.PrintMatrix(resultQR.X)

	fmt.Println("\nAutovalores (diagonal de Lambda):")
	diagonalView := resultQR.Lambda.DiagView()
	for i := range diagonalView.Diag() {
		fmt.Printf("λ%d = %.8f\n", i+1, diagonalView.At(i, i))
	}
}
