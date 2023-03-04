package main

import (
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%v+\n", store)

	server := NewServer(":3000", store)
	server.Run()
}
