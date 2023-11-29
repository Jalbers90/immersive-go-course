package main

import (
	"flag"
	"multiple-servers/static"
)

func main() {
	path := flag.String("path", "assets", "specify a path to serve static assets from")
	port := flag.String("port", "8082", "Port to run static server on")
	flag.Parse()

	static.Run(path, port)
}
