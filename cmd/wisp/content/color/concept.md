O `wisp.Color` representa uma cor no formato RGBA. Sua principal função é validar e normalizar códigos de cores do padrão web hexadecimal.

Ele armazena a cor internamente como `color.RGBA` (da biblioteca padrão), o que facilita a manipulação e conversão dos dados.

**Principais Funcionalidades:**

* **Parsing de Hex**: `ParseColor` entende os formatos `#RRGGBB` (6 dígitos) e `#RGB` (3 dígitos).
* **Normalização**: Converte a entrada para minúsculas e armazena o valor de forma canônica.
* **API de Conversão**: Oferece métodos como `Hex()` e `RGBA()` para obter a cor em diferentes representações.
