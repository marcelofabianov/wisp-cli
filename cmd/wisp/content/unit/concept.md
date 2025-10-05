O `wisp.Unit` é a base para o *value object* `Quantity`. Ele implementa um sistema de registro extensível que permite à aplicação definir seu próprio conjunto de unidades de medida válidas.

**Principais Funcionalidades:**

* **Sistema de Registro**: A aplicação define uma "lista de permissões" de unidades através da função `RegisterUnits` (ex: "KG", "UN", "M2").
* **Validação Centralizada**: O método `IsValid()` verifica se uma `Unit` foi previamente registrada, garantindo que o `Quantity` nunca seja criado com uma unidade desconhecida.
* **Flexibilidade**: Torna o tipo `Quantity` agnóstico ao domínio, permitindo que qualquer tipo de unidade de medida seja utilizado, desde que seja registrado primeiro.
