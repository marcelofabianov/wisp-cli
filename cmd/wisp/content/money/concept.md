O `wisp.Money` é um value object para representar valores monetários de forma segura.

Sua principal vantagem é evitar o uso de `float64` para cálculos financeiros, prevenindo erros de precisão. Ele armazena o valor internamente como um `int64` na sua menor unidade (ex: centavos) e sempre associa o montante a um `wisp.Currency`, garantindo que `100 BRL` seja diferente de `100 USD`.
