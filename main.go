package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Printf("%#v \n", store)
	server := newAPIServer(":8000", store)
	server.Run()
}
