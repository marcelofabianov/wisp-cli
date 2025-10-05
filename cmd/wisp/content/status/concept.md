O `wisp.Status` é um *value object* genérico e extensível para representar um estado dentro de um conjunto finito de possibilidades (uma máquina de estados).

Assim como `Role` e `Unit`, ele utiliza o padrão de registro. Isso permite que diferentes domínios da sua aplicação (usuários, faturas, pedidos) definam seus próprios conjuntos de status válidos.

**Principais Funcionalidades:**

* **Sistema de Registro**: A aplicação define os status permitidos para um determinado contexto com a função `RegisterStatuses`.
* **Criação Segura**: `NewStatus` garante que apenas um status previamente registrado possa ser instanciado, prevenindo estados inválidos em suas entidades.
* **Flexibilidade**: Permite que você crie *type aliases* específicos no seu domínio (ex: `type UserStatus wisp.Status`) para adicionar comportamento e clareza semântica.
