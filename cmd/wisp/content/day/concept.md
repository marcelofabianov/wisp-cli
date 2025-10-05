O `wisp.Day` é um *value object* que representa o dia de um mês, de 1 a 31.

É ideal para modelar regras de negócio que se repetem mensalmente em um dia específico, como datas de vencimento de faturas ou dias de fechamento de folha de pagamento. Ele valida que o valor esteja sempre no intervalo correto e oferece métodos de cálculo como `DaysUntil()` para saber quantos dias faltam para a próxima ocorrência.
