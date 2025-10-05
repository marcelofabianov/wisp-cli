O `wisp.CNPJ` é um *value object* que representa um Cadastro Nacional da Pessoa Jurídica (CNPJ) de forma segura e canônica.

Ele garante a integridade de um CNPJ no sistema, tratando-o como um conceito de negócio, não como uma simples `string`.

**Principais Funcionalidades e Garantias:**

* **Validação Matemática**: Executa o cálculo completo dos **dígitos verificadores**, garantindo que o número seja matematicamente válido de acordo com as regras da Receita Federal.
* **Normalização Automática**: Aceita CNPJs com ou sem máscara (`XX.XXX.XXX/XXXX-XX`) e armazena o valor internamente com apenas os 14 dígitos.
* **Segurança de Formato**: Rejeita entradas com tamanho incorreto ou sequências inválidas (como dígitos repetidos).
* **API Expressiva**: Oferece o método `Formatted()` para exibir o CNPJ com a máscara padrão, facilitando a apresentação em interfaces.
