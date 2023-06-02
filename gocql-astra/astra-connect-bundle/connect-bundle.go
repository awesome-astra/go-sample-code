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

	var cluster *gocql.ClusterConfig

	cluster, err := gocqlastra.NewClusterFromBundle(os.Getenv("ASTRA_DB_SECURE_BUNDLE_PATH"),
		"token", os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), 30*time.Second)

	if err != nil {
		panic("unable to load the bundle")
	}
	cluster.Timeout = 30 * time.Second

	session, err := gocql.NewSession(*cluster)

	if err != nil || session == nil {
		log.Fatalf("unable to connect session: %v", err)
	} else {
		fmt.Println("Success!")
	}

}
