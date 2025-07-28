package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	pvcprocessor "github.com/ArtroxGabriel/numeric-methods-2/unidade5/pvc-processor"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade5/result"
	"gonum.org/v1/gonum/mat"
)

func init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logDir := "log"
	if stat, err := os.Stat(logDir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(logDir, 0755); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else if !stat.IsDir() {
		panic("log exists but is not a directory")
	}
	timestamp := time.Now().Format("20060102_150405")
	logFileName := fmt.Sprintf("%s/pvc_%s.log", logDir, timestamp)
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	handler := slog.NewJSONHandler(logFile, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func main() {
	ctx := context.Background()

	initialCond := []float64{
		10, // x = 0
		1,  // x = 2
	}

	maskFunctions := []func(float64) float64{
		func(dx float64) float64 {
			return ((1.0 / (dx * dx)) - (7.0 / (2.0 * dx)))
		},
		func(dx float64) float64 {
			return ((-2.0 / (dx * dx)) - 1.0)
		},
		func(dx float64) float64 {
			return ((1.0 / (dx * dx)) + (7.0 / (2.0 * dx)))
		},
	}
	maskValues := make([]float64, len(maskFunctions))
	for i, maskFunction := range maskFunctions {
		maskValues[i] = maskFunction(0.1) // Using a step size of 0.1
	}

	pvcInput := &result.PVCInput{
		MaskValues:   maskValues,
		A:            0.0,
		B:            2.0,
		StepSize:     0.1,
		InitialCond:  initialCond,
		DefaultValue: 2.0,
	}

	pvc := pvcprocessor.NewPVC()

	pvcResult, err := pvc.Perfom(ctx, pvcInput)
	if err != nil {
		return
	}

	// fmt.Printf("Sistemas de equação:\n%.0f\n", mat.Formatted(pvcResult.Matrix))
	// fmt.Printf("Vetor B:\n%.0f\n", mat.Formatted(pvcResult.Vector))
	// fmt.Printf("Vetor solução:\n%.6f", mat.Formatted(pvcResult.Solution))

	printSolution(initialCond, pvcResult.Solution)
}

func printSolution(contourValues []float64, solution *mat.Dense) {
	// A solution vector is a matrix where the number of rows
	// represents the number of elements and the number of columns is 1.
	rows, _ := solution.Dims()

	// Iterate through each row of the vector to print the element.
	fmt.Printf("%.6f\n", contourValues[0])
	for i := range rows {
		// Access the element at row 'i' and column 0.
		fmt.Printf("%.6f\n", solution.At(i, 0))
	}
	fmt.Printf("%.6f\n", contourValues[1])
}
