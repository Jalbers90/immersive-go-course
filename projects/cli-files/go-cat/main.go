package main

import (
	"fmt"
	"go-cat/cmd"
	"os"
)

func main() {
	var path string
	if len(os.Args) < 2 {
		fmt.Println("Must provide a file path")
		return
	} else {
		path = os.Args[1]
	}

	cmd.Execute(path)
}
