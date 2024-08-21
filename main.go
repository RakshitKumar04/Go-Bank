package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n",store)

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8080", store)
	server.Run()
}
