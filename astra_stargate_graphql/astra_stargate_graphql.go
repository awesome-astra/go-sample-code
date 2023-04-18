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
		fmt.Println(fmt.Errorf("please set your environment variables or use 'astra db create-dotenv' to create a .env file"))
		return
	}

	client := astra_stargate_go.NewBasicAuthClient(os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_REGION"))
	if err != nil {
		fmt.Println(err)
	}

	query := "{\"query\":\"query GetTables {keyspace(name: \\\"library\\\") {name}}\"}"
	queryBody := []byte(query)
	bodyReader := bytes.NewBuffer(queryBody)

	if err != nil {
		panic(err)
	}
	req, err := client.APIPost("/api/graphql-schema", bodyReader)
	if err != nil {
		panic(err)
	}
	fmt.Println(req)

}
