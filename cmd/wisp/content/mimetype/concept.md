O `wisp.MIMEType` é um *value object* de segurança que representa um tipo de mídia (ex: "image/jpeg").

Assim como `FileExtension`, ele utiliza um sistema de registro (`RegisterMIMETypes`) para criar uma "lista de permissões" de tipos de conteúdo que a aplicação pode processar. Além disso, ele valida que a entrada siga o formato padrão `tipo/subtipo`. É uma ferramenta essencial para proteger rotinas de upload de arquivos.
