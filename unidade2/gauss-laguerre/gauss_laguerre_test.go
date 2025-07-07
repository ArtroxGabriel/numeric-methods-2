package gausslaguerre_test

import (
	"fmt"
	"math"
	"testing"

	gausslaguerre "github.com/ArtroxGabriel/numeric-methods-2/unidade2/gauss-laguerre"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name         string
	expectedArea float64
	tolerance    float64
	expr         func(float64) float64
}{
	{
		name:         "1 (constant)",
		expr:         func(x float64) float64 { return 1.0 },
		expectedArea: 1.0,
		tolerance:    1e-10,
	},
	{
		name:         "x (linear)",
		expr:         func(x float64) float64 { return x },
		expectedArea: 1.0,
		tolerance:    1e-10,
	},
	{
		name:         "x² (quadratic)",
		expr:         func(x float64) float64 { return x * x },
		expectedArea: 2.0,
		tolerance:    1e-10,
	},
	{
		name:         "x³ (cubic)",
		expr:         func(x float64) float64 { return x * x * x },
		expectedArea: 6.0,
		tolerance:    1e-10,
	},
	{
		name:         "x⁴ (quartic)",
		expr:         func(x float64) float64 { return x * x * x * x },
		expectedArea: 24.0,
		tolerance:    1e-8,
	},
	{
		name:         "x⁵ (quintic)",
		expectedArea: 120.0,
		tolerance:    1e-6,
		expr:         func(x float64) float64 { return x * x * x * x * x },
	},
	{
		name:         "e^(-x) (exponential)",
		expr:         func(x float64) float64 { return math.Exp(-x) },
		expectedArea: 0.5,
		tolerance:    1e-1,
	},
	{
		name:         "sin(x)",
		expr:         func(x float64) float64 { return math.Sin(x) },
		expectedArea: 0.5,
		tolerance:    1e-1,
	},
	{
		name:         "cos(x)",
		expr:         func(x float64) float64 { return math.Cos(x) },
		expectedArea: 0.5,
		tolerance:    1e-1,
	},
}

func TestIntegrateHermite(t *testing.T) {

	calculators := []struct {
		name       string
		calculator gausslaguerre.GaussLaguerreCalculator
	}{
		{"2 points", gausslaguerre.NewTwoPoints()},
		{"3 points", gausslaguerre.NewThreePoints()},
		{"4 points", gausslaguerre.NewFourPoints()},
	}

	for _, calc := range calculators {
		for _, tt := range testCases {
			testName := fmt.Sprintf("%s/%s", calc.name, tt.name)

			t.Run(testName, func(t *testing.T) {
				result := gausslaguerre.Integrate(calc.calculator, tt.expr)

				assert.InDelta(t, tt.expectedArea, result, tt.tolerance)
			})
		}
	}
}
