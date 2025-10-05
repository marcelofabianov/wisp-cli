O `wisp.Timezone` é um *value object* que representa um fuso horário do padrão IANA (ex: "America/Sao_Paulo", "Europe/London").

Sua principal função é garantir que a aplicação lide com fusos horários de forma segura e explícita, eliminando ambiguidades em operações com datas e horas, especialmente em sistemas com usuários em diferentes localizações geográficas.

**Principais Funcionalidades e Garantias:**

* **Sistema de Registro**: Para aumentar a segurança e o controle do domínio, ele utiliza um sistema de registro. A aplicação deve primeiro registrar os fusos horários que suporta através da função `RegisterTimezones`.
* **Validação Segura**: O construtor `NewTimezone` valida a entrada contra a lista de fusos registrados, garantindo que apenas os fusos pré-aprovados pelo seu negócio sejam utilizados. A validação usa a função `time.LoadLocation` da biblioteca padrão, garantindo a conformidade com o banco de dados IANA do sistema.
* **Lógica de Conversão**: O método `Convert(time.Time)` oferece uma maneira simples e segura de converter um instante no tempo para o fuso horário específico representado pelo objeto.
