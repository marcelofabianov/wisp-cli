O `wisp.NullableTime` é um *value object* crucial para representar timestamps que são opcionais, ou seja, que podem ser nulos.

Ele resolve um problema fundamental do `time.Time` padrão do Go, cujo valor "zero" (`0001-01-01`) é uma data válida, e não uma representação de `NULL`. Isso causa problemas ao interagir com bancos de dados e APIs JSON.

**Principais Funcionalidades e Garantias:**

* **Estado Nulo Explícito**: O `NullableTime` encapsula um `time.Time` e um booleano `Valid`. Se `Valid` for `false`, o objeto representa um valor `NULL`, tornando o estado nulo explícito e seguro de se verificar com o método `IsZero()`.

* **Integração Perfeita**: Lida de forma transparente com a serialização e desserialização de valores nulos.
    * **JSON**: Serializa para `null` quando `Valid` é `false`.
    * **SQL**: Mapeia para `NULL` no banco de dados quando `Valid` é `false`.

É a ferramenta ideal para campos como `deleted_at` (soft delete), `archived_at` (arquivamento), `processed_at` (data de processamento de um lote), ou qualquer outro timestamp que só existe após uma ação específica ocorrer.
