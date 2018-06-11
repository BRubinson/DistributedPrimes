package main

import (
	"log"
	"net/http"
)

var Leader = Overlord{
	-1,
	0,
	make(map[int64]Node)}

func main() {
	port := ":8080"
	router := NewRouter()
	log.Fatal(http.ListenAndServe(port, router))
}
