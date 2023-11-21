package cmd

import (
	"fmt"
	"os"
)

func Execute(path string) {
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read file %s error ::: %s", path, err)
		return
	}

	os.Stdout.Write(f)
}
