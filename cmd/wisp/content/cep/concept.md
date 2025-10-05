O `wisp.CEP` representa um Código de Endereçamento Postal brasileiro de forma segura.

Sua principal responsabilidade é garantir que o CEP tenha um formato estruturalmente válido.

**Principais Funcionalidades:**

* **Validação de Formato**: Garante que o valor, após a limpeza, contenha exatamente 8 dígitos numéricos.
* **Normalização**: Aceita CEPs com ou sem o hífen (ex: `74000-000`) e armazena internamente apenas os 8 dígitos.
* **Formatação**: Oferece o método `Formatted()` para exibir o CEP com a máscara `XXXXX-XXX`.
