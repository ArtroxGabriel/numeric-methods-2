package dino_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade2/dino"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name         string
	expectedArea float64
	tolerance    float64
	a            float64
	b            float64
	expr         func(float64) float64
}{
	{
		name:         "1 over sqrt(x)",
		a:            0,
		b:            1,
		expr:         func(x float64) float64 { return 1.0 / math.Sqrt(x) },
		expectedArea: 2,
		tolerance:    1e-1,
	},
}

func TestIntegrateDino(t *testing.T) {
	calculators := []struct {
		name       string
		calculator dino.DinoCalculator
	}{
		{"dino simples", dino.NewDinoSimples()},
		{"dino duo", dino.NewDinoDuo()},
	}

	for _, calc := range calculators {
		for _, tc := range testCases {
			testName := fmt.Sprintf("%s/%s", calc.name, tc.name)

			t.Run(testName, func(t *testing.T) {
				result := dino.IntegrateDino(calc.calculator, tc.expr, tc.a, tc.b)

				assert.InDelta(t, tc.expectedArea, result.Result, tc.tolerance)

				t.Logf("Número de iterações: %d", result.NumOfIterations)
			})
		}
	}
}
