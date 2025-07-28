# Solucionador de Problema de Valor de Contorno (PVC)

## Visão Geral

Este projeto implementa um solucionador numérico para problemas de valor de contorno usando o método de diferenças finitas.

## Desenvolvimento dos coeficientes para o exemplo

<img width="1600" height="804" alt="image" src="https://github.com/user-attachments/assets/b6c50b1f-075b-4a66-a27a-5a45331fbb60" />

## Estrutura do Projeto

```
.
├── main.go                     # Ponto de entrada principal da aplicação
├── go.mod                      # Definição do módulo Go
├── go.sum                      # Checksums do módulo Go
├── plotter.py                  # Ferramenta de visualização Python
├── go_output_plot.png          # Saída do gráfico gerado
├── pvc-processor/              # Pacote principal do solucionador PVC
│   ├── pvc-processor.go        # Implementação principal do solucionador
│   └── pvc-processor_test.go   # Testes unitários
├── result/                     # Estruturas de dados e tipos
│   └── result.go              # Tipos de entrada/saída e erros
└── log/                       # Diretório de logs da aplicação
    └── *.log                  # Arquivos de log formatados em JSON
```

## Instalação

### Pré-requisitos

- Go 1.24.5 ou superior
- Python 3.x com matplotlib (para visualização)

### Dependências

```bash
# Instalar dependências Go
go mod tidy

# Venv
python -m venv .venv
source .venv/bin/activate

# Instalar dependências Python (para plotagem)
pip install matplotlib
```

## Uso

### Uso Básico

```bash
# Executar o solucionador
go run main.goa

# Executar com visualização
python plotter.py
```

### Configuração

O solucionador é configurado através da struct `PVCInput` em `main.go`:

> Compativel com qualquer EDO de 1 dimensão

```go
pvcInput := &result.PVCInput{
    MaskValues:   []float64{65, -201, 135},  // Coeficientes do stencil de três pontos
    A:            0.0,                        // Contorno esquerdo (x = a)
    B:            2.0,                        // Contorno direito (x = b)
    StepSize:     0.1,                        // Espaçamento da grade (h)
    InitialCond:  []float64{10, 1},          // Valores de contorno [u(a), u(b)]
    DefaultValue: 2.0,                        // Valor padrão para o vetor RHS
}
```

### Explicação dos Parâmetros

- **MaskValues**: Coeficientes do stencil de diferenças finitas de três pontos `[u_{i-1}, u_i, u_{i+1}]`
- **A, B**: Fronteiras do domínio definindo o intervalo [a, b]
- **StepSize**: Espaçamento da grade h = (b-a)/n
- **InitialCond**: Condições de contorno `[u(a), u(b)]`
- **DefaultValue**: Valor padrão para o vetor do lado direito b

## Algoritmo

O solucionador implementa o método de diferenças finitas para resolver problemas de valor de contorno da forma:

```
au_{i-1} + bu_i + cu_{i+1} = f_i
```

Onde:

- `a`, `b`, `c` são os valores da máscara
- `f_i` é o valor do lado direito
- Condições de contorno são aplicadas nas extremidades

### Processo de Solução

1. **Geração da Grade**: Cria uma grade uniforme com tamanho de passo h
2. **Montagem da Matriz**: Constrói a matriz de coeficientes usando o stencil de três pontos
3. **Condições de Contorno**: Aplica condições de contorno de Dirichlet
4. **Solução do Sistema**: Resolve o sistema linear Ax = b usando o solucionador Gonum
5. **Saída**: Retorna o vetor solução junto com as matrizes do sistema

## Saída

O programa produz os valores da solução em cada ponto da grade:

```
10.000000    # u(x=0.0) - condição de contorno
8.234567     # u(x=0.1)
6.789012     # u(x=0.2)
...          # Pontos interiores
1.000000     # u(x=2.0) - condição de contorno
```

## Visualização

Use o script Python incluído para visualizar os resultados:

```bash
python plotter.py
```

Isso gera um gráfico mostrando a solução numérica ao longo do domínio.

## Log

A aplicação usa log estruturado em JSON com os seguintes níveis:

- **Info**: Informações gerais de operação
- **Debug**: Passos detalhados de computação
- **Error**: Condições de erro e falhas

Os logs são salvos em arquivos `log/pvc_AAAAMMDD_HHMMSS.log`.

## Testes

Execute o conjunto de testes:

```bash
# Executar todos os testes
go test ./...

# Executar testes com saída verbosa
go test -v ./...

# Executar testes para pacote específico
go test ./pvc-processor
```
