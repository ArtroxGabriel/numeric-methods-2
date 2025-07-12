package second

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Central)(nil)

// Central é uma struct que calcula a segunda derivada pela filosofia central.
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
	return b.formula(ctx, f, x, h), nil
}

func centralOrder1(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada central de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x-h) - 2*f(x) + f(x+h)) / (h * h)
}

func centralOrder2(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada central de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x-1.5*h) - f(x-0.5*(h)) - f(x+0.5*h) + f(x+1.5*h)) / (2 * h * h)
}

func centralOrder3(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada central de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x-h) - 2*f(x) + f(x+h)) / (h * h)
}

func centralOrder4(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada central de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	// 1/960 (-1239 A + 5695 B - 8950 C + 6030 D - 1795 E + 259 F)
	numerador := -1239.0*f(x-2.5*h) +
		5695.0*f(x-1.5*h) +
		-8950.0*f(x-0.5*h) +
		6030.0*f(x+0.5*h) +
		-1795.0*f(x+1.5*h) +
		259.0*f(x+2.5*h)
	return numerador / (960.0 * h * h)
}
