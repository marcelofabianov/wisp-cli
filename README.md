# wisp-cli

Uma ferramenta de linha de comando para explorar e aprender sobre os *value objects* do pacote `wisp`.

Esta CLI serve como uma documentação interativa, permitindo que você visualize rapidamente o conceito, um exemplo de uso prático e a URL da documentação para qualquer tipo disponível no pacote `wisp`.

## Demonstração Rápida

```sh
wisp-cli cpf
```

### Conceito
O wisp.CPF é um value object que representa um CPF brasileiro.

Ele garante que qualquer CPF no sistema seja válido, verificando não apenas o formato de 11 dígitos, mas também a lógica matemática dos dígitos verificadores. Ele normaliza a entrada (removendo máscaras como pontos e hífens) e oferece métodos para reformatar a exibição.

### Exemplo de Uso

```go
package main

import (
	"fmt"
	"log"

	"[github.com/marcelofabianov/wisp](https://github.com/marcelofabianov/wisp)"
)

func main() {
	// Exemplo de criação com um CPF válido e formatado
	validCPF, err := wisp.NewCPF("862.226.160-38")
	if err != nil {
		log.Fatalf("Falha inesperada: %v", err)
	}

	fmt.Printf("CPF Válido:\n")
	fmt.Printf("  - Valor Canônico: %s\n", validCPF.String())
	fmt.Printf("  - Valor Formatado: %s\n\n", validCPF.Formatted())

	// Exemplo de tentativa de criação com um CPF inválido
	fmt.Println("Tentando criar com um CPF inválido...")
	_, err = wisp.NewCPF("111.111.111-11")
	if err != nil {
		fmt.Printf("  -> Erro esperado recebido: %v\n", err)
	}
}
```

## Instalação

Existem duas maneiras principais de instalar a `wisp-cli`:

### Opção 1: `go install` (Recomendado para Desenvolvedores Go)

Se você tem o Go instalado e configurado em seu `PATH`, pode instalar a última versão diretamente do GitHub com um único comando:

```sh
go install github.com/marcelofabianov/wisp-cli/cmd/wisp@latest
```

Isso irá compilar e instalar o binário `wisp` no seu diretório `$GOPATH/bin`.

### Opção 2: Baixar o Binário (Para Todos os Usuários)

Você pode baixar a versão mais recente para o seu sistema operacional diretamente da página de **Releases** do repositório no GitHub.

## Como Usar

A `wisp-cli` é simples de usar. O principal comando espera o nome de um tipo do pacote `wisp` como argumento.

**Listando Todos os Tipos Disponíveis**

Para ver todos os tipos que a CLI conhece, execute o comando sem argumentos:

```sh
wisp-cli
```

Saída:

```
Nenhum tipo especificado. Tipos disponíveis:
  - cpf
  - money

Use 'wisp-cli [type]' para ver os detalhes.
```

**Exibindo Detalhes de um Tipo**

Para ver o conceito, exemplo e documentação de um tipo específico, passe seu nome como argumento. O nome do tipo não é case-sensitive.

Exemplo com `money`:

```sh
wisp-cli money
```

```
--- Conceito ---
O wisp.Money é um value object para representar valores monetários de forma segura.

Sua principal vantagem é evitar o uso de float64 para cálculos financeiros, prevenindo erros de precisão. Ele armazena o valor internamente como um int64 na sua menor unidade (ex: centavos) e sempre associa o montante a um wisp.Currency, garantindo que 100 BRL seja diferente de 100 USD.

--- Exemplo de Uso ---
package main
...

--- Documentação ---
Para mais detalhes, veja: [https://pkg.go.dev/github.com/marcelofabianov/wisp#Money](https://pkg.go.dev/github.com/marcelofabianov/wisp#Money)
```

**Compilar nova versão**

```
go build -o wisp-cli ./cmd/wisp
```


## Licença

Este projeto é licenciado sob a Licença MIT.
