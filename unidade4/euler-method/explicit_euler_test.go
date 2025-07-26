package eulermethod_test

import (
	"context"
	"log/slog"
	"math"
	"os"
	"testing"

	eulermethod "github.com/ArtroxGabriel/numeric-methods-2/unidade4/euler-method"
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

func TestExplicitEuler_Execute(t *testing.T) {
	delta := 1e-1

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fc               types.DerivativeFunc
		initialCondition *mat.VecDense
		initialTime      float64
		h                float64
		want             *eulermethod.EulerResult
	}{
		{
			name: "PVI - 1",
			fc: func(ctx context.Context, v *mat.Dense, i int, _ float64) *mat.VecDense {
				slog.InfoContext(ctx, "Executing PVI - 1")
				y := v.RowView(i)

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
			want: &eulermethod.EulerResult{
				Time:  0.5,
				State: mat.NewVecDense(1, []float64{2.79122}),
			},
		},
		{
			name: "PVI - 2",
			fc: func(ctx context.Context, v *mat.Dense, i int, _ float64) *mat.VecDense {
				slog.InfoContext(ctx, "Executing PVI - 2")
				y := v.RowView(i)

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
			want: &eulermethod.EulerResult{
				Time:  0.1,
				State: mat.NewVecDense(2, []float64{1.763, 150.3}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			em := eulermethod.NewExplicitEuler()
			got := em.Execute(context.Background(), tt.fc, tt.initialCondition, tt.initialTime, tt.h)

			assert.Equal(t, tt.want.Time, got.Time, "Time should match")

			expectedVector := tt.want.State.RawVector().Data
			actualVector := got.State.RawVector().Data

			// iterate over the slices and compare each element, compare with relative error:TestExplicitEuler_Execute
			compareSlices(t, expectedVector, actualVector, delta)
		})
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
