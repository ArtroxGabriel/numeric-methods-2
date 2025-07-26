package rungekutta

import (
	"context"
	"log/slog"

	eulermethod "github.com/ArtroxGabriel/numeric-methods-2/unidade4/euler-method"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade4/types"
	"gonum.org/v1/gonum/mat"
)

var _ RungeKuttaInterface = (*RungeKuttaFourthOrder)(nil)

type RungeKuttaFourthOrder struct {
	em eulermethod.ExplicitEuler
}

func NewRungeKuttaFour() *RungeKuttaFourthOrder {
	return &RungeKuttaFourthOrder{
		em: *eulermethod.NewExplicitEuler(),
	}
}

func (rk *RungeKuttaFourthOrder) RungeKutta() {}

func (rk *RungeKuttaFourthOrder) Execute(
	ctx context.Context,
	fc types.DerivativeFunc,
	initialCondition *mat.VecDense,
	initialTime, h float64,
) *RungeKuttaResult {
	slog.InfoContext(ctx, "Executing runge kutta",
		slog.Uint64("order", uint64(4)),
		slog.Any("initialCondition", initialCondition.RawVector().Data),
		slog.Float64("initialTime", initialTime),
		slog.Float64("h", h),
	)
	n := initialCondition.Len()

	// Create state matrix for intermediate calculations
	states := mat.NewDense(4, n, nil)

	// Set initial state: w_0 = y_0
	states.SetRow(0, initialCondition.RawVector().Data)

	// k1 = h * f(t_0, w_0)
	k1 := fc(ctx, states, 0)
	k1.ScaleVec(h, k1)

	// w_1 = y_0 + k1/2
	w1 := mat.NewVecDense(n, nil)
	w1.AddScaledVec(initialCondition, 0.5, k1)
	states.SetRow(1, w1.RawVector().Data)

	// k2 = h * f(t_0 + h/2, w_1)
	k2 := fc(ctx, states, 1)
	k2.ScaleVec(h, k2)

	// w_2 = y_0 + k2/2
	w2 := mat.NewVecDense(n, nil)
	w2.AddScaledVec(initialCondition, 0.5, k2)
	states.SetRow(2, w2.RawVector().Data)

	// k3 = h * f(t_0 + h/2, w_2)
	k3 := fc(ctx, states, 2)
	k3.ScaleVec(h, k3)

	// w_3 = y_0 + k3
	w3 := mat.NewVecDense(n, nil)
	w3.AddVec(initialCondition, k3)
	states.SetRow(3, w3.RawVector().Data)

	// k4 = h * f(t_0 + h, w_3)
	k4 := fc(ctx, states, 3)
	k4.ScaleVec(h, k4)

	// Final calculation: y_1 = y_0 + (h/6)(k1 + 2*k2 + 2*k3 + k4)
	result := mat.NewVecDense(n, nil)
	result.AddVec(result, k1)               // k1
	result.AddScaledVec(result, 2, k2)      // + 2*k2
	result.AddScaledVec(result, 2, k3)      // + 2*k3
	result.AddVec(result, k4)               // + k4
	result.ScaleVec(1.0/6.0, result)        // * (1/6)
	result.AddVec(initialCondition, result) // y_0 + (h/6)(...)

	// x_{i+1}
	nextStateTime := initialTime + h

	slog.InfoContext(ctx, "Computed next state refined",
		slog.Float64("nextTime", nextStateTime),
		slog.Any("nextState", result.RawVector().Data),
	)
	return NewRungeKuttaResult(nextStateTime, result)
}
