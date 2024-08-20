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
	fmt.Println("%+v\n",store)

	// server := NewAPIServer(":8080")
	// server.Run()
}
