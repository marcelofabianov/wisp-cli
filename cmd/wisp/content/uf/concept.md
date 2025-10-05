O `wisp.UF` é um *value object* do tipo enumeração para representar as Unidades Federativas do Brasil (ex: "GO", "SP", "RJ").

Ele utiliza o padrão de registro, exigindo que a aplicação defina uma lista de UFs permitidas na inicialização (`wisp.RegisterUFs`). Isso garante que apenas siglas de estados válidas e padronizadas sejam utilizadas em todo o sistema, prevenindo erros de digitação e inconsistências.
