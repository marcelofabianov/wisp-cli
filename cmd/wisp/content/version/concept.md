O `wisp.Version` é um *value object* que representa um número de versão, projetado especificamente para a implementação do padrão **Travamento Otimista** (Optimistic Locking).

Seu objetivo é previnir condições de corrida em ambientes concorrentes. Cada vez que uma entidade é atualizada, sua versão é incrementada. Ao tentar salvar, o sistema verifica se a versão da entidade em memória é a mesma que está no banco de dados. Se não for, significa que outro processo alterou o registro, e a operação atual é rejeitada, evitando a sobreposição de dados.

**Principais Funcionalidades:**

* **Ciclo de Vida Claro**: `InitialVersion()` retorna `1` (a versão de um registro recém-criado) e `Increment()` avança o número da versão.
* **Imutabilidade**: O método `Increment()` opera sobre um ponteiro, mas a intenção é que ele seja chamado dentro de métodos de comportamento da entidade (como `Audit.Touch()`), garantindo uma atualização controlada.
