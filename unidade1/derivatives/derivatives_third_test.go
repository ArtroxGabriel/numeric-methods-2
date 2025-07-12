package derivatives_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/derivatives/third"
	"github.com/stretchr/testify/assert"
)

var x = float64(2)

// f”'(x) = 24x
func cubicFuncD3(xi float64) float64 {
	return 24.0 * xi
}

func TestDerivatives_third_order1(t *testing.T) {
	// arrange log
	t.Parallel()
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

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
			name:             "thirdForwardO1",
			derivativeMethod: third.NewForward(order),
			expected:         cubicFuncD3(x),
		},
		{
			name:             "thirdBackwardO1",
			derivativeMethod: third.NewBackward(order),
			expected:         cubicFuncD3(x),
		},
		{
			name:             "thirdCentralO1",
			derivativeMethod: third.NewCentral(order),
			expected:         cubicFuncD3(x),
		},
	}

	// Loop através da tabela de casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Executa a função de derivada que está sendo testada
			ctx := context.Background()
			method := tc.derivativeMethod
			got, err := method.Calculate(ctx, cubicFunc, x, h)

			if !assert.NoError(t, err, "Erro ao calcular a derivada: %v", err) {
				t.SkipNow()
			}

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

func TestDerivatives_third_order2(t *testing.T) {
	// arrange log
	t.Parallel()
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	h := 1e-3
	order := uint64(2)

	// ta uma porcaria
	tolerance := 1e-1

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "thirdForwardO2",
			derivativeMethod: third.NewForward(order),
			expected:         cubicFuncD3(x),
		},
		{
			name:             "thirdBackwardO2",
			derivativeMethod: third.NewBackward(order),
			expected:         cubicFuncD3(x),
		},
		{
			name:             "thirdCentralO2",
			derivativeMethod: third.NewCentral(order),
			expected:         cubicFuncD3(x),
		},
	}

	// Loop através da tabela de casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Executa a função de derivada que está sendo testada
			ctx := context.Background()
			method := tc.derivativeMethod
			got, err := method.Calculate(ctx, cubicFunc, x, h)

			if !assert.NoError(t, err, "Erro ao calcular a derivada: %v", err) {
				t.SkipNow()
			}

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

func TestDerivatives_third_order3(t *testing.T) {
	// arrange log
	t.Parallel()
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	h := 1e-3
	order := uint64(3)

	tolerance := 1e-1

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "thirdForwardO3",
			derivativeMethod: third.NewForward(order),
			expected:         cubicFuncD3(x),
		},
		{
			name:             "thirdBackwardO3",
			derivativeMethod: third.NewBackward(order),
			expected:         cubicFuncD3(x),
		},
		{
			name:             "thirdCentralO3",
			derivativeMethod: third.NewCentral(order),
			expected:         cubicFuncD3(x),
		},
	}

	// Loop através da tabela de casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Executa a função de derivada que está sendo testada
			ctx := context.Background()
			method := tc.derivativeMethod
			got, err := method.Calculate(ctx, cubicFunc, x, h)

			if !assert.NoError(t, err, "Erro ao calcular a derivada: %v", err) {
				t.SkipNow()
			}

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

func TestDerivatives_third_order4(t *testing.T) {
	// arrange log
	t.Parallel()
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	h := 1e-3
	order := uint64(4)

	// tolerancia alta pq sao ruins
	tolerance := 1e-1

	tests := []struct {
		name             string
		derivativeMethod derivatives.DerivativeInterface
		expected         float64
	}{
		{
			name:             "thirdForwardO4",
			derivativeMethod: third.NewForward(order),
			expected:         cubicFuncD3(x),
		},
		{
			name:             "thirdBackwardO4",
			derivativeMethod: third.NewBackward(order),
			expected:         cubicFuncD3(x),
		},
		{
			name:             "thirdCentralO4",
			derivativeMethod: third.NewCentral(order),
			expected:         cubicFuncD3(x),
		},
	}

	// Loop através da tabela de casos de teste
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Executa a função de derivada que está sendo testada
			ctx := context.Background()
			method := tc.derivativeMethod
			got, err := method.Calculate(ctx, cubicFunc, x, h)

			if !assert.NoError(t, err, "Erro ao calcular a derivada: %v", err) {
				t.SkipNow()
			}

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
