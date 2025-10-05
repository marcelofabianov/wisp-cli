O `wisp.CreatedAt` é um *value object* simples que representa o timestamp de criação de uma entidade.

Embora seja um wrapper fino em torno do `time.Time`, seu propósito é adicionar **clareza semântica** ao modelo de domínio. Ao usar `CreatedAt` em vez de `time.Time` em uma `struct`, você deixa explícito que aquele campo representa o momento da criação e, conceitualmente, não deve ser alterado após a inicialização.

Ele é um componente essencial da `struct wisp.Audit`.
