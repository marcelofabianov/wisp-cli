O `wisp.Audit` é um *value object* composto, projetado para ser embutido (embedded) em entidades de domínio. Ele agrupa todos os campos e lógicas necessários para uma trilha de auditoria completa.

Ao compor outros tipos `wisp` (`CreatedAt`, `AuditUser`, `Version`, `NullableTime`, etc.), ele fornece uma solução "tudo em um" para rastrear o ciclo de vida de um registro.

**Principais Funcionalidades:**

* **Ciclo de Vida Completo**: Rastreia quem criou, quem atualizou, quando foi criado, quando foi atualizado, e a versão para travamento otimista.
* **Soft Delete e Arquivamento**: Suporta os padrões de `deleted_at` e `archived_at` através do `NullableTime`.
* **API de Comportamento**: Oferece métodos como `Touch()`, `Delete()` e `Archive()` que encapsulam a lógica de atualização dos campos de auditoria, garantindo consistência.
* **Métodos de Estado**: Funções como `IsActive()`, `IsDeleted()` e `IsArchived()` fornecem uma maneira clara de verificar o estado do registro.
