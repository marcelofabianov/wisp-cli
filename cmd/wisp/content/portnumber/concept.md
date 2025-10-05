O `wisp.PortNumber` representa um número de porta de rede (TCP/UDP).

Sua principal função é garantir que o valor da porta esteja sempre dentro do intervalo padrão definido pela IANA, que vai de **1 a 65535**. Isso previne o uso de portas inválidas (como 0 ou valores negativos) em configurações de serviços e comunicação de rede, tornando a aplicação mais segura e estável.
