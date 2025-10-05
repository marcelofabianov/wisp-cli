O `wisp.TimeRange` é um *value object* que representa um intervalo de tempo contínuo dentro de um dia, composto por um `TimeOfDay` de início e um de fim.

Sua principal responsabilidade é garantir a integridade do período, validando que o horário de início seja sempre anterior ao de término. É a base para a criação de conceitos mais complexos, como horários de funcionamento.

**Principais Funcionalidades e Garantias:**

* **Validação de Intervalo**: O construtor `NewTimeRange` retorna um erro se o horário de início não for estritamente anterior ao de fim, prevenindo a criação de períodos inválidos.
* **Lógica de Contenção**: O método `Contains(t TimeOfDay)` oferece uma forma clara de verificar se um determinado horário está dentro do intervalo. A verificação é inclusiva no início e exclusiva no fim `[start, end)`, o que é um padrão comum para agendamentos.
* **Imutabilidade**: O objeto é imutável, garantindo que um intervalo de tempo, uma vez criado, não possa ser alterado acidentalmente.
