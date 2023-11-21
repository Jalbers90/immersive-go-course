package main

import (
	"flag"
	"fmt"
	"go-ls/cmd"
	"os"
)

func main() {

	helpFlag := flag.Bool("h", false, "Show a helpful description")
	flag.Parse()

	if *helpFlag {
		fmt.Println("go-ls prints out file names in a given directory. Defaults to current working directory.")
		return
	}

	path := "./"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	cmd.Execute(path)
}
