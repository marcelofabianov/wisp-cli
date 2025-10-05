O `wisp.BirthDate` é um *value object* especializado que representa uma data de nascimento.

Construído sobre o `wisp.Date`, ele adiciona uma regra de negócio crucial: a data de nascimento não pode ser no futuro. Além disso, ele encapsula lógicas de negócio complexas, como o cálculo de idade (`Age()`) e a verificação de maioridade (`IsOfAge()`), tornando o código de domínio mais claro e seguro.
