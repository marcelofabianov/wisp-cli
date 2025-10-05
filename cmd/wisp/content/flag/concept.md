O `wisp.Flag[T]` é um *value object* genérico e poderoso, projetado para modelar um estado que só pode assumir um de dois valores permitidos.

Utilizando Generics do Go, ele pode ser usado com `string`, `int`, `bool` ou qualquer tipo comparável. Isso o torna ideal para representar conceitos como status (`"published"`/`"draft"`), confirmações (`1`/`0`), ou qualquer estado binário customizado.

**Principais Funcionalidades:**

* **Criação Segura**: O construtor `NewFlag(current, primary, secondary)` valida que o valor atual seja um dos dois estados permitidos.
* **API Expressiva**: Métodos como `IsPrimary()`, `IsSecondary()` e `Is(value)` oferecem maneiras claras e seguras de verificar o estado da flag sem usar "magic strings" ou valores literais no código de negócio.
