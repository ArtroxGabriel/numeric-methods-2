// Package second contém implementações de derivadas de segunda ordem.
package second

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Forward)(nil)

// Forward é uma struct que calcula a segunda derivada pela filosofia progressiva.
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

// Calculate executa o cálculo da derivada usando a fórmula que foi definida no NewForward.
func (b *Forward) Calculate(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	return b.formula(ctx, f, x, h), nil
}

func forwardOrder1(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x) - 2*f(x+h) + f(x+2*h)) / (h * h)
}

func forwardOrder2(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (2*f(x) - 5*f(x+h) + 4*f(x+2*h) - f(x+3*h)) / (h * h)
}

func forwardOrder3(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (35*f(x) - 104*f(x+h) + 114*f(x+2*h) - 56*f(x+3*h) + 11*f(x+4*h)) / (12 * h * h)
}

func forwardOrder4(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	numerador := 45*f(x) - 154*f(x+h) + 214*f(x+2*h) - 156*f(x+3*h) + 61*f(x+4*h) - 10*f(x+5*h)
	return numerador / (12 * h * h)
}
