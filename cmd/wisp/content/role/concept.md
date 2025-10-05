O `wisp.Role` é um *value object* extensível para representar papéis de usuário ou níveis de permissão (ex: "ADMIN", "STUDENT").

Ele utiliza o padrão de registro para garantir que apenas papéis pré-definidos pela sua aplicação possam ser utilizados, trazendo segurança de tipos para seu sistema de autorização e evitando o uso de "magic strings".

**Principais Funcionalidades:**

* **Sistema de Registro**: A aplicação define uma "lista de permissões" de papéis válidos através da função `RegisterRoles`.
* **Validação Segura**: O construtor `NewRole` valida a entrada contra a lista de papéis registrados, rejeitando qualquer valor desconhecido e garantindo a consistência do domínio.
* **Normalização**: A entrada é padronizada para maiúsculas e sem espaços.
