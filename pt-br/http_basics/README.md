# O pacote net/http
Você provavelmente já ouviu falar que Go é fantástico para a construção de
aplicações web de todas as formas e tamanhos.
Isto é em parte devido ao fantástico trabalho que tem sido
colocado em fazer a biblioteca padrão limpa, consistente e fácil de usar.

Talvez um dos pacotes mais importantes para qualquer desenvolvedor web Go novato
é o pacote `net/http`. Este pacote permite que você construa servidores
HTTP em Go com as suas poderosas construções composicionais.
Antes de começar a codificação, vamos fazer uma visão geral extremamente
rápida sobre HTTP.

## HTTP Básico
Quando nós falamos sobre aplicações web, geralmente significa que estamos
construindo servidores HTTP. HTTP é um protocolo que foi originalmente concebido
para transportar documentos HTML de um servidor para um cliente navegador web.
Hoje em dia, HTTP é usado para transportar muito mais do que HTML.

![](../../en//http_basics/http_diagram.png)

A coisa importante a notar neste diagrama é os dois pontos de interação
entre a *Server* e *Navegador*. O *Navegador* faz uma solicitação HTTP
com algumas informações, o *Server*, em seguida, processa esse pedido e
retorna uma *Resposta*.

Este padrão de requisição-resposta é um dos pontos chave focais na contrução de
aplicações web em Go. Na verdade, a peça mais importante do pacote `net/http`
é a Interface `http.Handler`.  

## A Interface http.Handler

Quando você se tornar mais familiar com Go, você irá notar o quanto *interfaces*
criam impacto no design de seus programas.
A interface `net/http` encapsula o padrão de requisição-resposta em um método:

``` go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
Implementadores desta interface são esperados para inspecionar e processar dados
vindos do objeto `http.Request` e escrever uma resposta para o objeto
`http.ResponseWriter`.

A interface `http.ResponseWriter` parece com isso:

``` go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(int)
}
```

## Compondo Web Services
Como grande parte do pacote `net/http` é construído fora de tipos de uma
interface bem definida nós podemos (e devemos) construir nossa aplicação web
com esta composição em mente.
Cada implementação `http.Handler` pode ser pensada como seu próprio servidor web.

Muitos padrões podem ser achados nesta simples mas poderosa suposição.
Ao longo deste livro, vamos cobrir alguns desses padrões e como podemos usá-los
para resolver problemas do mundo real.


## Exercício: 1 Line File Server
Vamos resolver um problema do mundo real em uma linha de código.

Há muito tempo pessoas precisavam servir apenas arquivos estáticos.
Pode ser que você tenha uma landing page estática em HTML e quer apenas servir
HTML, imagens, e CSS e chamar isto um dia. Verdade, você poderia colocar no
Apache ou no `SimpleHTTPServer` do Python, mas Apache é muito para este pequeno
site e `SimpleHTTPServer` é, assim, muito lento.

Nós iremos iniciar criando um novo projeto em nosso `GOPATH`.

``` bash
cd GOPATH/src
mkdir fileserver && cd fileserver
```

Crie um **main.go** com nosso típico padrão go.

``` go
package main

import "net/http"

func main() {
}
```

Todos nós precisamos importar o pacote `net/http` para este trabalho.
Lembre-se que tudo isto faz parte da biblioteca padrão em Go

Vamos escrever nosso código fileserver:

``` go
http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
```

A função `http.ListenAndServe` é usada para iniciar o servidor, ele se ligará
ao endereço que nós demos (`:8080`) e quando este receber uma requisição HTTP,
ele vai entregá-lo ao `http.Handler` que nós fornecemos como segundo argumento.
Em nosso caso é o `http.FilseServer`.

A função `http.FilseServer` constrói uma `http.Handler` que irá servir um
diretório inteiro de arquivos e descobrir qual arquivo para saque com base no
caminho da solicitação.
Nós dissemos para o FileServer servir o diretório de trabalho atual com
`http.Dir(".")`.

O programa completo fica assim:

``` go
package main

import "net/http"

func main() {
    http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
```

Vamos construir e rodar nosso programa fileserver:
``` bash
go build
./fileserver
```

Se visitar `localhost:8080/main.go` devemos ver o conteúdo do nosso arquivo
**main.go** em nosso navegador web. Podemos executar este programa a partir
de qualquer diretório e servir a árvore como um servidor de arquivo estático.
Tudo em uma linha de código Go.
