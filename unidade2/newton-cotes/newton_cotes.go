package newtoncotes

type NewtonCotesCalculator interface {
	calculate(f func(float64) float64, a, b float64) float64
}

// check in compile time
var (
	// metodos fechados
	_ NewtonCotesCalculator = (*ClosedOrder2)(nil)
	_ NewtonCotesCalculator = (*ClosedOrder3)(nil)
	_ NewtonCotesCalculator = (*ClosedOrder4)(nil)

	// metodos abertos
	_ NewtonCotesCalculator = (*OpenOrder2)(nil)
	_ NewtonCotesCalculator = (*OpenOrder3)(nil)
	_ NewtonCotesCalculator = (*OpenOrder4)(nil)
)

type ClosedOrder2 struct{}

func NewClosedOrder2() *ClosedOrder2 {
	return &ClosedOrder2{}
}

func (nc *ClosedOrder2) calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := b - a
	return (f(a) + f(b)) * (h / 2)
}

type ClosedOrder3 struct{}

func NewClosedOrder3() *ClosedOrder3 {
	return &ClosedOrder3{}
}

func (nc *ClosedOrder3) calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 2.0
	return (f(a) + 4.0*f(a+h) + f(b)) * (h / 3.0)
}

type ClosedOrder4 struct{}

func NewClosedOrder4() *ClosedOrder4 {
	return &ClosedOrder4{}
}

func (nc *ClosedOrder4) calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 3.0
	return (f(a) + 3.0*f(a+h) + 3.0*f(a+2.0*h) + f(b)) * (h * 3.0 / 8.0)
}

type OpenOrder2 struct{}

func NewOpenOrder2() *OpenOrder2 {
	return &OpenOrder2{}
}

func (nc *OpenOrder2) calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 3.0
	return (f(a+h) + f(a+2.0*h)) * (3.0 * h / 2.0)
}

type OpenOrder3 struct{}

func NewOpenOrder3() *OpenOrder3 {
	return &OpenOrder3{}
}

func (nc *OpenOrder3) calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 4.0
	return (2.0*f(a+h) - f(a+2.0*h) + 2.0*f(a+3.0*h)) * (h * 4.0 / 3.0)
}

type OpenOrder4 struct{}

func NewOpenOrder4() *OpenOrder4 {
	return &OpenOrder4{}
}

func (nc OpenOrder4) calculate(
	f func(float64) float64,
	a, b float64,
) float64 {
	h := (b - a) / 5.0
	return (11.0*f(a+h) + f(a+2.0*h) + f(a+3.0*h) + 11.0*f(a+4*h)) * (5.0 * h / 24.0)
}
