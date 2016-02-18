# Tips and Tricks

## Envolva um http.HandlerFunc com closure (contexto léxico)
As vezes você quer passar dados na inicialização de um http.HandlerFunc. Isto pode ser facilmente feito criando um `closure` de `http.HandlerFunc`:

``` go
func MyHandler(database *sql.DB) http.Handler {
  return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
    // agora, você tem acesso ao *sql.DB aqui
  })
}
```

## Usando `gorilla/context` para requisição específica de dados
Com frequência precisamos armazenar e recuperar dados que são específicos da requisição HTTP atual. Use o `gorilla/context` para mapear valores e recuperá-los mais tarde. Ele contém um `mutex` (mutual exclusion, ou exclusão mútua na tradução livre) global do mapa de objetos requisitados.

``` go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    val := context.Get(r, "myKey")

    // retorna ("bar", true)
    val, ok := context.GetOk(r, "myKey")
    // ...

}
```
