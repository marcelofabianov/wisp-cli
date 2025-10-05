O `wisp.Money` é um value object fundamental para qualquer aplicação que lide com finanças. Ele foi projetado para resolver dois dos maiores problemas em cálculos monetários com computadores: a imprecisão de tipos de ponto flutuante e a mistura acidental de moedas.

**Principais Funcionalidades e Garantias:**

* **Precisão Absoluta**: `wisp.Money` **nunca usa `float64`**. Internamente, ele armazena o valor monetário como um `int64` em sua menor unidade (centavos), eliminando completamente os erros de arredondamento que podem ser catastróficos em cálculos financeiros.

* **Segurança de Moeda**: Um valor monetário é inseparável de sua moeda. `wisp.Money` força a associação de cada valor a um `wisp.Currency` (ex: BRL, USD). Operações aritméticas, como `Add` ou `Subtract`, falharão se você tentar combinar moedas diferentes, prevenindo erros graves de lógica de negócio.

* **Imutabilidade**: Todas as operações que alteram o valor (como `Add`, `Subtract`, `Multiply`) não modificam o objeto original, mas sim retornam uma **nova instância** de `wisp.Money`. Isso torna os cálculos previsíveis, seguros e livres de efeitos colaterais.

* **Aritmética de Domínio**: Oferece métodos de negócio ricos, como `Split(n)`, que divide um valor monetário em `n` partes de forma justa, distribuindo corretamente os centavos restantes — uma tarefa complexa que agora está encapsulada e testada.

Usar `wisp.Money` é a garantia de que todas as operações financeiras no seu sistema serão precisas, seguras e semanticamente corretas.
