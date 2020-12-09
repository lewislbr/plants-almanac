package main

import (
	"fmt"
	"os"

	"users/src/api"
)

func main() {
	if err := api.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
