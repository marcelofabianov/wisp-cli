O `wisp.BusinessHours` é um *value object* de alto nível que modela os horários de funcionamento de um negócio ao longo da semana.

Ele funciona como um agregador, utilizando `wisp.DayOfWeek` como chave e `wisp.TimeRange` como valor para definir um cronograma de operação.

**Principais Funcionalidades e Garantias:**

* **Modelo de Domínio Rico**: Representa um conceito de negócio complexo (horário de funcionamento) de forma clara e coesa.
* **Lógica de Negócio Centralizada**: O método `IsOpen(time.Time)` encapsula toda a lógica para verificar se um instante específico (data e hora) cai dentro do expediente, considerando o dia da semana e o intervalo de tempo.
* **Flexibilidade**: Permite definir horários diferentes para cada dia da semana, incluindo dias em que o negócio está fechado (simplesmente omitindo o dia do mapa).
* **Integração**: Persiste sua estrutura complexa de forma transparente em uma única coluna de banco de dados (como JSON), facilitando o armazenamento.
