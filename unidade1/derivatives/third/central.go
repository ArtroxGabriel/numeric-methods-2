// Package third implementa a derivada terceira em diferentes filosofias e ordem de erro.
package third

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Central)(nil)

// Central é uma struct que calcula a terceira derivada pela filosofia progressiva.
type Central struct {
	// formula armazena a função de cálculo específica (com base na ordem de erro).
	formula func(ctx context.Context, f derivatives.Func, x, h float64) (float64, error)
}

func NewCentral(errorOrder uint64) *Central {
	var selectedFormula func(ctx context.Context, f derivatives.Func, x, h float64) (float64, error)

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
		panic(fmt.Sprintf("ordem de erro inválida para derivada progressiva: %d", errorOrder))
	}

	return &Central{
		formula: selectedFormula,
	}
}

// Calculate executa o cálculo da derivada usando a fórmula que foi definida no Newcentral.
func (b *Central) Calculate(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	// A mágica acontece aqui: chamamos a fórmula que foi "injetada".
	return b.formula(ctx, f, x, h)
}

func centralOrder1(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	result := (-f(x-1.5*h) + 3*f(x-0.5*h) - 3*f(x+0.5*h) + f(x+1.5*h)) / h3

	return result, nil
}

func centralOrder2(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	numerador := (-2*f(x-2*h) + 7*f(x-h) - 9*f(x) + 5*f(x+h) - f(x+2*h))

	return numerador / h3, nil
}

func centralOrder3(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	// -11 A + 35 B - 38 C + 14 D + E - F
	numerador := (-11*f(x-2.5*h) +
		35*f(x-1.5*h) +
		-38*f(x-0.5*h) +
		14*f(x+0.5*h) +
		f(x+1.5*h) +
		-f(x+2.5*h))

	return numerador / (8 * h3), nil
}

func centralOrder4(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))
	h3 := h * h * h

	// (-17 A + 70 B - 119 C + 108 D - 55 E + 14 F - G)
	// mid is D, the dx to other is mul of h
	numerador := (-17*f(x-3*h) +
		70*f(x-2*h) +
		-119*f(x-h) +
		108*f(x) +
		-55*f(x+h) +
		14*f(x+2*h) +
		-f(x+3*h))

	return numerador / (8 * h3), nil
}
