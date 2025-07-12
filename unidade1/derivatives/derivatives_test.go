package derivatives_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives/first"
	"github.com/stretchr/testify/assert"
)

// f(x) = x³
func cubicFunc(xi float64) float64 {
	return xi * xi * xi
}

// f'(x) = 3x²
func cubicFuncD1(xi float64) float64 {
	return 3 * xi * xi
}

// f”(x) = 6x
func cubicFuncD2(xi float64) float64 {
	return 6 * xi
}

// f”'(x) = 6x
func cubicFuncD3(_ float64) float64 {
	return 6.0
}

func TestDerivatives_order1(t *testing.T) {
	// arrange log
	t.Parallel()
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	x := float64(2)
	h := 1e-3
	order := uint64(1)

	// tolerancia alta pq sao ruins
	tolerance := 1e-2

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "FirstForwardO1",
			derivativeMethod: first.NewForward(order),
			expected:         cubicFuncD1(x),
		},
		{
			name:             "FirstBackwardO1",
			derivativeMethod: first.NewBackward(order),
			expected:         cubicFuncD1(x),
		},
		{
			name:             "FirstCentralO2",
			derivativeMethod: first.NewCentral(order),
			expected:         cubicFuncD1(x),
		},
	}

	// Loop através da tabela de casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Executa a função de derivada que está sendo testada
			ctx := context.Background()
			method := tc.derivativeMethod
			got := method.Calculate(ctx, cubicFunc, x, h)

			assert.InDeltaf(t,
				tc.expected,
				got,
				tolerance,
				"Função %s falhou: esperado %.8f, obtido %.8f",
				tc.name, tc.expected, got,
			)
		})
	}
}

func TestDerivatives_order2(t *testing.T) {
	// arrange log
	t.Parallel()
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	x := float64(2)
	h := 1e-3
	order := uint64(2)

	// tolerancia alta pq sao ruins
	tolerance := 1e-3

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "FirstForwardO2",
			derivativeMethod: first.NewForward(order),
			expected:         cubicFuncD1(x),
		},
		{
			name:             "FirstBackwardO2",
			derivativeMethod: first.NewBackward(order),
			expected:         cubicFuncD1(x),
		},
		{
			name:             "FirstCentralO2",
			derivativeMethod: first.NewCentral(order),
			expected:         cubicFuncD1(x),
		},
	}

	// Loop através da tabela de casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Executa a função de derivada que está sendo testada
			ctx := context.Background()
			method := tc.derivativeMethod
			got := method.Calculate(ctx, cubicFunc, x, h)

			assert.InDeltaf(t,
				tc.expected,
				got,
				tolerance,
				"Função %s falhou: esperado %.8f, obtido %.8f",
				tc.name, tc.expected, got,
			)
		})
	}
}

func TestDerivatives_order3(t *testing.T) {
	// arrange log
	t.Parallel()
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	x := float64(2)
	h := 1e-3
	order := uint64(3)

	// tolerancia alta pq sao ruins
	tolerance := 1e-4

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "FirstForwardO3",
			derivativeMethod: first.NewForward(order),
			expected:         cubicFuncD1(x),
		},
		{
			name:             "FirstBackwardO3",
			derivativeMethod: first.NewBackward(order),
			expected:         cubicFuncD1(x),
		},
		{
			name:             "FirstCentralO3",
			derivativeMethod: first.NewCentral(order),
			expected:         cubicFuncD1(x),
		},
	}

	// Loop através da tabela de casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Executa a função de derivada que está sendo testada
			ctx := context.Background()
			method := tc.derivativeMethod
			got := method.Calculate(ctx, cubicFunc, x, h)

			assert.InDeltaf(t,
				tc.expected,
				got,
				tolerance,
				"Função %s falhou: esperado %.8f, obtido %.8f",
				tc.name, tc.expected, got,
			)
		})
	}
}

func TestDerivatives_order4(t *testing.T) {
	// arrange log
	t.Parallel()
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	x := float64(2)
	h := 1e-3
	order := uint64(4)

	// tolerancia alta pq sao ruins
	tolerance := 1e-5

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "FirstForwardO4",
			derivativeMethod: first.NewForward(order),
			expected:         cubicFuncD1(x),
		},
		{
			name:             "FirstBackwardO4",
			derivativeMethod: first.NewBackward(order),
			expected:         cubicFuncD1(x),
		},
		{
			name:             "FirstCentralO3",
			derivativeMethod: first.NewCentral(order),
			expected:         cubicFuncD1(x),
		},
	}

	// Loop através da tabela de casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Executa a função de derivada que está sendo testada
			ctx := context.Background()
			method := tc.derivativeMethod
			got := method.Calculate(ctx, cubicFunc, x, h)

			assert.InDeltaf(t,
				tc.expected,
				got,
				tolerance,
				"Função %s falhou: esperado %.8f, obtido %.8f",
				tc.name, tc.expected, got,
			)
		})
	}
}
