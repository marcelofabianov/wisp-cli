O `wisp.AuditUser` é um *value object* que representa o ator responsável por uma mudança em uma entidade (quem criou, quem atualizou, etc.).

Ele é projetado para ser flexível, suportando dois tipos de identificadores:

1.  **E-mail**: Para ações realizadas por um usuário humano, o valor deve ser um e-mail válido. O tipo reutiliza `wisp.Email` para garantir a validação.
2.  **Sistema**: Para ações automatizadas ou internas, a string literal `"system"` pode ser usada.

Essa dualidade, gerenciada pelo construtor `NewAuditUser`, garante que o campo de auditoria seja sempre válido e semanticamente claro, com métodos auxiliares como `IsSystem()` e `IsEmail()` para verificar o tipo de ator.
