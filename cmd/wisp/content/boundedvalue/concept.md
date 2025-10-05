O `wisp.BoundedValue` é um *value object* que representa um valor numérico (`int64`) que não pode exceder um limite máximo, com um mínimo implícito de zero.

É ideal para modelar conceitos de negócio como capacidade de estoque, cotas de uso ou limites de saldo em uma carteira.

**Principais Funcionalidades:**

* **Garantia de Limite**: O método `Add(amount)` retorna o erro `wisp.ErrValueExceedsMax` se a operação violar o limite máximo, permitindo um controle de fluxo de negócio claro e explícito.
* **Imutabilidade**: Todas as operações (`Add`, `Subtract`, `Set`) retornam uma nova instância do objeto, garantindo que o estado seja sempre previsível.
* **API Expressiva**: Métodos como `AvailableSpace()` e `IsFull()` fornecem uma maneira clara de consultar o estado do valor em relação ao seu limite.
