# Controllers

Controllers são um tópico familiar em outras comunidades de desenvolvimento web.
Como a maioria dos desenvolvedores web se apoia na excelente biblioteca net/http,
poucas implementações de controllers ganharam popularidade. Apesar disso, existem
muitos benefícios em usar controllers. Eles permitem abstrações limpas e bem
definidas que vão além do que a interface handler do net/http oferece sozinha.

## Dependências

Nesse exemplo vamos experimentar construir nossa própria implementação de um
controller usando algumas features comuns do Go. Primeiro, vamos começar com o
problema que estamos tentando resolver. Digamos que estamos usando a biblioteca
`render` que discutimos em capítulos anteriores:

``` go
var Render = render.New(render.Options{})
```

Se quisermos que os nossos `http.Handler`s acessem a nossa instância do
`render.Render`, temos algumas opções.

**1. Usar uma variável global:** Isso não é tão ruim para programas pequenos,
mas quando o programa cresce fica um pesadelo para dar manutenção.

**2. Passar a variável através de um closure para o http.Handler:** Essa é 
uma ótima ideia e deveríamos usá-la a maior parte do tempo. A implementação
acaba ficando assim:

``` go
func MyHandler(r *render.Render) http.Handler {
  return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    // agora conseguimos acessar a variável r
  })
}
```

## Em defesa dos controllers

Quando o tamanho do programa aumenta, fica claro que a maior parte dos
`http.Handler`s compartilha as mesmas dependências e existirão vários desses
closures com os mesmos argumentos. A minha maneira preferida de limpar isso
é criar uma pequena implementação base de um controller que me dá algumas
vantagens:

1. Me permite compartilhar as dependências entre vários `http.Handler`s com objetivos ou conceitos parecidos.
2. Evita variáveis globais e funções para testes e mocks fáceis.
3. Me dá um mecanismo centralizado e idiomático para tratar erros.

A melhor parte dos controllers é que me dão tudo isso sem importar um
pacote externo! A maior parte dessa funcionalidade vem do uso de features
do Go, principalmente structs e embedding. Vamos dar uma olhada na implementação.

``` go
package main

import "net/http"

// Action define uma assinatura de função padronizada para usarmos
// quando criamos ações no controller. Uma ação no controller é
// basicamente só um método do controller.
type Action func(rw http.ResponseWriter, r *http.Request) error

// Esse é o nosso controller base.
type AppController struct{}

// A função Action facilita o tratamento de erros em um controller.
func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}
```

E é isso! Essa é a implementação completa que precisamos para ter as
vantagens de um controller ao nosso alcance. Tudo que falta é implementar
um controller de exemplo:

``` go
package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

type MyController struct {
	AppController
	*render.Render
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}
	http.ListenAndServe(":8080", c.Action(c.Index))
}
```

## Exercícios
1. Aumente o `MyController` para ter múltiplas ações com rotas diferentes.
2. Faça experiências com mais implementações de controllers, seja criativo.
3. Reescreva o método `Action` do `MyController` para renderizar uma página HTML de erro.
