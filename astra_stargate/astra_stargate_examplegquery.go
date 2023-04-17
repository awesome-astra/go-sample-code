package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/synedra/astra_stargate"

	"github.com/joho/godotenv"
)

type QueryRequestBody struct {
	Query string `json:"query"`
}

func main() {
	err := godotenv.Load()
	if len(os.Getenv("ASTRA_DB_APPLICATION_TOKEN")) == 0 {
		fmt.Println(fmt.Errorf("please set your environment variables or use astra db create-dotenv to create a .env file"))
		return
	}

	client := astra_stargate.NewBasicAuthClient(os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_REGION"))
	if err != nil {
		fmt.Println(err)
	}

	var query = QueryRequestBody{
		Query: `query GetTables {
			keyspace(name: "library") {
				name
			}
		  }
	`,
	}

	jsonStr, err := json.Marshal(query)

	if err != nil {
		panic(err)
	}
	req, err := client.APIPost("/graphql-schema", bytes.NewBuffer(jsonStr))
	fmt.Println(req)
	fmt.Println(err)

}
