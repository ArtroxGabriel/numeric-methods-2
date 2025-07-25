# Relatorio Unidade 2


### Dependências Go
- **Go 1.18+**: Versão mínima requerida

## Índice

1. [Métodos de Integração Newton-Cotes](#1-métodos-de-integração-newton-cotes)
2. [Métodos de Quadratura Gaussiana](#2-métodos-de-quadratura-gaussiana)
3. [Métodos DINO (Transformações Exponenciais)](#3-métodos-dino-transformações-exponenciais)
4. [Estrutura de Resultado](#4-estrutura-de-resultado)
5. [Executando Testes](#5-executando-testes)

## 1. Métodos de Integração Newton-Cotes

**Localização**: [newton-cotes/](./newton-cotes/)

O pacote Newton-Cotes implementa fórmulas de quadratura abertas e fechadas para integração numérica com substituição polinomial de graus 1 a 4.

### Métodos Disponíveis

#### Métodos Fechados (extremos incluídos)
- **ClosedOrder2**: Regra do trapézio - Usa valores da função nos extremos
- **ClosedOrder3**: Regra de Simpson 1/3 - Usa 3 pontos incluindo extremos
- **ClosedOrder4**: Regra de Simpson 3/8 - Usa 4 pontos incluindo extremos

#### Métodos Abertos (extremos excluídos)
- **OpenOrder2**: Regra do ponto médio - Usa 2 pontos interiores
- **OpenOrder3**: Fórmula aberta com 3 pontos interiores
- **OpenOrder4**: Fórmula aberta com 4 pontos interiores

### Características Principais

- **Integração Adaptativa**: Usa subdivisão recursiva para atingir tolerância desejada
- **Controle de Erro**: Compara integração do intervalo completo com soma dos meio-intervalos
- **Design Baseado em Interface**: Todos os métodos implementam a interface `NewtonCotesCalculator`

### Exemplo de Uso

```go
method := newtoncotes.NewClosedOrder3() // Regra de Simpson
f := func(x float64) float64 { return math.Sin(x) }
result := newtoncotes.Integrate(method, f, 0, math.Pi/2, 1e-6)
```

### Fundamentos Matemáticos

- **Ordem Fechada 2**: ∫f(x)dx ≈ h/2 * [f(a) + f(b)]
- **Ordem Fechada 3**: ∫f(x)dx ≈ h/3 * [f(a) + 4f(a+h) + f(b)]
- **Ordem Fechada 4**: ∫f(x)dx ≈ 3h/8 * [f(a) + 3f(a+h) + 3f(a+2h) + f(b)]

### Implementação

O método `integrateRecursive` divide recursivamente o intervalo [a, b] até que o erro estimado seja menor que a tolerância especificada. O erro é calculado comparando:
- Integral sobre o intervalo completo
- Soma das integrais sobre as duas metades do intervalo

## 2. Métodos de Quadratura Gaussiana

Os métodos de quadratura Gaussiana fornecem alta precisão ao escolher otimamente tanto as abscissas quanto os pesos. Cada método é especializado para diferentes tipos de integrais e funções peso.

### 2.1 Integração Gauss-Legendre

**Localização**: [gauss-legendre/](./gauss-legendre/)

Para integrais sobre intervalos finitos [-1, 1] com função peso w(x) = 1.

#### Ordens Disponíveis
- **TwoPoints**: Quadratura Gauss-Legendre de 2 pontos
- **ThreePoints**: Quadratura Gauss-Legendre de 3 pontos  
- **FourPoints**: Quadratura Gauss-Legendre de 4 pontos

#### Fundamentos Matemáticos
∫₋₁¹ f(x) dx ≈ Σᵢ wᵢ f(xᵢ)

Para intervalos arbitrários [a,b], usa transformação: x = (a+b+(a-b)s)/2

#### Características
- Exata para polinômios de grau ≤ 2n-1 (n = número de pontos)
- Integração adaptativa com controle de erro
- Transformação automática de intervalos

### 2.2 Integração Gauss-Hermite

**Localização**: [gauss-hermite/](./gauss-hermite/)

Para integrais sobre intervalos infinitos com função peso w(x) = e^(-x²).

#### Ordens Disponíveis
- **TwoPoints**: Usa ±√2/2 como abscissas, pesos = √π/2
- **ThreePoints**: Inclui zero e pontos simétricos
- **FourPoints**: Quatro pontos otimamente escolhidos

#### Fundamentos Matemáticos
∫₋∞^∞ e^(-x²) f(x) dx ≈ Σᵢ wᵢ f(xᵢ)

#### Aplicações
- Integrais com peso gaussiano
- Transformações exponenciais (métodos DINO)
- Problemas de probabilidade e estatística

### 2.3 Integração Gauss-Laguerre

**Localização**: [gauss-laguerre/](./gauss-laguerre/)

Para integrais sobre intervalos semi-infinitos [0, ∞) com função peso w(x) = e^(-x).

#### Ordens Disponíveis
- **TwoPoints**: Abscissas 2±√2, pesos (2±√2)/4
- **ThreePoints**: Três pontos incluindo aproximadamente 0.415, 2.294, 6.289
- **FourPoints**: Quatro pontos otimamente distribuídos

#### Fundamentos Matemáticos
∫₀^∞ e^(-x) f(x) dx ≈ Σᵢ wᵢ f(xᵢ)

#### Aplicações
- Integrais com decaimento exponencial
- Transformadas de Laplace
- Problemas de física estatística

### 2.4 Integração Gauss-Chebyshev

**Localização**: [gauss-chebyshev/](./gauss-chebyshev/)

Para integrais sobre [-1, 1] com função peso w(x) = 1/√(1-x²).

#### Ordens Disponíveis
- **TwoPoints**: Usa ±√2/2 como abscissas
- **ThreePoints**: Inclui 0, ±√3/2
- **FourPoints**: Quatro pontos de Chebyshev

#### Fundamentos Matemáticos
∫₋₁¹ f(x)/√(1-x²) dx ≈ Σᵢ wᵢ f(xᵢ)

Para N pontos: wᵢ = π/N (pesos iguais)

#### Características Especiais
- Todos os pesos são iguais (π/N)
- Relacionado aos polinômios de Chebyshev
- Excelente para aproximação de funções

## 3. Métodos DINO (Transformações Exponenciais)

**Localização**: [dino/](./dino/)

Os métodos DINO (Double exponential INtegratiOn) usam transformações de variáveis exponenciais para lidar com integrais difíceis, particularmente aquelas com singularidades nos extremos ou domínios infinitos.

### Métodos Disponíveis

#### 3.1 DINO Simples (Exponencial Simples)
**Tipo**: `DinoSimples`

Usa a transformação: x = (a+b+(b-a)tanh(s))/2

**Características:**
- **Propósito**: Lida com integrais com singularidades nos extremos
- **Transformação**: Mapeia (-∞, ∞) → [a, b]
- **Backend**: Usa quadratura Gauss-Hermite de 4 pontos
- **Derivada**: dx/ds = (b-a)/(2cosh²(s))

**Vantagens:**
- Converge rapidamente para integrandos com singularidades nos extremos
- Transforma integrais problemáticos em formas mais tratáveis
- Decaimento exponencial nas bordas do intervalo

#### 3.2 DINO Duo (Exponencial Dupla)
**Tipo**: `DinoDuo`

Usa a transformação: x = (a+b+(b-a)tanh(π/2·sinh(s)))/2

**Características:**
- **Propósito**: Convergência superior para integrandos suaves
- **Transformação**: Decaimento duplo exponencial nas bordas
- **Backend**: Usa quadratura Gauss-Hermite de 4 pontos
- **Derivada**: dx/ds = π(b-a)cosh(s)/(4cosh²(π/2·sinh(s)))

**Vantagens:**
- Convergência ainda mais rápida que DINO simples
- Ideal para funções muito suaves
- Decaimento super-exponencial

### Implementação dos Métodos DINO

Ambos os métodos seguem o mesmo padrão:

1. **Transformação de Variáveis**: Converte o intervalo [a,b] para (-∞,∞)
2. **Multiplicação pelo Jacobiano**: Inclui dx/ds na integral
3. **Peso Gaussiano**: Multiplica por e^(s²) para cancelar o peso de Hermite
4. **Quadratura Gauss-Hermite**: Aplica integração com 4 pontos

### Algoritmo Adaptativo

```go
func IntegrateDino(calculator DinoCalculator, f func(float64) float64, a, b float64) *result.IntegrateResult
```

- **Subdivisão Recursiva**: Divide intervalos até atingir tolerância (1e-5)
- **Estimativa de Erro**: Compara integral completa com soma das metades
- **Critério de Parada**: Erro < tolerância OU intervalo < 1e-9

### Exemplo de Uso

```go
// Exponencial simples - para singularidades nos extremos
dinoSimple := dino.NewDinoSimples()
result := dino.IntegrateDino(dinoSimple, f, a, b)

// Exponencial dupla - para funções muito suaves
dinoDuo := dino.NewDinoDuo()
result := dino.IntegrateDino(dinoDuo, f, a, b)
```

### Fundamentos Matemáticos

Ambos os métodos transformam a integral usando:
∫ₐᵇ f(x) dx = ∫₋∞^∞ f(x(s)) · x'(s) ds

Então aplicam quadratura Gauss-Hermite:
∫₋∞^∞ e^(-s²) g(s) ds ≈ Σᵢ wᵢ g(sᵢ)

onde g(s) = e^(s²) · f(x(s)) · x'(s)

## 4. Estrutura de Resultado

**Localização**: [result/](./result/)

Todos os métodos de integração retornam resultados usando a estrutura `IntegrateResult`:

```go
type IntegrateResult struct {
    Result          float64  // Valor da integral computada
    NumOfIterations int      // Número de subdivisões recursivas
}
```

### Informações Fornecidas

- **Result**: O valor numérico da integral calculada
- **NumOfIterations**: Indica o esforço computacional (número de subdivisões)

### Função Construtora

```go
func NewIntegrateResult(result float64, iterations int) *IntegrateResult
```

## 5. Executando Testes

Cada pacote inclui suítes de teste abrangentes que verificam a precisão e robustez dos métodos implementados.

### Comandos de Teste

```bash
# Testar todos os pacotes
go test ./...

# Testar pacote específico
go test ./newton-cotes
go test ./gauss-legendre
go test ./gauss-hermite
go test ./gauss-laguerre
go test ./gauss-chebyshev
go test ./dino

# Executar com saída verbosa
go test -v ./...

# Executar testes com cobertura
go test -cover ./...
```

### Funções de Teste Incluídas

Os testes incluem várias funções matemáticas desafiadoras:

1. **Funções Trigonométricas**:
   - `sin(x)` de 0 a π/2 (resultado esperado: 1)
   - `cos(x)` e combinações

2. **Funções Polinomiais**:
   - `x²`, `x³`, `x⁴` (para testar exatidão)
   - Polinômios compostos

3. **Funções Complexas**:
   - `(sin(2x) + 4x² + 3x)²` de 0 a 1
   - Funções logarítmicas: `ln(x+2)²`

4. **Casos Especiais**:
   - Funções com singularidades
   - Integrais impróprias
   - Funções oscilatórias

### Tolerâncias e Precisão

- **Newton-Cotes**: Tolerância padrão 1e-6
- **Gauss**: Tolerância adaptativa baseada na ordem
- **DINO**: Tolerância fixa 1e-5

## Estrutura do Projeto

```
├── newton-cotes/           # Métodos de integração Newton-Cotes
│   ├── newton_cotes.go     # Implementações dos métodos
│   └── newton_cotes_test.go # Testes abrangentes
├── gauss-legendre/         # Quadratura Gauss-Legendre
│   ├── gauss_legendre.go   # Implementação Legendre
│   └── gauss_legendre_test.go # Testes específicos
├── gauss-hermite/          # Quadratura Gauss-Hermite
│   ├── gauss_hermite.go    # Implementação Hermite
│   └── gauss_hermite_test.go # Testes específicos
├── gauss-laguerre/         # Quadratura Gauss-Laguerre
│   ├── gauss_laguerre.go   # Implementação Laguerre
│   └── gauss_laguerre_test.go # Testes específicos
├── gauss-chebyshev/        # Quadratura Gauss-Chebyshev
│   ├── gauss_chebyshev.go  # Implementação Chebyshev
│   └── gauss_chebyshev_test.go # Testes específicos
├── dino/                   # Métodos exponenciais DINO
│   ├── dino.go            # Implementações DINO
│   └── dino_test.go       # Testes DINO
├── result/                 # Estrutura de dados de resultado
│   └── result.go          # Definição da estrutura
├── go.mod                  # Definição do módulo Go
├── go.sum                  # Checksums das dependências
└── README.md              # Esta documentação
```
