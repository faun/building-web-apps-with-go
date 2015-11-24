# Usando o pacote render

Para deixar a renderização de JSON e HTML mais simples ainda, existe o
pacote `github.com/unrolled/render`. Esse pacote foi inspirado no pacote
`martini-contrib/render` e é sempre minha primeira opção para renderização
nas minhas aplicações web.

``` go
package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Bem-vindo, visite as sub páginas agora."))
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Alguns dados em binário aqui."))
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		// Assume que existe um template em ./templates chamado "example.tmpl"
		// $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
		r.HTML(w, http.StatusOK, "example", nil)
	})

	http.ListenAndServe(":8080", mux)
}
```

## Exercícios
1. Divirta-se com todas as opções disponíveis na hora de chamar `render.New()`.
2. Tente usar a função helper `.yield` (usando as chaves) e um layout com templates HTML.
