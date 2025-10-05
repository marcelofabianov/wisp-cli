O `wisp.PositiveInt` é um *value object* "primitivo seguro" que garante que um valor inteiro seja sempre estritamente maior que zero.

Sua validação no construtor `NewPositiveInt` rejeita qualquer valor menor ou igual a zero. É ideal para representar conceitos de negócio que, por definição, não podem ser nulos ou negativos, como o limite de vagas em um curso, a quantidade de itens em um pedido, ou o número de parcelas de um pagamento.

Usar `PositiveInt` em vez de `int` torna o modelo de domínio mais explícito e seguro, prevenindo a criação de entidades em estados inválidos.
