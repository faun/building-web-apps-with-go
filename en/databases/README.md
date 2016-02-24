# Databases
One of the most asked questions I get about web development in Go is how to
connect to a SQL database. Thankfully, Go has a fantastic SQL package in the
standard library that allows us to use a whole slew of drivers for different
SQL databases. In this example we will connect to a SQLite database, but the
syntax (minus some small SQL semantics) is the same for a MySQL or PostgreSQL
database.

Below is an example web application that has two endpoints: `POST /books` to
create a book and `GET /books` to show all books that have been created.

To create a new book, run:

    curl -vX POST http://localhost:8080/books \
    --header "Content-Type: application/json" \
    -d "{\"author\": \"George Orwell\", \"title\": \"1984\"}"

To create another:

    curl -vX POST http://localhost:8080/books \
    --header "Content-Type: application/json" \
    -d "{\"author\": \"Victor Hugo\", \"title\": \"Les Misérables\"}"

To view all books:

    curl -v http://localhost:8080/books --header "Content-Type: application/json"

``` go
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type errResponse struct {
	Message string `json:"message"`
}

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatalln("Could not connect to database")
	}

	r := httprouter.New()
	r.GET("/books", getHandler(db))
	r.POST("/books", postHandler(db))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)
}

func getHandler(db *sql.DB) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		books, err := getBooks(db)
		if err != nil {
			respondError(rw, err)
			return
		}

		rw.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(rw).Encode(books); err != nil {
			respondError(rw, err)
			return
		}
	}
}

func postHandler(db *sql.DB) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var b book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			respondError(rw, err)
			return
		}

		if err := createBook(db, b); err != nil {
			respondError(rw, err)
			return
		}

		rw.WriteHeader(http.StatusNoContent)
	}
}

func newDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := "CREATE TABLE IF NOT EXISTS books(title TEXT, author TEXT)"
	if _, err := db.Exec(q); err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

func getBooks(db *sql.DB) ([]book, error) {
	q := "SELECT title, author FROM books"
	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var books []book
	for rows.Next() {
		var b book
		if err := rows.Scan(&b.Title, &b.Author); err != nil {
			log.Println(err)
			return nil, err
		}

		books = append(books, b)
	}

	return books, nil
}

func createBook(db *sql.DB, b book) error {
	q := "INSERT INTO books(title, author) VALUES ($1, $2)"
	if _, err := db.Exec(q, b.Title, b.Author); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func respondError(rw http.ResponseWriter, err error) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)

	er := errResponse{
		Message: err.Error(),
	}

	if err := json.NewEncoder(rw).Encode(er); err != nil {
		log.Println(err)
	}
}
```

## Exercises
1. Make use of the `Query` function on our `sql.DB` instance to extract a collection of rows and map them to structs.
2. Add the ability to insert new records into our database by using an HTML form.
3. `go get github.com/jmoiron/sqlx` and observe the improvements made over the existing database/sql package in the standard library.
