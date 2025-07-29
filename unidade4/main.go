package main

import (
	"context"
	"fmt"
	"log/slog"
	"math"

	predictorcorrector "github.com/ArtroxGabriel/numeric-methods-2/unidade4/predictor-corrector"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

func main() {
	ctx := context.Background()
	ab := predictorcorrector.NewFourthOrder()

	fT := func(t float64) float64 {
		switch {
		case 0 <= t && t <= 0.5:
			return 4.0 * t
		case 0.5 < t && t <= 1.0:
			return -4.0*t + 4
		default:
			return 0.0
		}
	}

	var fc types.DerivativeFunc = func(
		ctx context.Context,
		d *mat.Dense,
		i int,
		t float64,
	) *mat.VecDense {
		slog.InfoContext(ctx, "Executing")
		y := d.RowView(i)

		result := mat.NewVecDense(y.Len(), nil)

		// element at 0:  f(t)/2.0 - 0.1*sqrt(2)*x'(t) + 2.0*x(t)
		a := fT(t)/2.0 - 0.1*math.Sqrt2*y.AtVec(0) + 2.0*y.AtVec(1)
		result.SetVec(0, a)

		// element at 1: v(t)
		result.SetVec(1, y.AtVec(0))

		slog.InfoContext(ctx, "Computed derivative",
			slog.Any("result", result.RawVector().Data),
		)
		return result
	}

	initialVec := mat.NewVecDense(2, []float64{
		0.0, // v_0
		0.0, // x_0
	})
	startTime := 0.0
	stepSize := 1.2
	errorTolerance := 1e-3

	got := ab.Execute(ctx, fc, initialVec, startTime, stepSize, errorTolerance)
	fmt.Printf("Time: %f\n", got.Time)
	fmt.Printf("%.6f", mat.Formatted(got.State))
}
