# Relatório Unidade 3

## 1) Métodos de Potência

Os métodos de potência são algoritmos iterativos usados para encontrar autovalores e autovetores de matrizes. Este projeto implementa três variações principais:

### 1.1) Método de Potência Regular (`potencia_regular.go`)

**Objetivo**: Encontra o maior autovalor (em valor absoluto) e seu autovetor correspondente.

**Algoritmo**:
- Inicia com um vetor aproximação inicial `x_0`
- Aplica iterativamente: `y = A·x` e normaliza `x = y/||y||inf`
- O autovalor converge para `lambda = ||y||inf`
- Critério de parada: `abs((lambda_k - lambda_k-1)/lambda_k) < tolerância`

**Parâmetros**:
- `matrixA`: Matriz quadrada de entrada
- `initialVector`: Vetor inicial para iteração
- `tolerance`: Tolerância para convergência
- `maxIterations`: Número máximo de iterações

### 1.2) Método de Potência Inversa (`potencia_inversa.go`)

**Objetivo**: Encontra o menor autovalor (em valor absoluto) e seu autovetor correspondente.

**Algoritmo**:
- Calcula a matriz inversa `A^-1`
- Aplica o método de potência regular em `A^-1`
- O menor autovalor de `A` é o inverso do maior autovalor de `A^-1`
- Utiliza decomposição LU interna do gonum para inversão

**Vantagens**: Converge rapidamente para o menor autovalor
**Limitações**: Requer que a matriz seja invertível

### 1.3) Método de Potência com Deslocamento (`potencia_desloc.go`)

**Objetivo**: Encontra autovalores próximos a um valor específico `μ` (deslocamento).

**Algoritmo**:
- Cria matriz deslocada: `A_hat = A - μI`
- Aplica potência inversa em `A_hat`
- O autovalor de `A` é: `lambda = 1/lambda_A_hat + μ`
- Permite buscar autovalores em regiões específicas do espectro

**Parâmetros adicionais**:
- `mu`: Valor do deslocamento (aproximação do autovalor desejado)

**Uso estratégico**: Útil quando se conhece aproximadamente onde estão os autovalores de interesse.

### Estrutura Comum

Todos os métodos retornam um `PowerMethodResult` contendo:
- `Eigenvalue`: Autovalor calculado
- `Eigenvector`: Autovetor correspondente (normalizado)

## 2) Método de Householder + QR

Implementação completa para encontrar todos os autovalores e autovetores de matrizes simétricas.

### 2.1) Transformações de Householder (`householder.go`)

**Objetivo**: Reduzir uma matriz simétrica à forma tridiagonal, preservando autovalores.

**Algoritmo**:
1. Para cada coluna `i`, constrói uma matriz de reflexão `Hᵢ`
2. A reflexão zera elementos abaixo da diagonal secundária
3. Aplica: `A_{k+1} = H_i^T · A_k · H_i`
4. Acumula as transformações: `H = H_1 · H_2 · ... · H_n-2`

**Estruturas retornadas**:
- `T`: Matriz tridiagonal resultante
- `H`: Matriz acumulada das transformações de Householder

**Funções principais**:
- `householderMatrix()`: Cria matriz de reflexão para uma coluna específica
- `HouseholderMethod()`: Aplica todas as transformações sequencialmente
- `NewIdentityMatrix()`: Utilitário para criar matrizes identidade

### 2.2) Método QR (`qr_method.go`)

**Objetivo**: Encontrar todos os autovalores e autovetores da matriz tridiagonal.

**Algoritmo QR iterativo**:
1. Decompõe a matriz: `A = Q · R` (usando rotações de Givens)
2. Recompõe: `A_{k+1} = R · Q`
3. Acumula autovetores: `X = X · Q`
4. Repete até convergência (elementos sub-diagonais ≈ 0)

**Decomposição QR com Rotações de Givens**:
- Para cada elemento sub-diagonal `A[i,j]` com `i > j`
- Calcula rotação que zera o elemento: `cos = a/r`, `sen = -b/r`
- Aplica rotação à direita em `R` e à esquerda em `Q`

**Critério de convergência**:
- Soma dos valores absolutos dos elementos sub-diagonais < `epsilon`
- Indica que a matriz está suficientemente próxima da forma diagonal

**Estruturas retornadas**:
- `Lambda`: Matriz diagonal com autovalores na diagonal principal
- `X`: Matriz cujas colunas são os autovetores

### 2.3) Fluxo Completo de Execução

1. **Entrada**: Matriz simétrica `A`
2. **Householder**: `A → T` (tridiagonalização) + matriz `H`
3. **QR**: `T → Λ` (diagonalização) + autovetores `V`
4. **Resultado**: Autovalores em `Λ`, autovetores finais em `H·V`

### Vantagens da Abordagem

- **Estabilidade numérica**: Transformações ortogonais preservam normas
- **Eficiência**: Redução prévia à forma tridiagonal acelera o método QR
- **Completude**: Encontra todos os autovalores/autovetores
- **Precisão**: Adequado para matrizes simétricas de tamanho moderado

## 3) Estrutura do Projeto

```
power-methods/
├── common.go                 # Estruturas compartilhadas
├── potencia_regular.go       # Método de potência regular
├── potencia_inversa.go       # Método de potência inversa  
├── potencia_desloc.go        # Método de potência com deslocamento
└── *_test.go                 # Testes unitários para cada método

householder-qr/
├── householder.go            # Transformações de Householder
├── qr_method.go             # Método QR com rotações de Givens
└── *_test.go                # Testes para ambos os métodos
```

## 4) Dependências

- **gonum.org/v1/gonum/mat**: Biblioteca para operações com matrizes
- **Go 1.21+**: Versão mínima requerida

## 5) Execução dos Testes

```bash
# Testar métodos de potência
go test ./power-methods/

# Testar Householder + QR  
go test ./householder-qr/

# Executar todos os testes
go test ./...
```
