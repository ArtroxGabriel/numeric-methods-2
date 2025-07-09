package gausschebyshev_test

import (
	"fmt"
	"math"
	"testing"

	gausschebyshev "github.com/ArtroxGabriel/numeric-methods-2/unidade2/gauss-chebyshev"
	"github.com/stretchr/testify/assert"
)

func TestIntegrateLaguerre_TwoPoints(t *testing.T) {
	testCases := []struct {
		name         string
		expectedArea float64
		tolerance    float64
		expr         func(float64) float64
	}{
		{
			name:         "1 (constant)",
			expr:         func(x float64) float64 { return 1.0 },
			expectedArea: math.Pi,
			tolerance:    1e-10,
		},
		{
			name:         "x (odd function)",
			expectedArea: 0.0,
			tolerance:    1e-10,
			expr:         func(x float64) float64 { return x },
		},
		{
			name:         "x² (even function)",
			expr:         func(x float64) float64 { return x * x },
			expectedArea: math.Pi / 2.0,
			tolerance:    1e-10,
		},
		{
			name:         "x³ (odd function)",
			expr:         func(x float64) float64 { return x * x * x },
			expectedArea: 0.0,
			tolerance:    1e-10,
		},
		{
			name:         "x⁵ (odd function)",
			expectedArea: 0.0,
			tolerance:    1e-10,
			expr:         func(x float64) float64 { return x * x * x * x * x },
		},
		{
			name:         "cos(x)",
			expectedArea: math.Pi * 0.7652,
			tolerance:    1e-1,
			expr:         func(x float64) float64 { return math.Cos(x) },
		},
		{
			name:         "sin(x)",
			expectedArea: 0.0,
			tolerance:    1e-10,
			expr:         func(x float64) float64 { return math.Sin(x) },
		},
		{
			name:         "1/(1+x²)",
			expectedArea: math.Pi / math.Sqrt(2.0),
			tolerance:    10e-1,
			expr:         func(x float64) float64 { return 1.0 / (1.0 + x*x) },
		},
	}

	for _, tt := range testCases {
		testName := fmt.Sprintf("%s/%s", "2 points", tt.name)

		calc := gausschebyshev.NewTwoPoints()
		t.Run(testName, func(t *testing.T) {
			result := gausschebyshev.Integrate(calc, tt.expr)

			assert.InDelta(t, tt.expectedArea, result, tt.tolerance)
		})
	}
}

func TestIntegrateLaguerre_ThreeAndFourPoints(t *testing.T) {
	calculators := []struct {
		name       string
		calculator gausschebyshev.GaussChebyshevCalculator
	}{
		{"3 points", gausschebyshev.NewThreePoints()},
		{"4 points", gausschebyshev.NewFourPoints()},
	}

	testCases := []struct {
		name         string
		expectedArea float64
		tolerance    float64
		expr         func(float64) float64
	}{
		{
			name:         "1 (constant)",
			expr:         func(x float64) float64 { return 1.0 },
			expectedArea: math.Pi,
			tolerance:    1e-10,
		},
		{
			name:         "x (odd function)",
			expectedArea: 0.0,
			tolerance:    1e-10,
			expr:         func(x float64) float64 { return x },
		},
		{
			name:         "x² (even function)",
			expr:         func(x float64) float64 { return x * x },
			expectedArea: math.Pi / 2.0,
			tolerance:    1e-10,
		},
		{
			name:         "x³ (odd function)",
			expr:         func(x float64) float64 { return x * x * x },
			expectedArea: 0.0,
			tolerance:    1e-10,
		},
		{
			name:         "x⁴ (even function)",
			expr:         func(x float64) float64 { return x * x * x * x },
			expectedArea: 3.0 * math.Pi / 8.0,
			tolerance:    1e-10,
		},
		{
			name:         "x⁵ (odd function)",
			expectedArea: 0.0,
			tolerance:    1e-10,
			expr:         func(x float64) float64 { return x * x * x * x * x },
		},
		{
			name:         "cos(x)",
			expectedArea: math.Pi * 0.7652,
			tolerance:    1e-1,
			expr:         func(x float64) float64 { return math.Cos(x) },
		},
		{
			name:         "sin(x)",
			expectedArea: 0.0,
			tolerance:    1e-10,
			expr:         func(x float64) float64 { return math.Sin(x) },
		},
		{
			name:         "1/(1+x²)",
			expectedArea: math.Pi / math.Sqrt(2.0),
			tolerance:    10e-1,
			expr:         func(x float64) float64 { return 1.0 / (1.0 + x*x) },
		},
	}

	for _, calc := range calculators {
		for _, tt := range testCases {
			testName := fmt.Sprintf("%s/%s", calc.name, tt.name)

			t.Run(testName, func(t *testing.T) {
				result := gausschebyshev.Integrate(calc.calculator, tt.expr)

				assert.InDelta(t, tt.expectedArea, result, tt.tolerance)
			})
		}
	}
}
