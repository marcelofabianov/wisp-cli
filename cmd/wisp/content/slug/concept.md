O `wisp.Slug` é um *value object* que representa uma string otimizada para ser um identificador amigável e único em URLs.

Ele encapsula um pipeline de normalização robusto para transformar qualquer string em um formato seguro para a web, melhorando a usabilidade e o SEO de links.

**Principais Funcionalidades:**

* **Transliteração**: Remove acentos e caracteres especiais (ex: `Geração` -> `Geracao`).
* **Conversão para Minúsculas**: Padroniza todo o texto para lowercase.
* **Substituição Inteligente**: Substitui símbolos comuns por palavras (`%` -> `percent`) e outros caracteres inválidos por hífens.
* **Limpeza**: Garante que não haja hífens duplicados ou no início/fim da string.
* **Validação**: Assegura que o resultado final, após toda a normalização, não seja uma string vazia.
