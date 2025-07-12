package imaging

import (
	"image"
	"image/color"
	"math"
)

// Convolve aplica um kernel a uma imagem em tons de cinza.
func Convolve(img *image.Gray, kernel [][]float64) *image.Gray {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Cria uma imagem de saída com o mesmo tamanho
	out := image.NewGray(bounds)

	kernelSize := len(kernel)
	kernelRadius := kernelSize / 2

	// Itera sobre cada pixel da imagem (ignorando as bordas para simplificar)
	for y := kernelRadius; y < height-kernelRadius; y++ {
		for x := kernelRadius; x < width-kernelRadius; x++ {
			var sum float64 = 0
			// Aplica o kernel
			for ky := range kernelSize {
				for kx := range kernelSize {
					// Pega o valor do pixel da imagem original
					pixelValue := float64(img.GrayAt(x-kernelRadius+kx, y-kernelRadius+ky).Y)
					// Multiplica pelo valor do kernel
					sum += pixelValue * kernel[ky][kx]
				}
			}
			// Normaliza o valor para o range de 0-255 e atribui ao pixel de saída
			out.SetGray(x, y, color.Gray{Y: uint8(math.Max(0, math.Min(255, sum)))})
		}
	}
	return out
}
