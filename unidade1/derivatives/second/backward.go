package second

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Backward)(nil)

// Backward é uma struct que calcula a segunda derivada pela filosofia regressiva.
type Backward struct {
	// formula armazena a função de cálculo específica (com base na ordem de erro).
	formula func(ctx context.Context, f derivatives.Func, x, h float64) float64
}

func NewBackward(errorOrder uint64) *Backward {
	var selectedFormula func(ctx context.Context, f derivatives.Func, x, h float64) float64

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
		panic(fmt.Sprintf("ordem de erro inválida para derivada regressiva: %d", errorOrder))
	}

	return &Backward{
		formula: selectedFormula,
	}
}

// Calculate executa o cálculo da derivada usando a fórmula que foi definida no NewBackward.
func (b *Backward) Calculate(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	return b.formula(ctx, f, x, h), nil
}

func backwardOrder1(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada regressiva de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x-2*h) - 2*f(x-h) + f(x)) / (h * h)
}

func backwardOrder2(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada regressiva de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (-f(x-3*h) + 4*f(x-2*h) - 5*f(x-h) + 2*f(x)) / (h * h)
}

func backwardOrder3(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada regressiva de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (11*f(x-4*h) - 56*f(x-3*h) + 114*f(x-2*h) - 104*f(x-h) + 35*f(x)) / (12 * h * h)
}

func backwardOrder4(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada regressiva de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	numerador := -10*f(x-5*h) + 61*f(x-4*h) - 156*f(x-3*h) + 214*f(x-2*h) - 154*f(x-h) + 45*f(x)
	return numerador / (12 * h * h)
}
