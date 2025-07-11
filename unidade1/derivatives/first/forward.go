package first

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Forward)(nil)

// Forward é uma struct que calcula a primeira derivada pela filosofia progressiva.
type Forward struct {
	// formula armazena a função de cálculo específica (com base na ordem de erro).
	formula func(ctx context.Context, f derivatives.Func, x, h float64) float64
}

func NewForward(errorOrder uint64) *Forward {
	var selectedFormula func(ctx context.Context, f derivatives.Func, x, h float64) float64

	switch errorOrder {
	case 1:
		selectedFormula = forwardOrder1
	case 2:
		selectedFormula = forwardOrder2
	case 3:
		selectedFormula = forwardOrder3
	case 4:
		selectedFormula = forwardOrder4
	default:
		panic(fmt.Sprintf("ordem de erro inválida para derivada progressiva: %d", errorOrder))
	}

	return &Forward{
		formula: selectedFormula,
	}
}

// Calculate executa o cálculo da derivada usando a fórmula que foi definida no NewBackward.
func (b *Forward) Calculate(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	// A mágica acontece aqui: chamamos a fórmula que foi "injetada".
	return b.formula(ctx, f, x, h), nil
}

// backwardOrder1 implementa a fórmula regressiva com erro O(h).
// backward euler method
func forwardOrder1(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x+h) - f(x)) / h
}

// forwardOrder2 implementa a fórmula regressiva com erro O(h²).
func forwardOrder2(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (-f(x+2*h) + 4*f(x+h) - 3*f(x)) / (2 * h)
}

func forwardOrder3(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	// 1/6 (-11 A + 18 B - 9 C + 2 D)
	return (-11.0*f(x) + 18.0*f(x+h) - 9.0*f(x+2*h) + 2.0*f(x+3*h)) / (6.0 * h)
}

func forwardOrder4(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (-3*f(x+4*h) + 16*f(x+3*h) - 36*f(x+2*h) + 48*f(x+h) - 25*f(x)) / (12 * h)
}
