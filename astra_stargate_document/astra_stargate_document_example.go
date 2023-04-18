package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/awesome-astra/astra_stargate_go"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if len(os.Getenv("ASTRA_DB_APPLICATION_TOKEN")) == 0 {
		fmt.Println(fmt.Errorf("please set your environment variables or use astra db create-dotenv to create a .env file"))
		return
	}

	client := astra_stargate_go.NewBasicAuthClient(os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_REGION"))
	if err != nil {
		fmt.Println(err)
	}

	// Document: Create a collection
	fmt.Println("Create 'library' collection in the library keyspace")
	jsonStr := []byte(`{"name":"library"}`)
	responsebody, err := client.APIPost("/api/rest/v2/namespaces/library/collections", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// Document: Retrieve collections
	fmt.Println("Retrieve collections")
	responsebody, err = client.APIGet("/api/rest/v2/namespaces/library/collections")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// Document: Create a document without specifying the ID
	fmt.Println("Create a document without specifying an ID")
	jsonStr = []byte(`{
		"stuff": "Random ramblings",
		"other": "I do not care much about the ID for this document."
	  }`)
	responsebody, err = client.APIPost("/api/rest/v2/namespaces/library/collections/library", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// Document: Create a document with a specific ID (long-ID-number)
	fmt.Println("Create a document specifying an ID")
	jsonStr = []byte(`{
		"stuff": "long-ID-number",
		"other": "I need a document with a set value for a test."
	  }`)
	responsebody, err = client.APIPut("/api/rest/v2/namespaces/library/collections/library/long-ID-number", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// Document: Create a document with some data
	fmt.Println("Create a document specifying an ID")
	jsonStr = []byte(`{
		"reader": {
		   "name": "Amy Smith",
		   "user_id": "12345",
		   "birthdate": "10-01-1980",
		   "email": {
			   "primary": "asmith@gmail.com",
			   "secondary": "amyispolite@aol.com"
		   },
		   "address": {
			   "primary": {
				   "street": "200 Antigone St",
				   "city": "Nevertown",
				   "state": "MA",
				   "zip-code": 55555
			   },
			   "secondary": {
				   "street": "850 2nd St",
				   "city": "Evertown",
				   "state": "MA",
				   "zip-code": 55556
			   }
		   },
		   "reviews": [
			   {
				   "book-title": "Moby Dick", 
				   "rating": 4, 
				   "review-date": "04-25-2002",
				   "comment": "It was better than I thought."
			   },
			   {
				   "book-title": "Pride and Prejudice", 
				   "rating": 2, 
				   "review-date": "12-02-2002",
				   "comment": "It was just like the movie."
			   }
		   ]
		}
	}`)

	responsebody, err = client.APIPost("/api/rest/v2/namespaces/library/collections/library", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// Document: Create a document with some data
	fmt.Println("Create a book document with data specifying an ID")
	jsonStr = []byte(`{
        "book": {
            "title": "Native Son",
            "isbn": "12322",
            "author": [
                "Richard Wright"
            ],
            "pub-year": 1930,
            "genre": [
                "poverty",
                "action"
            ],
            "format": [
                "hardback",
                "paperback",
                "epub"
            ],
            "languages": [
                "English",
                "German",
                "French"
            ]
        }
    }`)

	responsebody, err = client.APIPut("/api/rest/v2/namespaces/library/collections/library/native-son-doc-id", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// Document: Create a document with some data
	fmt.Println("Create a reader document with data specifying an ID")
	jsonStr = []byte(`{
			"reader": {
			   "name": "John Smith",
			   "user_id": "12346",
			   "birthdate": "11-01-1992",
			   "email": {
				   "primary": "jsmith@gmail.com",
				   "secondary": "john.smith@aol.com"
			   },
			   "address": {
				   "primary": {
					   "street": "200 Z St",
					   "city": "Evertown",
					   "state": "MA",
					   "zip-code": 55555
				   },
				   "secondary": {
					   "street": "850 2nd St",
					   "city": "Evertown",
					   "state": "MA",
					   "zip-code": 55556
				   }
			   },
			   "reviews": [
				   {
					   "book-title": "Moby Dick", 
					   "rating": 3, 
					   "review-date": "02-02-2002",
					   "comment": "It was okay."
				   },
				   {
					   "book-title": "Pride and Prejudice", 
					   "rating": 5, 
					   "review-date": "03-02-2002",
					   "comment": "It was a wonderful book! I loved reading it."
				   }
			   ]
			}
		}`)

	responsebody, err = client.APIPut("/api/rest/v2/namespaces/library/collections/library/John-Smith", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// Document: Find all documents
	fmt.Println("Retrieve documents")
	responsebody, err = client.APIGet("/api/rest/v2/namespaces/library/collections/library")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// Document: Delete the collection
	fmt.Println("Delete collection")
	responsebody, err = client.APIDelete("/api/rest/v2/namespaces/library/collections/library")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

}
