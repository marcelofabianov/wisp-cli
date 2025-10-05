O `wisp.UUID` é um *value object* que encapsula um Identificador Único Universal, garantindo que os IDs em seu sistema sejam sempre válidos e gerados de forma consistente.

Este tipo atua como um wrapper em torno da biblioteca `github.com/google/uuid`, mas com um foco específico e opinativo.

**Principais Funcionalidades e Garantias:**

* **Geração Segura (UUIDv7)**: O construtor `NewUUID()` gera, por padrão, um **UUID v7**. Esta é uma versão moderna que combina a unicidade aleatória com um timestamp Unix, tornando os IDs sequenciais e muito mais eficientes como chaves primárias em bancos de dados (melhorando a localidade e o desempenho de índices).

* **Validação de Formato**: Ao criar um UUID a partir de uma string (`ParseUUID`), o tipo garante que a entrada esteja no formato canônico correto, rejeitando qualquer string inválida.

* **Segurança de Nulos**: O tipo possui um valor `wisp.Nil` explícito, e métodos como `IsNil()` fornecem uma forma clara e segura de verificar se um ID está vazio, evitando o uso acidental de IDs nulos.

* **Integração**: Assim como outros tipos `wisp`, ele implementa as interfaces padrão `json.Marshaler`/`Unmarshaler` e `sql.Valuer`/`Scanner` para uma integração transparente com APIs e bancos de dados.

Usar `wisp.UUID` em vez de `string` para seus IDs é uma prática que aumenta a robustez e a intenção do seu modelo de domínio.
