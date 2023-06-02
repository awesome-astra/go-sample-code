package main

import (
	"fmt"
	"log"
	"os"
	"time"

	gocqlastra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
)

func main() {

	var err error

	var cluster *gocql.ClusterConfig

	if len(os.Getenv("ASTRA_DB_APPLICATION_TOKEN")) > 0 {

		if len(os.Getenv("ASTRA_DB_ID")) == 0 {
			panic("database ID is required when using a token")
		}
	}

	cluster, err = gocqlastra.NewClusterFromURL("https://api.astra.datastax.com", os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), 10*time.Second)

	if err != nil {
		fmt.Errorf("unable to load cluster %s from astra: %v", os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), err)
	}

	cluster.Timeout = 30 * time.Second
	start := time.Now()
	session, err := gocql.NewSession(*cluster)
	elapsed := time.Now().Sub(start)

	if err != nil {
		log.Fatalf("unable to connect session: %v", err)
	}

	fmt.Println("Making the query now")

	iter := session.Query("SELECT release_version FROM system.local").Iter()

	var version string

	for iter.Scan(&version) {
		fmt.Println(version)
	}

	if err = iter.Close(); err != nil {
		log.Printf("error running query: %v", err)
	}

	fmt.Printf("Connection process took %s\n", elapsed)

	if err != nil {
		log.Fatalf("unable to connect session: %v", err)
	}

}
