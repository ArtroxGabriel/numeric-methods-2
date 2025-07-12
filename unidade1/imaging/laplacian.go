package imaging

import (
	"image"
	"image/color"
	"math"
)

// DetectEdgesLaplacian implementa o Algoritmo 2.
func DetectEdgesLaplacian(inputImg *image.Gray, tolerance float64) *image.Gray {
	// 1) Suavize a imagem (isso é chamado de "Laplacian of Gaussian" ou LoG)
	blurred := Convolve(inputImg, GaussianKernel5x5)

	// 2) Aplique o filtro de Laplace
	imgA := Convolve(blurred, Laplacian)

	bounds := inputImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	finalImg := image.NewGray(bounds)

	// 3) Gere a matriz final baseada na tolerância (detecção de "zero-crossing" simplificada)
	for y := range height {
		for x := range width {
			laplaceVal := float64(imgA.GrayAt(x, y).Y)

			// Nota: O valor do pixel é 0-255. A convolução com Laplace pode dar valores negativos
			// mas a nossa função Convolve limita a 0. Um detector de "zero-crossing" real
			// procuraria por mudanças de sinal nos vizinhos. A sua versão é mais simples.
			// Para o nosso caso, vamos considerar o valor absoluto.
			if math.Abs(laplaceVal) > tolerance {
				finalImg.SetGray(x, y, color.Gray{Y: 0}) // Borda (preto)
			} else {
				finalImg.SetGray(x, y, color.Gray{Y: 255}) // Fundo (branco)
			}
		}
	}
	return finalImg
}
