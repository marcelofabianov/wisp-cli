O `wisp.Quantity` é um *value object* genérico e poderoso para representar uma quantidade numérica associada a uma unidade de medida.

Diferente de `Weight` ou `Length` que possuem unidades fixas, `Quantity` é extensível através do sistema de registro do `wisp.Unit`, permitindo que a aplicação defina suas próprias unidades de medida.

**Principais Funcionalidades:**

* **Precisão Configurável**: Evita erros de `float64` armazenando o valor como um `int64` escalado. A precisão pode ser definida globalmente com `SetDefaultPrecision` ou por instância com `NewQuantityWithPrecision`.
* **Segurança de Unidade**: Garante que uma quantidade só possa ser criada com uma `wisp.Unit` previamente registrada.
* **Integração Financeira**: O método `MultiplyByMoney(Money)` permite calcular um valor financeiro total a partir de uma quantidade e um preço unitário de forma segura e com arredondamento correto.
