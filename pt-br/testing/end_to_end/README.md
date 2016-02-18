# Testes fim-a-fim

Testes fim-a-fim nos permitem testar a aplicação ao longo do ciclo inteiro
de uma requisição. Enquanto os testes de unidade são feitos para testar
somente uma função em particular, os testes fim-a-fim rodam as camadas intermediárias,
o roteamento e todo tipo de código através do qual uma requisição passa.

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func HelloWorld(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprint(res, "Hello World")
}

func App() http.Handler {
	n := negroni.Classic()

	m := func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		fmt.Fprint(res, "Antes...")
		next(res, req)
		fmt.Fprint(res, "...Depois")
	}
	n.Use(negroni.HandlerFunc(m))

	r := httprouter.New()

	r.GET("/", HelloWorld)
	n.UseHandler(r)
	return n
}

func main() {
	http.ListenAndServe(":3000", App())
}
```

Esse é o arquivo de teste. Ele deve ser colocado no mesmo diretório 
da sua aplicação com o nome `main_test.go`.

```go
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	exp := "Antes...Hello World...Depois"

	if exp != string(body) {
		t.Fatalf("Esperado %s recebido %s", exp, body)
	}
}
```

# Exercícios
1. Crie outra camada de middleware que modifique o código HTTP (status) da requisição.
2. Crie uma requisição POST e teste que ela é tratada corretamente.
