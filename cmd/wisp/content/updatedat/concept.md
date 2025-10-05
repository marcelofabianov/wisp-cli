O `wisp.UpdatedAt`, similar ao `CreatedAt`, representa o timestamp da última modificação de uma entidade.

Sua principal vantagem é a **clareza de intenção** através do método `Touch()`. Em vez de espalhar `meuObjeto.Campo = time.Now()` por todo o código, você chama `meuObjeto.Campo.Touch()`. Isso centraliza a lógica de atualização do timestamp e torna a ação de "marcar como atualizado" explícita e legível.

Usar `UpdatedAt` promove um código mais limpo e expressivo para o rastreamento de modificações.
