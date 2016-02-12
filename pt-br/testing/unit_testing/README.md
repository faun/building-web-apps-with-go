# Testes unitários

Testes unitários nos permitem testar uma função do tipo `http.HandlerFunc`
diretamente sem rodar camadas intermediárias, roteamento ou qualquer outro
tipo de código em torno da função.

```go
package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":3000", nil)
}
```

Esse é o arquivo de teste. Ele deve ser colocado no mesmo diretório
da sua aplicação com o nome `main_test.go`.

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloWorld(res, req)

	exp := "Hello World"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Esperado %s recebido %s", exp, act)
	}
}
```

## Exercícios
1. Modifique a saída do `HelloWorld` para imprimir um parâmetro e teste que esse parâmetro foi renderizado.
2. Crie uma requisição POST e teste que ela é tratada corretamente.
