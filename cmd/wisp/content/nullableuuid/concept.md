O `wisp.NullableUUID` é um *value object* projetado para representar um `wisp.UUID` que pode ser opcional, ou seja, nulo. É a solução ideal para modelar chaves estrangeiras (FKs) que não são obrigatórias.

Enquanto `wisp.UUID` pode usar `wisp.Nil` para representar um valor nulo, isso pode gerar ambiguidade. O `NullableUUID` resolve isso encapsulando um `UUID` e um booleano `Valid`.

**Principais Vantagens:**

* **Clareza de Intenção**: Um campo do tipo `wisp.NullableUUID` em uma `struct` documenta explicitamente que aquele relacionamento é opcional.
* **Segurança**: Força o desenvolvedor a verificar o estado `Valid` antes de usar o UUID, prevenindo o uso acidental de um ID nulo e causando erros em tempo de execução.
* **Consistência**: Alinha-se com outros tipos nuláveis do `wisp`, como o `NullableTime`, criando uma API coesa e previsível.
