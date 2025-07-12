package third

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
)

var _ derivatives.DerivativeInterface = (*Forward)(nil)

// Forward é uma struct que calcula a terceira derivada pela filosofia progressiva.
type Forward struct {
	// formula armazena a função de cálculo específica (com base na ordem de erro).
	formula func(ctx context.Context, f derivatives.Func, x, h float64) (float64, error)
}

func NewForward(errorOrder uint64) *Forward {
	var selectedFormula func(ctx context.Context, f derivatives.Func, x, h float64) (float64, error)

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

// Calculate executa o cálculo da derivada usando a fórmula que foi definida no Newforward.
func (b *Forward) Calculate(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	// A mágica acontece aqui: chamamos a fórmula que foi "injetada".
	return b.formula(ctx, f, x, h)
}

// forwardOrder1 implementa a fórmula regressiva com erro O(h).
// forward euler method
func forwardOrder1(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return 0.0, errors.New("forwardOrder1: not implemented yet")
}

// forwardOrder2 implementa a fórmula regressiva com erro O(h²).
func forwardOrder2(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return 0.0, errors.New("forwardOrder2: not implemented yet")
}

func forwardOrder3(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return 0.0, errors.New("forwardOrder3: not implemented yet")
}

func forwardOrder4(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	return 0.0, errors.New("forwardOrder3: not implemented yet")
}
