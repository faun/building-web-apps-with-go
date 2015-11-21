# Roteamento de URL

Para algumas aplicações simples, o pacote padrão `http.ServeMux` pode te levar muito longe. Se você precisar de mais poder na forma como você analisa os parâmetros de URL e os encaminha para
o handler apropriado, você pode precisar usar um pacote de roteamento de terceiros.
Neste tutorial, nós usaremos o popular pacote
`github.com/julienschmidt/httprouter` como nosso roteador.
`github.com/julienschmidt/httprouter` é uma ótima opção para um roteador, pois é uma implementação muito simples com uma dos melhores benckmarks de performance de todos os roteadores Go de terceiros.

Neste exemplo, vamos criar alguns roteamentos para um recurso RESTful chamado "posts". A seguir, definimos mecanismos para ver o index, mostrar, criar, atualizar, deletar e editar os posts.

``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", HomeHandler)

	// Coleção de Posts
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	// Posts singulares
	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Iniciando serviço em :8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "index de posts")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "criando posts")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "mostrando post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "atualizando post")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "deletando post")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "editando post")
}
```

## Exercícios

1. Explore a documentação do pacote `github.com/julienschmidt/httprouter`.
2. Saiba como o `github.com/julienschmidt/httprouter` roda bem com `http.Handler`s existentes como o `http.FileServer`
3. `httprouter` tem uma interface muito simples. Explore qual tipo de abstração pode ser criado em cima deste rápido roteador para construir coisas como roteamento RESTful mais fácil.
