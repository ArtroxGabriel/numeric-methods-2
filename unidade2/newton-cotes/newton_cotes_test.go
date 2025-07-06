package newtoncotes_test

import (
	"fmt"
	"math"
	"testing"

	newtoncotes "github.com/ArtroxGabriel/numeric-methods-2/unidade2/newton-cotes"
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
	const e = 1e-6
	const tolerance = 1e-6

	calculators := []struct {
		name       string
		calculator newtoncotes.NewtonCotesCalculator
	}{
		{"ClosedOrder2", newtoncotes.NewClosedOrder2()},
		{"ClosedOrder3", newtoncotes.NewClosedOrder3()},
		{"ClosedOrder4", newtoncotes.NewClosedOrder4()},
		{"OpenOrder2", newtoncotes.NewOpenOrder2()},
		{"OpenOrder3", newtoncotes.NewOpenOrder3()},
		{"OpenOrder4", newtoncotes.NewOpenOrder4()},
	}

	for _, calc := range calculators {
		for _, tc := range testCases {
			testName := fmt.Sprintf("%s/%s", calc.name, tc.name)

			t.Run(testName, func(t *testing.T) {
				result := newtoncotes.Integrate(calc.calculator, tc.f, tc.a, tc.b, e)

				assert.InDelta(t, tc.expected, result.Result, tolerance)

				t.Logf("Número de iterações: %d", result.NumOfIterations)
			})
		}
	}
}
