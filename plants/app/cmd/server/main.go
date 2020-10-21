package main

import (
	"fmt"
	"os"

	"plants/pkg/api/graphql"
)

func main() {
	if err := graphql.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
