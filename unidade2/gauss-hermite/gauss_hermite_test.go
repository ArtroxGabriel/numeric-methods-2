package gausshermite_test

import (
	"fmt"
	"math"
	"testing"

	gausshermite "github.com/ArtroxGabriel/numeric-methods-2/unidade2/gauss-hermite"
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
		expectedArea: math.Sqrt(math.Pi),
		tolerance:    1e-10,
	},
	{
		name:         "x² (even polynomial)",
		expr:         func(x float64) float64 { return x * x },
		expectedArea: math.Sqrt(math.Pi) / 2.0,
		tolerance:    1e-10,
	},
	{
		name:         "x⁶ (even polynomial)",
		expr:         func(x float64) float64 { return x * x * x * x * x * x },
		expectedArea: 15.0 * math.Sqrt(math.Pi) / 8.0,
		tolerance:    4,
	},
	{
		name:         "x (odd polynomial)",
		expr:         func(x float64) float64 { return x },
		expectedArea: 0.0,
		tolerance:    1e-10,
	},
	{
		name:         "x³ (odd polynomial)",
		expr:         func(x float64) float64 { return x * x * x },
		expectedArea: 0.0,
		tolerance:    1e-10,
	},
	{
		name:         "x⁵ (odd polynomial)",
		expr:         func(x float64) float64 { return x * x * x * x * x },
		expectedArea: 0.0,
		tolerance:    1e-10,
	},
	{
		name:         "e^(-x²) (Gaussian)",
		expr:         func(x float64) float64 { return math.Exp(-x * x) },
		expectedArea: math.Sqrt(math.Pi) / math.Sqrt(2.0),
		tolerance:    0.2,
	},
	{
		name:         "cos(x)",
		expr:         func(x float64) float64 { return math.Cos(x) },
		expectedArea: math.Sqrt(math.Pi) * math.Exp(-0.25),
		tolerance:    1e-1,
	},
	{
		name:         "sin(x)",
		expr:         func(x float64) float64 { return math.Sin(x) },
		expectedArea: 0.0,
		tolerance:    1e-10,
	},
}

func TestIntegrateHermite(t *testing.T) {

	calculators := []struct {
		name       string
		calculator gausshermite.GaussHermiteCalculator
	}{
		{"2 points", gausshermite.NewTwoPoints()},
		{"3 points", gausshermite.NewThreePoints()},
		{"4 points", gausshermite.NewFourPoints()},
	}

	for _, calc := range calculators {
		for _, tt := range testCases {
			testName := fmt.Sprintf("%s/%s", calc.name, tt.name)

			t.Run(testName, func(t *testing.T) {
				result := gausshermite.Integrate(calc.calculator, tt.expr)

				assert.InDelta(t, tt.expectedArea, result, tt.tolerance)
			})
		}
	}
}
