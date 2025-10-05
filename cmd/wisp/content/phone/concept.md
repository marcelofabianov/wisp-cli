O `wisp.Phone` é um *value object* específico para validar e normalizar números de telefone do Brasil.

Ele encapsula as regras complexas do plano de numeração brasileiro, como a validação do DDD, a diferenciação entre celulares e fixos, e a regra do nono dígito.

**Principais Funcionalidades:**

* **Normalização**: Aceita diversos formatos de entrada (com `+55`, parênteses, hífens, espaços) e armazena o número internamente no formato canônico E.164 (ex: `5562988887777`).
* **Validação de Regras**: Verifica o comprimento correto para fixos (10 dígitos) e celulares (11 dígitos, incluindo o DDD) e a presença do nono dígito `9` para celulares.
* **API Expressiva**: Oferece métodos como `Formatted()`, `AreaCode()`, `IsMobile()` e `IsLandline()` para extrair informações e formatar a exibição.
