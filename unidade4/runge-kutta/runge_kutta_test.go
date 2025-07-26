package rungekutta_test

import (
	"context"
	"log/slog"
	"math"
	"os"
	"testing"

	rungekutta "github.com/ArtroxGabriel/numeric-methods-2/unidade4/runge-kutta"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func TestRungeKutta_Execute(t *testing.T) {
	tests := []struct {
		name             string
		fc               types.DerivativeFunc
		initialCondition *mat.VecDense
		initialTime      float64
		h                float64
		want             *rungekutta.RungeKuttaResult
	}{
		{
			name: "PVI-1",
			fc: func(ctx context.Context, v *mat.Dense, t int) *mat.VecDense {
				slog.InfoContext(ctx, "Executing PVI - 1")
				y := v.RowView(t)

				result := mat.NewVecDense(y.Len(), nil)
				result.ScaleVec(2.0/3.0, y)

				slog.InfoContext(ctx, "Computed derivative",
					slog.Any("result", result.RawVector().Data),
				)
				return result
			},
			initialCondition: mat.NewVecDense(1, []float64{2.0}),
			initialTime:      0.0,
			h:                0.5,
			want: &rungekutta.RungeKuttaResult{
				Time:  0.5,
				State: mat.NewVecDense(1, []float64{2.79122}),
			},
		},
		{
			name: "PVI-2",
			fc: func(ctx context.Context, v *mat.Dense, t int) *mat.VecDense {
				slog.InfoContext(ctx, "Executing PVI - 2")
				y := v.RowView(t)

				result := mat.NewVecDense(y.Len(), nil)

				// element at 0:  -g * k/m*v(t)
				// where g = 10, and k/m = 0.5/0.5
				result.SetVec(0, -10-(1)*y.AtVec(0))

				// element at 1: v(t)
				result.SetVec(1, y.AtVec(0))

				slog.InfoContext(ctx, "Computed derivative",
					slog.Any("result", result.RawVector().Data),
				)
				return result
			},
			initialCondition: mat.NewVecDense(2, []float64{3, 150}),
			initialTime:      0.0,
			h:                0.1,
			want: &rungekutta.RungeKuttaResult{
				Time:  0.1,
				State: mat.NewVecDense(2, []float64{1.765, 150.235}),
			},
		},
	}

	orders := []struct {
		name      string
		rk        rungekutta.RungeKuttaInterface
		tolerance float64
	}{
		{
			name:      "Second Order",
			rk:        rungekutta.NewRungeKuttaSecond(),
			tolerance: 1e-2,
		},
		{
			name:      "Third Order",
			rk:        rungekutta.NewRungeKuttaThird(),
			tolerance: 4e-3,
		},
		{
			name:      "Fourth Order",
			rk:        rungekutta.NewRungeKuttaFour(),
			tolerance: 2e-3,
		},
	}

	for _, tt := range tests {
		for _, order := range orders {
			t.Run(tt.name+"_"+order.name, func(t *testing.T) {
				ctx := context.Background()
				got := order.rk.Execute(ctx, tt.fc, tt.initialCondition, tt.initialTime, tt.h)

				assert.Equal(t, tt.want.Time, got.Time, "Time should match")

				expectedVector := tt.want.State.RawVector().Data
				actualVector := got.State.RawVector().Data

				compareSlices(t, expectedVector, actualVector, order.tolerance)
			})
		}
	}
}

func compareSlices(t *testing.T, expected, actual []float64, delta float64) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Errorf("expected length %d, got %d", len(expected), len(actual))
		return
	}
	for i := range expected {
		relativeError := math.Abs(expected[i]-actual[i]) / expected[i]
		if relativeError > delta {
			t.Errorf("at index %d: expected %f, got %f (relative error: %f)", i, expected[i], actual[i], relativeError)
		}
	}
}
