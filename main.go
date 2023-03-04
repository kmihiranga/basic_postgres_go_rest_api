package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	store, err := NewPostgressStore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%v+\n", store)

	if err := store.Init(ctx); err != nil {
		log.Fatalf("error initializing postgres tables. %v", err)
	}
	server := NewServer(store)
	server.Run(":3000")
}
