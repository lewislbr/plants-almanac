package main

import (
	"log"

	"plants/src/api"
)

func main() {
	if err := api.Start(); err != nil {
		log.Fatalln(err)
	}
}
