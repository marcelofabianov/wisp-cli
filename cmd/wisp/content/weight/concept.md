O `wisp.Weight` é um *value object* para representar medidas de massa (peso) de forma segura e precisa, evitando os erros de arredondamento de `float64`.

**Principais Funcionalidades:**

* **Armazenamento Preciso**: O valor é armazenado internamente como um `int64` na menor unidade comum (miligramas), garantindo a exatidão em cálculos.
* **Conversão de Unidades**: O método `In(unit)` permite converter o valor de forma segura entre as unidades suportadas: Quilogramas (`kg`), Gramas (`g`), Libras (`lb`) e Onças (`oz`).
* **Aritmética Segura**: Oferece métodos como `Add()` e `Subtract()` para realizar operações matemáticas, onde o resultado de uma subtração pode ser negativo para representar cenários de déficit.
* **Validação**: O construtor `NewWeight` garante que um peso não pode ser criado com um valor inicial negativo.
