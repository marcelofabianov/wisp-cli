O `wisp.Percentage` é um *value object* para representar valores percentuais com alta precisão, crucial para cálculos financeiros e fiscais.

Ele resolve o problema de imprecisão de tipos de ponto flutuante (`float64`) armazenando o valor internamente como um `int64` escalado (com 4 casas decimais de precisão).

**Principais Funcionalidades:**

* **Criação Segura**: O construtor `NewPercentageFromFloat` aceita um `float64` (ex: `0.155` para 15.5%) e o converte de forma segura para a representação interna.
* **Precisão Matemática**: Garante que os cálculos sejam exatos, sem os erros de arredondamento de `float64`.
* **Lógica de Aplicação**: O método `ApplyTo(Money)` é a sua principal funcionalidade. Ele calcula o valor percentual de um montante `wisp.Money`, utilizando arredondamento bancário ("half to even") para garantir justiça financeira.
