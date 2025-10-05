O `wisp.Length` é um *value object* para representar medidas de comprimento. Similar ao `Weight`, ele foi projetado para garantir precisão e segurança em cálculos.

**Principais Funcionalidades:**

* **Armazenamento Preciso**: O valor é armazenado internamente como um `int64` em micrômetros para evitar erros de ponto flutuante.
* **Conversão de Unidades**: O método `In(unit)` permite converter o valor entre as unidades suportadas: Metros (`m`), Centímetros (`cm`), Milímetros (`mm`), Quilômetros (`km`), Polegadas (`in`) e Pés (`ft`).
* **Aritmética Segura**: Fornece métodos como `Add()` e `Subtract()` para operações de medida.
* **Validação**: `NewLength` garante que um comprimento não pode ser criado com um valor inicial negativo.
