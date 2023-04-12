package main

import (
    "fmt"
    "os"

    "github.com/datastax-ext/astra-go-sdk"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()

    token := os.Getenv("ASTRA_DB_APPLICATION_TOKEN")
    secureBundle := os.Getenv("ASTRA_DB_SECURE_BUNDLE_PATH")
    keyspace := os.Getenv("ASTRA_DB_KEYSPACE")

    c, err := astra.NewStaticTokenClient(
        token,
        astra.WithSecureConnectBundle(secureBundle),
        astra.WithDefaultKeyspace(keyspace),
    )
    if err != nil {
        fmt.Println("Error:")
        fmt.Println(err)
    }

    fmt.Println("SELECTing from system.local")

    rows, err := c.Query("SELECT cluster_name FROM system.local").Exec()
    if err != nil {
        fmt.Println(err)
    }

    for _, r := range rows {
        vals := r.Values()
        strClusterName := vals[0].(string)
        fmt.Println("cluster_name:", strClusterName)
    }
}
