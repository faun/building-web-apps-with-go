# JSON

JSON vem se tornando rapidamente o formato de serialização dominante para APIs
web, o que faz dele o mais relevante quando se está aprendendo a construir aplicações
web usando Go. Felizmente, é muito simples trabalhar com JSON em Go - é extremamente
fácil transformar structs Go existentes em JSON usando o pacote `encoding/json`
da biblioteca padrão.

``` go
package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
```

## Exercícios
1. Leia a documentação da API JSON e descubra como renomear e ignorar campos na serialização para JSON.
2. Em vez de usar o método `json.Marshal`, tente usar a API `json.Encoder`.
3. Descubra como fazer pretty print do JSON com o pacote `encoding/json`.
