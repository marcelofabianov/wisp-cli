O `wisp.DayOfWeek` representa um dia da semana (Domingo a Sábado) de forma segura e tipada.

Construído sobre o `time.Weekday` da biblioteca padrão, ele enriquece a API com constantes (`wisp.Monday`), métodos de conveniência (`IsWeekend()`) e um parser de strings. A serialização JSON é customizada para usar strings legíveis (ex: "monday"), melhorando a clareza em APIs.
