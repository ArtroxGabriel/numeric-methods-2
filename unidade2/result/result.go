package result

type IntegrateResult struct {
	Result          float64
	NumOfIterations int
}

func NewIntegrateResult(result float64, iterations int) *IntegrateResult {
	return &IntegrateResult{
		Result:          result,
		NumOfIterations: iterations,
	}
}
