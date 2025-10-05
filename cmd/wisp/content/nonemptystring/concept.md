O `wisp.NonEmptyString` é um *value object* "primitivo seguro" que garante que uma string nunca seja vazia.

Sua principal função é a validação no momento da criação (`NewNonEmptyString`), que apara os espaços em branco no início e no fim da string de entrada e retorna um erro se o resultado for vazio.

É a ferramenta ideal para campos de domínio que são obrigatórios, como nomes, títulos ou descrições, garantindo que " " (apenas espaços) não seja considerado um valor válido e prevenindo a propagação de dados vazios pelo sistema.
