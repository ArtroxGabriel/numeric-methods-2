// Package derivatives provides interfaces for calculating derivatives of functions.
package derivatives

import "context"

type Func func(float64) float64

type DerivativeInterface interface {
	Calculate(context.Context, Func, float64, float64) float64
}
