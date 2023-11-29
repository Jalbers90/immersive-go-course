package main

import (
	"flag"
	"multiple-servers/api"

	_ "github.com/lib/pq"
)

func main() {
	port := flag.String("port", "8081", "Port to run static server on")
	flag.Parse()
	dbURL := "postgres://postgres:mysecretpassword@localhost:7676/postgres?sslmode=disable"
	api.Run(port, dbURL)
}
