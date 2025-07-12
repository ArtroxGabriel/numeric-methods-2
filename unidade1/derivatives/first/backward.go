// Package first contém implementações de derivadas de primeira ordem.
package first

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Backward)(nil)

// Backward é uma struct que calcula a primeira derivada pela filosofia regressiva.
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
		// Panicking aqui é uma boa escolha, pois é um erro de programação
		// usar uma ordem que não existe.
		panic(fmt.Sprintf("ordem de erro inválida para derivada regressiva: %d", errorOrder))
	}

	return &Backward{
		formula: selectedFormula,
	}
}

// Calculate executa o cálculo da derivada usando a fórmula que foi definida no NewBackward.
func (b *Backward) Calculate(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	// A mágica acontece aqui: chamamos a fórmula que foi "injetada".
	return b.formula(ctx, f, x, h)
}

// backwardOrder1 implementa a fórmula regressiva com erro O(h).
// backward euler method
func backwardOrder1(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada regressiva de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (f(x) - f(x-h)) / h
}

// backwardOrder2 implementa a fórmula regressiva com erro O(h²).
func backwardOrder2(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada regressiva de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (3*f(x) - 4*f(x-h) + f(x-2*h)) / (2 * h)
}

func backwardOrder3(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada regressiva de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (11*f(x) - 18*f(x-h) + 9*f(x-2*h) - 2*f(x-3*h)) / (6 * h)
}

func backwardOrder4(ctx context.Context, f derivatives.Func, x, h float64) float64 {
	slog.DebugContext(ctx, "Calculando a derivada regressiva de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return (25*f(x) - 48*f(x-h) + 36*f(x-2*h) - 16*f(x-3*h) + 3*f(x-4*h)) / (12 * h)
}
