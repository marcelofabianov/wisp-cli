O `wisp.Currency` é um *value object* que representa um código de moeda (padrão ISO 4217), como "BRL" ou "USD".

Seu principal objetivo é garantir que apenas moedas pré-aprovadas pela aplicação sejam utilizadas, evitando o uso de "magic strings" e erros de digitação.

**Principais Funcionalidades:**

* **Sistema de Registro**: Utiliza o padrão de registro. A aplicação deve primeiro registrar as moedas que suporta (ex: `wisp.RegisterCurrencies("BRL", "USD")`).
* **Validação Segura**: O construtor `NewCurrency` valida a entrada contra a lista de moedas registradas, rejeitando qualquer código desconhecido.
* **Normalização**: A entrada é automaticamente convertida para maiúsculas e tem seus espaços removidos, garantindo um formato canônico.
