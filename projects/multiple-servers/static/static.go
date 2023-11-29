package static

import (
	"fmt"
	"log"
	"net/http"
)

func Run(path, port *string) {
	log.Printf("path: %s\n", *path)
	log.Printf("port: %s\n", *port)

	http.Handle("/", http.FileServer(http.Dir(*path)))

	listenAddr := fmt.Sprintf(":%s", *port)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
