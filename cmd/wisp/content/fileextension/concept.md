O `wisp.FileExtension` representa a extensão de um arquivo (ex: "pdf") de forma segura e padronizada. É um *value object* crucial para sistemas que lidam com upload e manipulação de arquivos.

**Principais Funcionalidades:**

* **Sistema de Registro**: A aplicação define uma "lista de permissões" de extensões válidas através da função `RegisterFileExtensions`.
* **Validação Segura**: `NewFileExtension` garante que apenas extensões previamente registradas sejam criadas, protegendo o sistema contra tipos de arquivo não desejados.
* **Normalização**: A entrada é normalizada para um formato canônico (minúsculas, sem o ponto inicial), aceitando entradas como ".JPG" ou "jpg" e tratando-as da mesma forma.
