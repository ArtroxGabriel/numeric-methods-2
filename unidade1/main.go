package main

import (
	"log/slog"
	"os"

	"github.com/ArtroxGabriel/numeric-methods-2/unidade1/imaging"
)

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("Iniciando o processamento de imagens...")

	originalImg := imaging.LoadImageGrayscale("pngwing.com.png")

	// --- Executa o Algoritmo 1: Sobel ---
	slog.Info("Aplicando detecção de bordas com Sobel...")
	sobelThreshold := 100.0 // Valor experimental, ajuste conforme necessário
	sobelEdges := imaging.DetectEdgesSobel(originalImg, sobelThreshold)
	imaging.SaveImage("resultado_sobel.png", sobelEdges)
	slog.Info("Resultado do Sobel salvo em 'resultado_sobel.png'")

	// --- Executa o Algoritmo 2: Laplace ---
	slog.Info("Aplicando detecção de bordas com Laplace...")
	laplaceTolerance := 5.0 // Valor experimental, ajuste conforme necessário
	laplaceEdges := imaging.DetectEdgesLaplacian(originalImg, laplaceTolerance)
	imaging.SaveImage("resultado_laplace.png", laplaceEdges)
	slog.Info("Resultado do Laplace salvo em 'resultado_laplace.png'")

	centralO4Threshold := 15.0
	centralO4Edges := imaging.DetectEdgesCentralO4(originalImg, centralO4Threshold)

	// --- Executa o Algoritmo 3: Central O(h⁴) ---
	imaging.SaveImage("resultado_central.png", centralO4Edges)
	slog.Info("Resultado do Central O(h⁴) salvo em 'resultado_central.png'")

	backwardO4Threshold := 15.0
	backwardO4Edges := imaging.DetectEdgesBackward04(originalImg, backwardO4Threshold)

	// --- Executa o Algoritmo 4: Backward O(h⁴) ---
	imaging.SaveImage("resultado_backward.png", backwardO4Edges)
	slog.Info("Resultado do Backward O(h⁴) salvo em 'resultado_backward.png'")

	// forward
	forwardO4Threshold := 15.0
	forwardO4Edges := imaging.DetectEdgesForward04(originalImg, forwardO4Threshold)

	// --- Executa o Algoritmo 5: Forward O(h⁴) ---
	imaging.SaveImage("resultado_forward.png", forwardO4Edges)
	slog.Info("Resultado do Forward O(h⁴) salvo em 'resultado_forward.png'")

	slog.Info("Processamento concluído.")
}
