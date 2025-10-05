O `wisp.DateRange` é um *value object* que representa um período de tempo, composto por dois `wisp.Date`: um início e um fim.

Sua principal responsabilidade é garantir a integridade do intervalo, assegurando que a data de início nunca seja posterior à data de fim. Ele encapsula o conceito de "período", fornecendo métodos úteis como `Contains(date)` para verificar se uma data está dentro do intervalo.
