package cmd

import (
	"fmt"
	"os"
)

func Execute(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read directory %s\n error %s", path, err)
		return
	}

	for _, file := range dir {
		fmt.Println(file.Name())
	}
}
