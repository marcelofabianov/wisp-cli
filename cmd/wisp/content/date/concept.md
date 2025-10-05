O `wisp.Date` é um *value object* para representar uma data de calendário (Ano, Mês, Dia) de forma pura, sem a complexidade de horas, minutos ou fusos horários.

Ele resolve uma classe comum de bugs em que o `time.Time` padrão é usado para conceitos de data, o que pode levar a erros de comparação e de fuso horário. O `wisp.Date` garante que o valor seja sempre um dia de calendário, armazenando-o internamente como um `time.Time` com a hora zerada e em UTC.
