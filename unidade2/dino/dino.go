// Package dino implements the Dino method for numerical integration using Gauss-Hermite quadrature.
package dino

import (
	"math"

	gausshermite "github.com/ArtroxGabriel/numeric-methods-2/unidade2/gauss-hermite"
	"github.com/ArtroxGabriel/numeric-methods-2/unidade2/result"
)

type interval struct {
	a, b float64
}

type DinoCalculator interface {
	Calculate(func(float64) float64, float64, float64) float64
}

// IntegrateDino calcula a integral de uma função f no intervalo [a, b]
// usando o DinoCalculator fornecido e refinando a área até atingir a tolerância.
func IntegrateDino(calculator DinoCalculator, f func(float64) float64, a, b float64) *result.IntegrateResult {
	iterations := 0
	tolerance := 1e-5

	// A função recursiva interna que fará o trabalho de integração adaptativa.
	// Retorna a área calculada para o subintervalo e o número de iterações.
	var integrateRecursive func(float64, float64) (float64, int)

	integrateRecursive = func(currentA, currentB float64) (float64, int) {
		iterations++ // Incrementa o contador de iterações global

		// Calcula a área inicial para o intervalo atual
		areaFull := calculator.Calculate(f, currentA, currentB)

		// Divide o intervalo ao meio
		mid := (currentA + currentB) / 2.0

		// Calcula as áreas para as duas metades
		areaLeft := calculator.Calculate(f, currentA, mid)
		areaRight := calculator.Calculate(f, mid, currentB)

		// Soma das áreas das duas metades
		areaHalves := areaLeft + areaRight

		// Calcula a diferença entre a área total e a soma das metades.
		// Esta é uma heurística comum para estimar o erro em métodos adaptativos.
		// Se a área mudou pouco ao ser dividida, é um bom sinal de convergência.
		errorEstimate := math.Abs(areaFull - areaHalves)

		// Se a estimativa de erro for menor que a tolerância (dividida por 2,
		// pois cada subintervalo contribui para a tolerância total),
		// ou se o intervalo for muito pequeno, retornamos a área das metades
		// como a melhor aproximação para este subintervalo.
		// O `currentB - currentA < 1e-9` é para evitar divisões infinitas em
		// casos onde a tolerância nunca é atingida (e também se o intervalo é 0).
		if errorEstimate < tolerance || (currentB-currentA) < 1e-9 {
			return areaHalves, 1 // Retorna a área e 1 iteração (para este subproblema)
		}

		// Se não convergiu, chama recursivamente para as duas metades
		areaLeftRecursive, iterLeft := integrateRecursive(currentA, mid)
		areaRightRecursive, iterRight := integrateRecursive(mid, currentB)

		return areaLeftRecursive + areaRightRecursive, iterLeft + iterRight
	}

	// Chama a função recursiva para o intervalo completo
	finalArea, _ := integrateRecursive(a, b) // Não precisamos das iterações individuais aqui, o `iterations` global já conta.

	return result.NewIntegrateResult(finalArea, iterations)
}

var (
	_ DinoCalculator = (*DinoSimples)(nil)
	_ DinoCalculator = (*DinoDuo)(nil)
)

// DinoSimples is exponencial simples
type DinoSimples struct {
	hermite gausshermite.GaussHermiteCalculator
}

func NewDinoSimples() *DinoSimples {
	return &DinoSimples{
		hermite: gausshermite.NewFourPoints(),
	}
}

func (d *DinoSimples) Calculate(f func(float64) float64, a, b float64) float64 {
	xS := func(s float64) float64 {
		return ((a + b) + (b-a)*math.Tanh(s)) / 2.0
	}
	dxS := func(s float64) float64 {
		return (b - a) / (2.0 * math.Pow(math.Cosh(s), 2))
	}

	fHat := func(s float64) float64 {
		return f(xS(s)) * dxS(s)
	}

	fHat2 := func(s float64) float64 {
		return math.Exp(s*s) * fHat(s)
	}

	return d.hermite.Calculate(fHat2)
}

// DinoDuo is exponencial dupla
type DinoDuo struct {
	hermite gausshermite.GaussHermiteCalculator
}

func NewDinoDuo() *DinoDuo {
	return &DinoDuo{
		hermite: gausshermite.NewFourPoints(),
	}
}

func (d *DinoDuo) Calculate(f func(float64) float64, a, b float64) float64 {
	piOverTwo := math.Pi / 2.0
	xS := func(s float64) float64 {
		return ((a + b) + (b-a)*math.Tanh(math.Sinh(s)*piOverTwo)) / 2.0
	}
	dxS := func(s float64) float64 {
		return piOverTwo * ((b - a) * math.Cosh(s)) / (2.0 * math.Pow(math.Cosh(piOverTwo*math.Sinh(s)), 2))
	}

	fHat := func(s float64) float64 {
		return f(xS(s)) * dxS(s)
	}

	fhat2 := func(s float64) float64 {
		return math.Exp(s*s) * fHat(s)
	}

	return d.hermite.Calculate(fhat2)
}
