O `wisp.Email` é um *value object* que garante que uma string representa um endereço de e-mail estruturalmente válido.

Em vez de usar expressões regulares (regex), que são notoriamente complexas e falhas para validar e-mails, este tipo utiliza o parser `net/mail.ParseAddress` da biblioteca padrão do Go, que segue a especificação da RFC 5322.

**Principais Funcionalidades:**

* **Validação Robusta**: Garante que o formato do e-mail seja válido.
* **Normalização**: Converte o e-mail para letras minúsculas e remove espaços em branco no início e no fim, garantindo um formato de armazenamento canônico.
