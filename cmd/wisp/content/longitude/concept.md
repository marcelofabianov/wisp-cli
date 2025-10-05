O `wisp.Longitude` é um *value object* que representa a coordenada geográfica de longitude, a posição Leste-Oeste na superfície da Terra.

Sua função primária é a validação estrita do valor, garantindo que ele esteja sempre dentro do intervalo globalmente aceito de **-180.0 a 180.0**.

Ao isolar a longitude em seu próprio tipo, o sistema pode encapsular lógicas de negócio específicas, como determinar se uma localização está no Hemisfério Ocidental ou Oriental, ou agrupar dados por proximidade ao Meridiano de Greenwich. Isso previne dados de geolocalização inválidos e torna o modelo de domínio mais preciso.
