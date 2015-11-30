# Banco de Dados

Uma das perguntas que eu mais recebo sobre desenvolvimento web em Go é como
conectar em um banco de dados SQL. Felizmente, Go tem um pacote de SQL fantástico
na biblioteca padrão que nos permite usar toda uma gama de drivers de diferentes
bancos de dados SQL. Nesse exemplo vamos conectar com um banco de dados SQLite,
mas a sintaxe (exceto alguns pequenos detalhes semânticos do SQL) é a mesma
para um banco MySQL or PostgreSQL.

``` go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := NewDB()
	log.Println("Ouvindo em :8080")
	http.ListenAndServe(":8080", ShowBooks(db))
}

func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var title, author string
		err := db.QueryRow("select title, author from books").Scan(&title, &author)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(rw, "O primeiro livro é '%s' by '%s'", title, author)
	})
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists books(title text, author text)")
	if err != nil {
		panic(err)
	}

	return db
}
```

## Exercícios
1. Use a função `Query` na sua instância do `sql.DB` para extrair uma coleção de linhas e mapeie elas a structs.
2. Adicione a habilidade de inserir novos registros no nosso banco de dados usando um form HTML.
3. Rode `go get github.com/jmoiron/sqlx` e observe as melhorias feitas em cima do pacote `database/sql` da biblioteca padrão.
