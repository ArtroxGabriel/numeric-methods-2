// Package types defines the types and interfaces used in the numerical methods for solving ordinary differential equations (ODEs).
package types

import (
	"context"

	"gonum.org/v1/gonum/mat"
)

type DerivativeFunc func(context.Context, *mat.Dense, int) *mat.VecDense
