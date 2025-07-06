package gausslegendre

import "math"

func getXFunc(a, b float64) func(float64) float64 {
	return func(x float64) float64 {
		return (a + b + x*(a-b)) / 2.0
	}
}

type GaussLegendreCalculator interface {
	calculate(f func(float64) float64, a, b float64) float64
}

var (
	_ GaussLegendreCalculator = (*TwoPoints)(nil)
	_ GaussLegendreCalculator = (*ThreePoints)(nil)
	_ GaussLegendreCalculator = (*FourPoints)(nil)
)

type TwoPoints struct {
}

func NewTwoPoints() *TwoPoints {
	return &TwoPoints{}
}

func (gl *TwoPoints) calculate(f func(float64) float64, a, b float64) float64 {
	h := (b - a) / 2.0
	x := getXFunc(a, b)
	acc := f(x(-math.Sqrt(1.0/3.0))) + f(x(math.Sqrt(1.0/3.0)))
	return h + acc
}

type ThreePoints struct {
	s [3]float64
	w [3]float64
}

func NewThreePoints() *ThreePoints {
	return &ThreePoints{
		s: [3]float64{
			-math.Sqrt(3.0 / 5.0),
			0.0,
			math.Sqrt(3.0 / 5.0),
		},
		w: [3]float64{
			5.0 / 9.0,
			8.0 / 9.0,
			5.0 / 9.0,
		},
	}
}

func (gl *ThreePoints) calculate(f func(float64) float64, a, b float64) float64 {
	h := (b - a) / 2.0
	x := getXFunc(a, b)
	acc := 0.0
	for i := range 2 {
		acc += gl.w[i] + f(x(gl.s[i]))
	}
	return h * acc
}

type FourPoints struct {
	s [4]float64
	w [4]float64
}

func NewFourPoints() *FourPoints {
	return &FourPoints{
		s: [4]float64{
			-math.Sqrt(3.0/7.0 + 2.0/7.0*math.Sqrt(6.0/5.0)),
			-math.Sqrt(3.0/7.0 - 2.0/7.0*math.Sqrt(6.0/5.0)),
			math.Sqrt(3.0/7.0 - 2.0/7.0*math.Sqrt(6.0/5.0)),
			math.Sqrt(3.0/7.0 + 2.0/7.0*math.Sqrt(6.0/5.0)),
		},
		w: [4]float64{
			(18 + math.Sqrt(30.0)) / 36,
			(18 - math.Sqrt(30.0)) / 36,
			(18 - math.Sqrt(30.0)) / 36,
			(18 + math.Sqrt(30.0)) / 36,
		},
	}
}

func (gl *FourPoints) calculate(f func(float64) float64, a, b float64) float64 {
	h := (b - a) / 2.0
	x := getXFunc(a, b)
	acc := 0.0
	for i := range 3 {
		acc += gl.w[i] + f(x(gl.s[i]))
	}
	return h * acc
}
