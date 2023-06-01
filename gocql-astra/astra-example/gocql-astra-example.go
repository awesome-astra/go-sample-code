package main

import (
	"fmt"
	"os"
	"time"
	"log"

	gocqlastra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     gocql.UUID
	Title  string
	Author string
	Year   int
}

func main() {

	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	var cluster *gocql.ClusterConfig

	if len(os.Getenv("ASTRA_DB_APPLICATION_TOKEN")) > 0 {

		if len(os.Getenv("ASTRA_DB_ID")) == 0 {
			panic("database ID is required when using a token")
		}
	}

	fmt.Println("Creating the cluster now")

	cluster, err = gocqlastra.NewClusterFromURL("https://api.astra.datastax.com", os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), 20*time.Second)
		if err != nil {
			fmt.Errorf("unable to load cluster %s from astra: %v", os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), err)
		}
		cluster.Timeout = 30*time.Second
		session, err := gocql.NewSession(*cluster)

		if err != nil {
			log.Fatalf("unable to connect session: %v", err)
		}

		defer session.Close()

	// Delete the table
	if err := session.Query(`DROP TABLE IF EXISTS library.books`).Exec(); 
	err != nil {
		fmt.Println("Dropping table to clean for examples.")
	}

	// Create the table
	fmt.Println("Creating the table now")
	if err := session.Query(`CREATE TABLE library.books ( id uuid PRIMARY KEY, title text, author text, year int );`).Exec(); 
	err != nil {
		log.Fatal(err)
	}
	// Create Rows
	newBook := Book{
		ID:     gocql.TimeUUID(),
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2023,
	}

	secondBook := Book{
		ID:     gocql.TimeUUID(),
		Title:  "Zen and the Art of Go",
		Author: "Jane Doe",
		Year:   2023,
	}

	fmt.Println("Adding the books now:")
	
	if err := session.Query(`INSERT INTO library.books (id, title, author, year) VALUES (?, ?, ?, ?)`,
		newBook.ID, newBook.Title, newBook.Author, newBook.Year).Exec(); err != nil {
		log.Fatal(err)
	}

	if err := session.Query(`INSERT INTO library.books (id, title, author, year) VALUES (?, ?, ?, ?)`,
	secondBook.ID, secondBook.Title, secondBook.Author, secondBook.Year).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Querying the books now:")
	iter := session.Query(`SELECT id, title, author, year FROM library.books`).Iter()
	
	var books []Book
	var book Book
	for iter.Scan(&book.ID, &book.Title, &book.Author, &book.Year) {
		books = append(books, book)
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	// Print all books
	for _, book := range books {
		fmt.Println("Book:", book)
	}

	// Update
	updatedBook := Book{
		ID:     newBook.ID,
		Title:  "Advanced Go Programming",
		Author: "John Doe",
		Year:   2024,
	}

	if err := session.Query(`UPDATE library.books SET title = ?, author = ?, year = ? WHERE id = ?`,
		updatedBook.Title, updatedBook.Author, updatedBook.Year, updatedBook.ID).Exec(); err != nil {
		log.Fatal(err)
	}

	// Delete
	if err := session.Query(`DELETE FROM library.books WHERE id = ?`, newBook.ID).Exec(); err != nil {
		log.Fatal(err)
	}


}


