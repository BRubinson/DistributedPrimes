package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var overlordPort string = ":8080"
var overlordIp string = "http://localhost"
var OverlordAddr = overlordIp + overlordPort
var node Node

func main() {
	/* port := ":7777"
	router := NewRouter()
	go log.Fatal(http.ListenAndServe(port, router)) */
	node.connect()
	fmt.Println(node)
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
