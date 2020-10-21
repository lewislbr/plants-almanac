package main

import (
	"fmt"
	"os"

	"users/pkg/api/rest"
)

func main() {
	if err := rest.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
