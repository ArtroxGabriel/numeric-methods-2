package first

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Central)(nil)

// Central é uma struct que calcula a primeira derivada pela filosofia progressiva.
type Central struct {
	// formula armazena a função de cálculo específica (com base na ordem de erro).
	formula func(ctx context.Context, f derivatives.Func, x, h float64) float64
}

func NewCentral(errorOrder uint64) *Central {
	var selectedFormula func(ctx context.Context, f derivatives.Func, x, h float64) float64

	switch errorOrder {
	case 1:
		selectedFormula = centralOrder1
	case 2:
		selectedFormula = centralOrder2
	case 3:
		selectedFormula = centralOrder3
	case 4:
		selectedFormula = centralOrder4
	default:
		panic(fmt.Sprintf("ordem de erro inválida para derivada central: %d", errorOrder))
	}

	return &Central{
		formula: selectedFormula,
	}
}

// Calculate executa o cálculo da derivada usando a fórmula que foi definida no Newcentral.
func (b *Central) Calculate(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	// A mágica acontece aqui: chamamos a fórmula que foi "injetada".
	return b.formula(ctx, f, x, h), nil
}

// centralOrder1 implementa a fórmula regressiva com erro O(h).
// central euler method
func centralOrder1(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada central de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x+0.5*h) - f(x-0.5*h)) / (h)
}

// centralOrder2 implementa a fórmula regressiva com erro O(h²).
func centralOrder2(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada central de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x+h) - f(x-h)) / (2.0 * h)
}

func centralOrder3(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada central de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x-1.5*h) - 27*f(x-0.5*h) + 27*f(x+0.5*h) - f(x+1.5*h)) / (24.0 * h)
}

func centralOrder4(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada central de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))
	return (f(x-2*h) - 8*f(x-h) + 8*f(x+h) - f(x+2*h)) / (12 * h)
}
