// Package imaging fornece funções para carregar e salvar imagens em tons de cinza.
package imaging

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log/slog"
	"os"
	"path/filepath"
)

// LoadImageGrayscale carrega uma imagem do disco e a converte para tons de cinza.
func LoadImageGrayscale(filename string) *image.Gray {
	// ./data + filename
	filename = filepath.Join("data", filename)

	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Falha ao abrir o arquivo: %v", slog.Any("error", err))
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		slog.Error("Falha ao decodificar a imagem", slog.Any("error", err))
		os.Exit(1)
	}

	// Converte para tons de cinza para simplificar os cálculos
	grayImg := image.NewGray(img.Bounds())
	draw.Draw(grayImg, grayImg.Bounds(), img, image.Point{}, draw.Src)
	return grayImg
}

// SaveImage salva uma imagem em tons de cinza no disco.
func SaveImage(path string, img *image.Gray) {
	path = filepath.Join("data", path)
	file, err := os.Create(path)
	if err != nil {
		slog.Error("Falha ao criar o arquivo de saída", slog.Any("error", err))
		os.Exit(1)
	}
	defer file.Close()

	ext := filepath.Ext(path)
	switch ext {
	case ".png":
		png.Encode(file, img)
	case ".jpg", ".jpeg":
		jpeg.Encode(file, img, nil)
	default:
		slog.Error("Formato de arquivo não suportado", slog.String("extensão", ext))
		os.Exit(1)
	}
}
