O `wisp.Preferences` é um *value object* para armazenar de forma segura um conjunto de dados flexíveis no formato chave-valor, geralmente persistido como JSON.

Ele substitui o uso direto de `map[string]any` ou `json.RawMessage`, trazendo mais segurança e uma API mais conveniente.

**Principais Funcionalidades:**

* **Imutabilidade**: O método `Set(key, value)` não altera o objeto original, mas retorna uma **nova** instância com os dados atualizados, prevenindo efeitos colaterais.
* **Segurança de Tipo**: Garante que os dados subjacentes sejam sempre um objeto JSON válido.
* **API Conveniente**: Oferece métodos como `Get()` e `Set()` que abstraem a necessidade de serializar e desserializar o JSON manualmente para cada operação.
