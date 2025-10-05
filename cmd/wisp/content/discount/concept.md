O `wisp.Discount` é um *value object* de negócio que encapsula a lógica de um desconto, que pode ser de dois tipos: um valor **fixo** (`wisp.Money`) ou um valor **percentual** (`wisp.Percentage`).

Esta natureza polimórfica permite que o sistema trate diferentes tipos de desconto de forma uniforme e segura.

**Principais Funcionalidades:**

* **Polimorfismo Seguro**: Construtores separados (`NewFixedDiscount` e `NewPercentageDiscount`) garantem a criação de um tipo de desconto válido.
* **Lógica Centralizada**: O método `ApplyTo(Money)` é o coração do tipo. Ele aplica o desconto (seja ele fixo ou percentual) a um valor `Money` e retorna o preço final, garantindo que o cálculo seja feito corretamente e que o resultado nunca seja negativo.
* **Serialização Inteligente**: A representação JSON é customizada para indicar o tipo e o valor do desconto, facilitando a comunicação com APIs.
