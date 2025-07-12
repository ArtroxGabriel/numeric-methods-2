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

// f(x) = x^4
func cubicFunc(xi float64) float64 {
	return xi * xi * xi * xi
}

// f'(x) = 4x^3
func cubicFuncD1(xi float64) float64 {
	return 4.0 * xi * xi * xi
}

func TestDerivatives_first_order1(t *testing.T) {
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
	tolerance := 1e-1

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
			name:             "FirstCentralO1",
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
			got, err := method.Calculate(ctx, cubicFunc, x, h)
			assert.NoError(t, err)

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

func TestDerivatives_first_order2(t *testing.T) {
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

			expected: cubicFuncD1(x),
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
			got, err := method.Calculate(ctx, cubicFunc, x, h)
			assert.NoError(t, err)

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

func TestDerivatives_first_order3(t *testing.T) {
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
			got, err := method.Calculate(ctx, cubicFunc, x, h)

			assert.NoError(t, err, "Erro ao calcular a derivada: %v", err)

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

func TestDerivatives_first_order4(t *testing.T) {
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
			got, err := method.Calculate(ctx, cubicFunc, x, h)

			assert.NoError(t, err, "Erro ao calcular a derivada: %v", err)

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
