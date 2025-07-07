package gausslegendre_test

import (
	"fmt"
	"math"
	"testing"

	gausslegendre "github.com/ArtroxGabriel/numeric-methods-2/unidade2/gauss-legendre"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	f        func(float64) float64
	a, b     float64
	expected float64
}{
	{
		name:     "Integral de ln(x+2)^2 de 2 a 3",
		f:        func(x float64) float64 { return math.Sin(x) },
		a:        0,
		b:        math.Pi / 2.0,
		expected: 1,
	},
	{
		name:     "Integral de (sen(2x) + 4x^2 + 3x)^2 de 0 a 1",
		f:        func(x float64) float64 { return math.Pow(math.Sin(2*x)+4*x*x+3*x, 2) },
		a:        0,
		b:        1,
		expected: 17.8764703,
	},
}

func TestIntegrate(t *testing.T) {
	t.Parallel()

	const tolerance = 1e-6

	calculators := []struct {
		name       string
		calculator gausslegendre.GaussLegendreCalculator
	}{
		{"2 points", gausslegendre.NewTwoPoints()},
		{"3 points", gausslegendre.NewThreePoints()},
		{"4 points", gausslegendre.NewFourPoints()},
	}

	for _, calc := range calculators {
		for _, tc := range testCases {
			testName := fmt.Sprintf("%s/%s", calc.name, tc.name)

			t.Run(testName, func(t *testing.T) {
				result := gausslegendre.Integrate(calc.calculator, tc.f, tc.a, tc.b, tolerance)

				assert.InDelta(t, tc.expected, result.Result, tolerance)

				t.Logf("Número de iterações: %d", result.NumOfIterations)
			})
		}
	}
}
