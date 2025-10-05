O `wisp.Latitude` é um *value object* que representa a coordenada geográfica de latitude, a posição Norte-Sul na superfície da Terra.

Sua principal responsabilidade é garantir a integridade do dado, validando que o valor esteja sempre no intervalo globalmente aceito de **-90.0 (Polo Sul) a 90.0 (Polo Norte)**.

O uso de um tipo específico para Latitude, em vez de um `float64` genérico, permite encapsular lógicas de negócio que dependem exclusivamente desta coordenada. Por exemplo, determinar se um ponto está no Hemisfério Norte ou Sul, ou classificar uma localização em zonas climáticas (tropical, temperada, polar). Isso torna o código mais seguro e expressivo.
