package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/awesome-astra/astra_stargate_go"
	graphql "github.com/hasura/go-graphql-client"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if len(os.Getenv("ASTRA_DB_APPLICATION_TOKEN")) == 0 {
		fmt.Println(fmt.Errorf("please set your environment variables or use 'astra db create-dotenv' to create a .env file"))
		return
	}

	astraClient := astra_stargate_go.NewBasicAuthClient(os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_REGION"))
	if err != nil {
		fmt.Println(err)
	}

	client := graphql.NewClient(astraClient.GetURL()+"/graphql-schema)", &http.Client{
		Transport: headerRoundTripper{
			setHeaders: func(req *http.Request) {
				req.Header.Set("x-cassandra-token", os.Getenv("ASTRA_DB_APPLICATION_TOKEN"))
			},
			rt: http.DefaultTransport,
		},
	})

	var query struct {
		Query struct {
			Keyspaces struct {
				Name string
			}
		}
	}

	err = client.Query(context.Background(), &query, nil)

	print(query)
}

type headerRoundTripper struct {
	setHeaders func(req *http.Request)
	rt         http.RoundTripper
}

func (h headerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	h.setHeaders(req)
	return h.rt.RoundTrip(req)
}

func print(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}
