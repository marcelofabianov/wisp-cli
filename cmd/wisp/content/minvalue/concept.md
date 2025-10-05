O `wisp.MinValue` é um *value object* que representa um valor numérico (`int64`) que não pode cair abaixo de um limite mínimo.

É o complemento do `BoundedValue`. Ideal para modelar conceitos de negócio onde um valor tem um piso, mas não necessariamente um teto, como o nível mínimo de estoque de segurança de um produto ou o saldo mínimo em uma conta.

**Principais Funcionalidades:**

* **Garantia de Limite**: O método `Subtract(amount)` retorna o erro `wisp.ErrValueSubceedsMin` se a operação violar o limite mínimo.
* **Imutabilidade**: Todas as operações retornam uma nova instância do objeto.
* **API Expressiva**: Métodos como `IsAtMin()` fornecem uma maneira clara de consultar o estado do valor.
