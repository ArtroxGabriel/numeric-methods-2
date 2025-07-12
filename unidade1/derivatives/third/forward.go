package third

import (
	"context"
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

func forwardOrder1(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 1",
		slog.Uint64("ordem", uint64(1)),
		slog.Float64("x", x),
		slog.Float64("h", h))
	h3 := h * h * h

	result := (-f(x) + 3*f(x+h) - 3*f(x+2*h) + f(x+3*h)) / h3

	return result, nil
}

func forwardOrder2(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 2",
		slog.Uint64("ordem", uint64(2)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	numerador := (-4*f(x) + 15*f(x+h) - 21*f(x+2*h) + 13*f(x+3*h) - 3*f(x+4*h))

	return numerador / h3, nil
}

func forwardOrder3(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 3",
		slog.Uint64("ordem", uint64(3)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	//  (-23 A + 95 B - 154 C + 122 D - 47 E + 7 F)
	numerador := (-23*f(x) +
		95*f(x+h) +
		-154*f(x+2*h) +
		122*f(x+3*h) +
		-47*f(x+4*h) +
		7*f(x+5*h))

	return numerador / (4 * h3), nil
}

func forwardOrder4(ctx context.Context, f derivatives.Func, x, h float64) (float64, error) {
	slog.DebugContext(ctx, "Calculando a derivada progressiva de ordem 4",
		slog.Uint64("ordem", uint64(4)),
		slog.Float64("x", x),
		slog.Float64("h", h))

	h3 := h * h * h

	//   (-61 A + 280 B - 533 C + 544 D - 319 E + 104 F - 15 G)
	numerador := (-61*f(x) +
		280*f(x+h) +
		-533*f(x+2*h) +
		544*f(x+3*h) +
		-319*f(x+4*h) +
		104*f(x+5*h) +
		-15*f(x+6*h))

	return numerador / (8 * h3), nil
}
