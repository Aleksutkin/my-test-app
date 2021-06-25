package main

import (
	"log"
	"my-test-app/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		log.Fatal(err)
	}
}
