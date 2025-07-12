package third

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Backward)(nil)

// Backward é uma struct que calcula a terceira derivada pela filosofia progressiva.
type Backward struct {
	// formula armazena a função de cálculo específica (com base na ordem de erro).
	formula func(ctx context.Context, f derivatives.Func, x, h float64) (float64, error)
}

func NewBackward(errorOrder uint64) *Backward {
	var selectedFormula func(ctx context.Context, f derivatives.Func, x, h float64) (float64, error)

	switch errorOrder {
	case 1:
		selectedFormula = backwardOrder1
	case 2:
		selectedFormula = backwardOrder2
	case 3:
		selectedFormula = backwardOrder3
	case 4:
		selectedFormula = backwardOrder4
	default:
		panic(fmt.Sprintf("ordem de erro inválida para derivada progressiva: %d", errorOrder))
	}

	return &Backward{
		formula: selectedFormula,
	}
}

// Calculate executa o cálculo da derivada usando a fórmula que foi definida no NewBackward.
func (b *Backward) Calculate(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	// A mágica acontece aqui: chamamos a fórmula que foi "injetada".
	return b.formula(ctx, f, x, h)
}

// backwardOrder1 implementa a fórmula regressiva com erro O(h).
// backward euler method
func backwardOrder1(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	result := (-f(x-3*h) + 3*f(x-2*h) - 3*f(x-h) + f(x)) / h3

	return result, nil
}

// backwardOrder2 implementa a fórmula regressiva com erro O(h²).
func backwardOrder2(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	numerador := (-f(x-4*h) + 3*f(x-3*h) - 3*f(x-2*h) + f(x-h))

	return numerador / h3, nil
}

func backwardOrder3(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	//   (-13 A + 65 B - 134 C + 142 D - 77 E + 17 F)
	numerador := (-13*f(x-5*h) +
		65*f(x-4*h) +
		-134*f(x-3*h) +
		142*f(x-2*h) +
		-77*f(x-h) +
		+17*f(x))

	return numerador / (4 * h3), nil
}

func backwardOrder4(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	//  3 A - 56 B + 235 C - 448 D + 449 E - 232 F + 49 G)
	numerador := (3*f(x-6*h) +
		-56*f(x-5*h) +
		235*f(x-4*h) +
		-448*f(x-3*h) +
		449*f(x-2*h) +
		-232*f(x-h) +
		49*f(x))

	return numerador / (8 * h3), nil
}
