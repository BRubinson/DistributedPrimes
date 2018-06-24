package main

import (
	"log"
	"net/http"
	"time"
)

var Leader = Overlord{
	Node{-2, "", ""},
	[]Node{},
	[]int{},
	0,
	time.Now()}

func main() {
	OverlordDisplay()
	Leader.ManageMaster()
	port := ":8080"
	router := NewRouter()
	log.Fatal(http.ListenAndServe(port, router))
}
