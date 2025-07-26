# Relatório Unidade 4

## Sumário

1. [Visão Geral](#visão-geral)
2. [Estrutura do Projeto](#estrutura-do-projeto)
3. [Método de Euler](#método-de-euler)
4. [Métodos de Runge-Kutta](#métodos-de-runge-kutta)
5. [Métodos Preditor-Corretor](#métodos-preditor-corretor)
6. [Testes e Validação](#testes-e-validação)
7. [Como Executar](#como-executar)
8. [Dependências](#dependências)

## Visão Geral

Este projeto implementa soluções numéricas para problemas de valor inicial (PVI) da forma:

```
y'(t) = f(t, y(t))
y(t₀) = y₀
```

Onde:

- `y'(t)` é a derivada de y em relação ao tempo t
- `f(t, y(t))` é uma função conhecida
- `y₀` é a condição inicial no tempo `t₀`

Os métodos implementados permitem aproximar a solução y(t) em pontos discretos, calculando valores sucessivos y₁, y₂, ..., yₙ.

## Estrutura do Projeto

```
unidade4/
├── types/                  # Definições de tipos e interfaces
│   └── types.go           # DerivativeFunc e interfaces base
├── euler-method/          # Implementação dos métodos de Euler
│   ├── explicit_euler.go  # Euler explícito
│   ├── implicit_euler.go  # Euler implícito
│   ├── result.go         # Estrutura de resultado
│   └── *_test.go         # Testes unitários
├── runge-kutta/          # Implementação dos métodos de Runge-Kutta
│   ├── runge_kutta_second.go  # RK de 2ª ordem
│   ├── runge_kutta_third.go   # RK de 3ª ordem
│   ├── runge_kutta_fourth.go  # RK de 4ª ordem (clássico)
│   ├── result.go             # Estrutura de resultado
│   └── *_test.go             # Testes unitários
└── predictor-corrector/  # Implementação dos métodos preditor-corretor
    ├── second_order.go   # Adams-Bashforth de 2ª ordem
    ├── third_order.go    # Adams-Bashforth de 3ª ordem
    ├── fourth_order.go   # Adams-Bashforth de 4ª ordem
    ├── result.go        # Estrutura de resultado
    └── *_test.go        # Testes unitários
```

## Método de Euler

O método de Euler é o mais simples dos métodos numéricos para EDOs. Implementamos duas variantes:

### Euler Explícito

**Fórmula:** `y_{n+1} = y_n + h·f(t_n, y_n)`

**Características:**

- Método de primeira ordem
- Simples de implementar
- Erro de truncamento local O(h²)
- Condicionalmente estável

**Implementação:**

```go
type ExplicitEuler struct{}

func (em *ExplicitEuler) Execute(
    ctx context.Context,
    fc types.DerivativeFunc,
    initialCondition *mat.VecDense,
    initialTime, h float64,
) *EulerResult
```

### Euler Implícito

**Fórmula:** `y_{n+1} = y_n + h·f(t_{n+1}, y_{n+1})`

**Características:**

- Método de primeira ordem
- Mais estável que o explícito
- Requer resolução de equação não-linear
- Implementado usando predição com Euler explícito

**Implementação:**
O método implícito usa uma abordagem de predição-correção:

1. **Predição:** Usa Euler explícito para obter estimativa inicial
2. **Correção:** Aplica a fórmula implícita usando a predição

## Métodos de Runge-Kutta

Os métodos de Runge-Kutta oferecem maior precisão que o método de Euler, avaliando a derivada em múltiplos pontos dentro do intervalo.

### Runge-Kutta de 2ª Ordem (Método do Ponto Médio)

**Fórmulas:**

```
k₁ = h·f(t_n, y_n)
k₂ = h·f(t_n + h/2, y_n + k₁/2)
y_{n+1} = y_n + k₂
```

**Características:**

- Erro de truncamento local O(h³)
- Boa precisão para h moderado
- Balanço entre simplicidade e precisão

### Runge-Kutta de 3ª Ordem

**Fórmulas:**

```
k₁ = h·f(t_n, y_n)
k₂ = h·f(t_n + h/2, y_n + k₁/2)
k₃ = h·f(t_n + h, y_n - k₁ + 2k₂)
y_{n+1} = y_n + (k₁ + 4k₂ + k₃)/6
```

**Características:**

- Erro de truncamento local O(h⁴)
- Melhor precisão que RK2
- Compromisso entre precisão e custo computacional

### Runge-Kutta de 4ª Ordem (Clássico)

**Fórmulas:**

```
k₁ = h·f(t_n, y_n)
k₂ = h·f(t_n + h/2, y_n + k₁/2)
k₃ = h·f(t_n + h/2, y_n + k₂/2)
k₄ = h·f(t_n + h, y_n + k₃)
y_{n+1} = y_n + (k₁ + 2k₂ + 2k₃ + k₄)/6
```

**Características:**

- Erro de truncamento local O(h⁵)
- Excelente precisão
- Método mais utilizado na prática
- Requer 4 avaliações da função por passo

## Métodos Preditor-Corretor

Os métodos preditor-corretor combinam eficiência computacional com boa estabilidade, usando informações de pontos anteriores.

### Adams-Bashforth (Preditor)

Utiliza interpolação polinomial baseada em valores anteriores da derivada para predizer o próximo valor.

**2ª Ordem:**

```
y_{n+1}^P = y_n + h/2 · (3f_n - f_{n-1})
```

**3ª Ordem:**

```
y_{n+1}^P = y_n + h/12 · (23f_n - 16f_{n-1} + 5f_{n-2})
```

**4ª Ordem:**

```
y_{n+1}^P = y_n + h/24 · (55f_n - 59f_{n-1} + 37f_{n-2} - 9f_{n-3})
```

### Adams-Moulton (Corretor)

Usado para corrigir a predição usando o valor predito.

**Características:**

- Métodos multistep
- Requerem valores iniciais calculados por métodos single-step (Runge-Kutta)
- Eficientes para problemas de longa duração
- Boa estabilidade quando bem implementados

**Implementação:**
O projeto implementa uma abordagem adaptativa que:

1. Calcula valores iniciais usando Runge-Kutta
2. Aplica predição Adams-Bashforth
3. Verifica convergência dentro de tolerância especificada
4. Itera até convergência ou máximo de iterações

## Testes e Validação

Todos os métodos são testados com problemas de valor inicial conhecidos:

### Problema Teste Principal: PVI-1

```
y' = (2/3)y
y(0) = 2
```

**Solução Analítica:** `y(t) = 2e^(2t/3)`

Este problema permite validação direta dos métodos numéricos contra a solução exata.

### Estrutura de Testes

Cada método possui testes unitários que verificam:

- Precisão em comparação com solução analítica
- Comportamento com diferentes tamanhos de passo
- Estabilidade numérica
- Casos extremos

## Como Executar

### Pré-requisitos

- Go 1.24.5 ou superior

````

### Executar Testes

```bash
# Todos os testes
go test ./...

# Testes específicos
go test ./euler-method
go test ./runge-kutta
go test ./predictor-corrector

# Testes com saída detalhada
go test -v ./...
````
