# Métodos Numéricos 2 - Unidade 1

Este projeto implementa algoritmos de cálculo de derivadas numéricas e processamento de imagens usando Go. O código está organizado em dois módulos principais: **derivadas numéricas** e **processamento de imagens**.

## 📊 Derivadas Numéricas

### Estrutura do Módulo `derivatives`

O módulo de derivadas implementa três tipos de aproximações numéricas para derivadas de primeira, segunda e terceira ordem:

#### Estrutura de Pastas

```
derivatives/
├── derivatives.go          # Interface principal DerivativeInterface
├── derivatives_*_test.go   # Testes para cada ordem de derivada
├── first/                  # Primeira derivada
│   ├── forward.go         # Aproximação progressiva (O(h¹) até O(h⁴))
│   ├── backward.go        # Aproximação regressiva (O(h¹) até O(h⁴))
│   └── central.go         # Aproximação central (O(h¹) até O(h⁴))
├── second/                 # Segunda derivada
│   ├── forward.go         # Aproximação progressiva (O(h¹) até O(h⁴))
│   ├── backward.go        # Aproximação regressiva (O(h¹) até O(h⁴))
│   └── central.go         # Aproximação central (O(h¹) até O(h⁴))
└── third/                  # Terceira derivada
    ├── forward.go         # Aproximação progressiva (O(h¹) até O(h⁴))
    ├── backward.go        # Aproximação regressiva (O(h¹) até O(h⁴))
    └── central.go         # Aproximação central (O(h¹) até O(h⁴))
```

#### Como Funciona

**Interface Principal:**
```go
type DerivativeInterface interface {
    Calculate(context.Context, Func, float64, float64) (float64, error)
}
```

**Método de Desenvolvimento:**
- Todas as fórmulas foram desenvolvidas utilizando o **método de interpolação polinomial de Newton**
- Cada implementação oferece 4 ordens de precisão: O(h¹), O(h²), O(h³) e O(h⁴)
- Maior ordem = maior precisão numérica

**Exemplo de Uso:**
```go
// Criar uma instância de derivada central de primeira ordem com precisão O(h²)
derivativa := first.NewCentral(2)

// Calcular f'(2) onde f(x) = x⁴
resultado, err := derivativa.Calculate(context.Background(), func(x float64) float64 {
    return x * x * x * x  // f(x) = x⁴
}, 2.0, 1e-3)  // x=2, h=0.001
```

#### Testes

- `derivatives_first_test.go` - Testa todas as aproximações de primeira derivada
- `derivatives_second_test.go` - Testa todas as aproximações de segunda derivada  
- `derivatives_third_test.go` - Testa todas as aproximações de terceira derivada

Função teste: f(x) = x⁴ com derivadas analíticas conhecidas para validação.

## 🖼️ Processamento de Imagens

### Estrutura do Módulo `imaging`

O módulo de processamento implementa algoritmos de detecção de bordas e filtros de convolução:

```
imaging/
├── convolution.go      # Operações de convolução básicas
├── custom.go          # Detectores customizados usando derivadas numéricas
├── filters.go         # Kernels de filtros (Sobel, Gaussian, etc.)
├── laplacian.go       # Detector de bordas Laplaciano
├── sobel.go           # Detector de bordas Sobel
└── utils.go           # Utilitários (carregar/salvar imagens)
```

### Algoritmos Implementados

#### 1. **Algoritmo Sobel** (`sobel.go`)
- **Função:** `DetectEdgesSobel(inputImg *image.Gray, threshold float64)`
- **Processo:**
  1. Suavização com filtro Gaussiano 5x5
  2. Aplicação do kernel Sobel X (gradiente horizontal)
  3. Aplicação do kernel Sobel Y (gradiente vertical) 
  4. Cálculo da magnitude: `sqrt(SobelX² + SobelY²)`
  5. Aplicação de threshold para binarização

#### 2. **Algoritmo Laplaciano** (`laplacian.go`)
- **Função:** `DetectEdgesLaplacian(inputImg *image.Gray, tolerance float64)`
- **Processo:**
  1. Suavização com filtro Gaussiano
  2. Aplicação do kernel Laplaciano (segunda derivada)
  3. Detecção de cruzamentos por zero com tolerância

#### 3. **Detector Central O(h⁴)** (`custom.go`)
- **Função:** `DetectEdgesCentralO4(inputImg *image.Gray, threshold float64)`
- Utiliza aproximação central de **quarta ordem** das derivadas numéricas
- Calcula gradientes com maior precisão numérica

#### 4. **Detector Backward O(h⁴)** (`custom.go`)
- **Função:** `DetectEdgesBackward04(inputImg *image.Gray, threshold float64)`
- Utiliza aproximação regressiva de **quarta ordem** das derivadas numéricas
- Ideal para bordas próximas ao final da imagem

### Convolução

**Função Principal:** `Convolve(img *image.Gray, kernel [][]float64)`

Implementa a operação matemática de convolução 2D:
```
resultado[x,y] = Σ Σ imagem[x+i,y+j] × kernel[i,j]
```

### Kernels Disponíveis

- **Sobel X/Y:** Detecta gradientes horizontais/verticais
- **Laplaciano:** Segunda derivada para detecção de bordas
- **Gaussiano 5x5:** Suavização e redução de ruído
- **Derivadas numéricas customizadas:** Forward, Backward, Central em diferentes ordens

## 🚀 Executando o Projeto

### Processamento de Imagens

```bash
go run main.go
```

Este comando:
1. Carrega a imagem `data/pngwing.com.png`
2. Aplica os 4 algoritmos de detecção de bordas
3. Salva os resultados em `data/resultado_*.png`

### Executando Testes

```bash
# Todos os testes
go test ./...

# Apenas testes de derivadas
go test ./derivatives/...

# Teste específico
go test ./derivatives/ -run TestDerivatives_first_order1
```

## 📁 Arquivos de Saída

- `data/resultado_sobel.png` - Bordas detectadas com Sobel
- `data/resultado_laplace.png` - Bordas detectadas com Laplaciano  
- `data/resultado_central.png` - Bordas detectadas com Central O(h⁴)
- `data/resultado_backward.png` - Bordas detectadas com Backward O(h⁴)
- `data/resultado_forward.png` - Bordas detectadas com Forward O(h⁴)

## 🔧 Dependências

- **Go 1.21+**