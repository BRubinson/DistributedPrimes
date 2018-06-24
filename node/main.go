package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var overlordPort string = ":8080"
var overlordIp string = "http://localhost"
var OverlordAddr = overlordIp + overlordPort
var NodePort string
var DNode Node
var Primes []int

func main() {
	portIntpointer := flag.Int("port", 7777, "an int")
	flag.Parse()
	NodePort = ":" + strconv.Itoa(*portIntpointer)

	DNode.connectLocal(NodePort)
	fmt.Println(DNode)

	go DNode.run()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(NodePort, router))

}
