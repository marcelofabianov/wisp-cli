O `wisp.IPAddress` é um *value object* para representar endereços de rede IPv4 ou IPv6 de forma segura.

Ele utiliza o parser `net.ParseIP` da biblioteca padrão do Go, garantindo uma validação robusta e aderente aos padrões de rede. O tipo armazena internamente a representação canônica do endereço.

**Principais Funcionalidades:**

* **Validação Unificada**: `NewIPAddress` valida tanto o formato IPv4 quanto o IPv6.
* **Verificação de Versão**: Oferece métodos de conveniência como `IsV4()` e `IsV6()` para facilmente distinguir entre os tipos de endereço.
