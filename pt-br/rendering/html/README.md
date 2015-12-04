# Templates HTML

Servir HTML é um trabalho importante para algumas aplicações web. Go tem uma
das minhas linguagens de template preferidas. Não pelos features, mas pela
simplicidade e a segurança. Renderizar templates HTML é quase tão fácil
quanto renderizar JSON usando o pacote `html/template` da biblioteca padrão.
Esse é o código de renderização de templates HTML:

``` go
package main

import (
	"html/template"
	"net/http"
	"path"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
```

Esse é o template que vamos usar. Ele deve ser colocado no arquivo
`templates/index.html` no diretório que o seu programa está sendo
executado:

``` html
<html>
  <h1>{{ .Title }}</h1>
  <h3>por {{ .Author }}</h3>
</html>
```

## Exercícios

1. Leia a documentação dos pacotes `text/template` e `html/template`. Brinque um pouco com a linguagem de template para ganhar uma noção dos objetivos dela, além dos seus pontos fortes e fracos.
2. No exemplo parseamos os arquivos em cada request, o que pode causar um impacto de performance. Experimente parsear os arquivos no começo do programa e executá-los no seu `http.Handler` (dica: use o método `Copy()` do `html.Template`).
3. Experimente parsear e usar múltiplos templates.
