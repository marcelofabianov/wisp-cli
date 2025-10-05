O `wisp.TimeOfDay` é um *value object* que representa uma hora específica do dia (ex: "17:30") de forma isolada, sem qualquer associação com uma data ou fuso horário.

É o bloco de construção fundamental para lógicas de negócio que dependem apenas do horário, como a definição de horários de funcionamento ou agendamento de tarefas diárias.

**Principais Funcionalidades e Garantias:**

* **Validação de Formato**: O construtor `ParseTimeOfDay` garante que a entrada esteja estritamente no formato "HH:MM" (ex: "09:00", "23:59"), rejeitando outros formatos para manter a consistência.

* **Armazenamento Eficiente**: Internamente, o valor é armazenado como um `int` representando o total de minutos desde a meia-noite. Isso torna as comparações (`Before`, `After`) extremamente rápidas e livres de complexidades de parsing.

* **API Clara**: Oferece métodos como `Hour()` e `Minute()` para extrair os componentes do horário de forma segura.
