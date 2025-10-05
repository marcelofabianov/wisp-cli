O `wisp.RangedValue` é o mais completo dos tipos de valor delimitado. Ele representa um valor numérico (`int64`) que deve permanecer contido dentro de um intervalo `[min, max]`.

Ele une os conceitos de `MinValue` e `BoundedValue`, sendo perfeito para modelar regras de negócio que possuem tanto um piso quanto um teto, como a quantidade de alunos em uma turma, a temperatura de um equipamento ou a dosagem de um componente.

**Principais Funcionalidades:**

* **Validação Dupla**: O construtor `NewRangedValue` valida se `min <= current <= max`, garantindo a integridade do valor desde a sua criação.
* **Operações Seguras**: Os métodos `Add()` e `Subtract()` retornam os erros `ErrValueExceedsMax` e `ErrValueSubceedsMin`, respectivamente, se qualquer um dos limites for violado.
* **Imutabilidade**: Garante que o estado seja sempre previsível, com todas as operações retornando uma nova instância.
