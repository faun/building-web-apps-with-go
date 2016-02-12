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
