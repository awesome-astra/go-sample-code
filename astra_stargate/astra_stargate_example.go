package main

import (
	"fmt"
	"os"

	"github.com/synedra/astra_stargate"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if len(os.Getenv("ASTRA_DB_APPLICATION_TOKEN")) == 0 {
		fmt.Println(fmt.Errorf("Please set your environment variables or use astra db create-dotenv to create a .env file"))
		return
	}

	client := astra_stargate.NewBasicAuthClient(os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_REGION"))
	if err != nil {
		fmt.Println(err)
	}
	responsebody, err := client.APICall("/api/rest/v1/keyspaces", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)
}
