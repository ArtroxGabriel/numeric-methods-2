// imaging/custom.go

package imaging

import (
	"image"
	"image/color"
	"math"
)

// DetectEdgesCentralO4 implementa a detecção de bordas com o kernel Central O(h⁴).
func DetectEdgesCentralO4(inputImg *image.Gray, threshold float64) *image.Gray {
	// O divisor da fórmula é 12.
	divisor := 12.0

	// 1) Suavize a imagem
	blurred := Convolve(inputImg, GaussianKernel5x5)

	// 2) Aplique os kernels CentralO4
	imgA := Convolve(blurred, CentralO4X)
	imgB := Convolve(blurred, CentralO4Y)

	bounds := inputImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	finalImg := image.NewGray(bounds)

	// 3) Calcule a magnitude, aplicando o divisor
	for y := range height {
		for x := range width {
			// Aplica o divisor ao resultado da convolução
			valA := float64(imgA.GrayAt(x, y).Y) / divisor
			valB := float64(imgB.GrayAt(x, y).Y) / divisor

			magnitude := math.Sqrt(valA*valA + valB*valB)

			if magnitude > threshold {
				finalImg.SetGray(x, y, color.Gray{Y: 0}) // Borda (preto)
			} else {
				finalImg.SetGray(x, y, color.Gray{Y: 255}) // Fundo (branco)
			}
		}
	}
	return finalImg
}

// DetectEdgesBackward04 implementa a detecção de bordas com o kernel Backward O(h⁴).
func DetectEdgesBackward04(inputImg *image.Gray, threshold float64) *image.Gray {
	// O divisor da fórmula é 12.
	divisor := 12.0

	// 1) Suavize a imagem
	blurred := Convolve(inputImg, GaussianKernel5x5)

	// 2) Aplique os kernels ForwardO4
	imgA := Convolve(blurred, BackwardO4X)
	imgB := Convolve(blurred, BackwardO4Y)

	bounds := inputImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	finalImg := image.NewGray(bounds)

	// 3) Calcule a magnitude, aplicando o divisor
	for y := range height {
		for x := range width {
			// Aplica o divisor ao resultado da convolução
			valA := float64(imgA.GrayAt(x, y).Y) / divisor
			valB := float64(imgB.GrayAt(x, y).Y) / divisor

			magnitude := math.Sqrt(valA*valA + valB*valB)

			if magnitude > threshold {
				finalImg.SetGray(x, y, color.Gray{Y: 0}) // Borda (preto)
			} else {
				finalImg.SetGray(x, y, color.Gray{Y: 255}) // Fundo (branco)
			}
		}
	}
	return finalImg
}

// DetectEdgesForward04 implementa a detecção de bordas com o kernel Forward O(h⁴).
func DetectEdgesForward04(inputImg *image.Gray, threshold float64) *image.Gray {
	// O divisor da fórmula é 12.
	divisor := 12.0

	// 1) Suavize a imagem
	blurred := Convolve(inputImg, GaussianKernel5x5)

	// 2) Aplique os kernels ForwardO4
	imgA := Convolve(blurred, ForwardO4X)
	imgB := Convolve(blurred, ForwardO4Y)

	bounds := inputImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	finalImg := image.NewGray(bounds)

	// 3) Calcule a magnitude, aplicando o divisor
	for y := range height {
		for x := range width {
			// Aplica o divisor ao resultado da convolução
			valA := float64(imgA.GrayAt(x, y).Y) / divisor
			valB := float64(imgB.GrayAt(x, y).Y) / divisor

			magnitude := math.Sqrt(valA*valA + valB*valB)

			if magnitude > threshold {
				finalImg.SetGray(x, y, color.Gray{Y: 0}) // Borda (preto)
			} else {
				finalImg.SetGray(x, y, color.Gray{Y: 255}) // Fundo (branco)
			}
		}
	}
	return finalImg
}
