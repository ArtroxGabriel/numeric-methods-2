package imaging

import (
	"image"
	"image/color"
	"math"
)

// DetectEdgesSobel implementa o Algoritmo 1.
func DetectEdgesSobel(inputImg *image.Gray, threshold float64) *image.Gray {
	// 1) Suavize a imagem
	blurred := Convolve(inputImg, GaussianKernel5x5)

	// 2.1) Aplique Sobel X
	imgA := Convolve(blurred, SobelX)

	// 2.2) Aplique Sobel Y
	imgB := Convolve(blurred, SobelY)

	bounds := inputImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	finalImg := image.NewGray(bounds)

	// 2.3 e 2.4) Calcule a magnitude do gradiente e aplique o threshold
	for y := range height {
		for x := range width {
			// Valor do gradiente em X e Y
			valA := float64(imgA.GrayAt(x, y).Y)
			valB := float64(imgB.GrayAt(x, y).Y)

			// Magnitude: C = sqrt(A² + B²)
			magnitude := math.Sqrt(valA*valA + valB*valB)

			// 4) Gere a matriz final
			if magnitude > threshold {
				finalImg.SetGray(x, y, color.Gray{Y: 0}) // Borda (preto)
			} else {
				finalImg.SetGray(x, y, color.Gray{Y: 255}) // Fundo (branco)
			}
		}
	}
	return finalImg
}
