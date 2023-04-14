package main

import (
	"fmt"
	"os"

	"github.com/synedra/astra_stargate"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("HELLO?")
	err := godotenv.Load()
	fmt.Println(err)
	fmt.Println(os.Getenv("ASTRA_DB_ID"))
	if len(os.Getenv("ASTRA_DB_APPLICATION_TOKEN")) == 0 {
		fmt.Errorf("Please set your environment variables or use astra db create-dotenv to create a .env file")
		panic("Unable to query API")
	}
	client := astra_stargate.NewBasicAuthClient(os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_REGION"))
	if err != nil {
		return
	}
	responsebody, err := client.APICall("/api/rest/v1/keyspaces", "")
	fmt.Println(responsebody)
}
