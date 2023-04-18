package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"

	"github.com/synedra/astra_stargate"

	"github.com/joho/godotenv"
)

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

	// Basic REST Query
	fmt.Println("Basic REST query for keyspaces")
	responsebody, err := client.APIGet("/api/rest/v1/keyspaces")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// REST: Create a table
	fmt.Println("Create a table using REST")
	var jsonStr = []byte(`{
		"name": "users",
		"columnDefinitions":
		  [
			{
			  "name": "firstname",
			  "typeDefinition": "text"
			},
			{
			  "name": "lastname",
			  "typeDefinition": "text"
			},
			{
			  "name": "favorite color",
			  "typeDefinition": "text"
			}
		  ],
		"primaryKey":
		  {
			"partitionKey": ["firstname"],
			"clusteringKey": ["lastname"]
		  },
		"tableOptions":
		  {
			"defaultTimeToLive": 0,
			"clusteringExpression":
			  [{ "column": "lastname", "order": "ASC" }]
		  }
	}`)

	responsebody, err = client.APIPost("/api/rest/v2/schemas/keyspaces/library/tables", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	whereString := "{\"firstname\":{\"$in\":[\"Mookie\",\"Janesha\"]}}"
	whereEncoded := url.QueryEscape(whereString)

	// REST: Check the table
	fmt.Println("Check table")
	responsebody, err = client.APIGet("/api/rest/v2/keyspaces/library/users?where=" + whereEncoded)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// REST: Create a row
	fmt.Println("Create rows in the users table using REST")
	jsonStr = []byte(`{
		"firstname": "Mookie",
		"lastname": "Betts",
		"favorite color": "blue"
	}'`)

	responsebody, err = client.APIPost("/api/rest/v2/keyspaces/library/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	jsonStr = []byte(`{
		"firstname": "Janesha",
		"lastname": "Doesha",
		"favorite color": "grey"
	}'`)

	responsebody, err = client.APIPost("/api/rest/v2/keyspaces/library/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// REST: Check the table
	fmt.Println("Check table")
	responsebody, err = client.APIGet("/api/rest/v2/keyspaces/library/users?where=" + whereEncoded)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// REST: Update a row
	fmt.Println("Update a row")
	jsonStr = []byte(`{ "favorite color": "Fuchsia"}`)

	responsebody, err = client.APIPut("/api/rest/v2/keyspaces/library/users/Janesha/Doesha", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// REST: Check the table
	fmt.Println("Check table")
	responsebody, err = client.APIGet("/api/rest/v2/keyspaces/library/users?where=" + whereEncoded)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// REST: Delete a user
	fmt.Println("Delete a user")
	responsebody, err = client.APIDelete("/api/rest/v2/keyspaces/library/users/Janesha/Doesha")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

	// REST: Delete the table
	fmt.Println("Delete the table")
	responsebody, err = client.APIDelete("/api/rest/v2/schemas/keyspaces/library/tables/users")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsebody)

}
