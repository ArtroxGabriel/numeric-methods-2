package derivatives_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives/second"
	"github.com/stretchr/testify/assert"
)

// f”(x) = 12x^2
func cubicFuncD2(xi float64) float64 {
	return 12.0 * xi * xi
}

func TestDerivatives_second_order1(t *testing.T) {
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
			name:             "SecondForwardO1",
			derivativeMethod: second.NewForward(order),
			expected:         cubicFuncD2(x),
		},
		{
			name:             "SecondBackwardO1",
			derivativeMethod: second.NewBackward(order),
			expected:         cubicFuncD2(x),
		},
		{
			name:             "SecondCentralO1",
			derivativeMethod: second.NewCentral(order),
			expected:         cubicFuncD2(x),
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

func TestDerivatives_second_order2(t *testing.T) {
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
			name:             "SecondForwardO2",
			derivativeMethod: second.NewForward(order),
			expected:         cubicFuncD2(x),
		},
		{
			name:             "SecondBackwardO2",
			derivativeMethod: second.NewBackward(order),
			expected:         cubicFuncD2(x),
		},
		{
			name:             "SecondCentralO2",
			derivativeMethod: second.NewCentral(order),
			expected:         cubicFuncD2(x),
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

func TestDerivatives_second_order3(t *testing.T) {
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
	tolerance := 1e-5

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "SecondForwardO3",
			derivativeMethod: second.NewForward(order),
			expected:         cubicFuncD2(x),
		},
		{
			name:             "SecondBackwardO3",
			derivativeMethod: second.NewBackward(order),
			expected:         cubicFuncD2(x),
		},
		{
			name:             "SecondCentralO3",
			derivativeMethod: second.NewCentral(order),
			expected:         cubicFuncD2(x),
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

func TestDerivatives_second_order4(t *testing.T) {
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
	tolerance := 1e-4

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "SecondForwardO4",
			derivativeMethod: second.NewForward(order),
			expected:         cubicFuncD2(x),
		},
		{
			name:             "SecondBackwardO4",
			derivativeMethod: second.NewBackward(order),
			expected:         cubicFuncD2(x),
		},
		{
			name:             "SecondCentralO4",
			derivativeMethod: second.NewCentral(order),
			expected:         cubicFuncD2(x),
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
