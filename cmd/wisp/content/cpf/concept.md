O `wisp.CPF` é um value object projetado para representar um Cadastro de Pessoas Físicas (CPF) de forma segura e canônica. Ele elimina os riscos de se usar uma `string` primitiva para um conceito de negócio tão importante.

A principal responsabilidade deste tipo é garantir a **integridade absoluta** de um CPF dentro do seu sistema.

**Principais Funcionalidades e Garantias:**

* **Validação Matemática**: A validação não se limita ao formato. O `wisp.CPF` executa o cálculo completo dos **dígitos verificadores**, garantindo que o número seja matematicamente válido de acordo com as regras da Receita Federal.

* **Normalização Automática**: Ele aceita CPFs em diversos formatos na entrada (ex: `123.456.789-00` ou `12345678900`) e normaliza o valor internamente para armazenar apenas os 11 dígitos, garantindo um formato de dados consistente.

* **Segurança de Formato**: Rejeita entradas com tamanho incorreto ou sequências inválidas conhecidas (como `111.111.111-11`).

* **API Expressiva**: Oferece métodos de conveniência como `Formatted()`, que retorna o CPF com a máscara `XXX.XXX.XXX-XX`, facilitando a exibição para o usuário final.

Ao usar `wisp.CPF`, você move toda a complexidade de validação e formatação para um único lugar, tornando o resto da sua aplicação mais limpa e segura.
