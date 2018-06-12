package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var overlordPort string = ":8080"
var overlordIp string = "http://localhost"
var OverlordAddr = overlordIp + overlordPort
var NodePort string
var DNode Node

func isNodeConnected(masterIp string) {
	if r, err := http.Get(masterIp + "/api/online"); err != nil {
		panic(err)
	} else {
		defer r.Body.Close()
		if body, err := ioutil.ReadAll(r.Body); err != nil {
			panic(err)
		} else {
			fmt.Println(string(body))
		}

	}
}

func getNumber(masterIp string) {
	if r, err := http.Get(masterIp + "/api/getNumber"); err != nil {
		panic(err)
	} else {
		defer r.Body.Close()
		if body, err := ioutil.ReadAll(r.Body); err != nil {
			panic(err)
		} else {
			fmt.Println(string(body))

			time.Sleep(200)
			getNumber(masterIp)

		}

	}
}

func main() {
	portIntpointer := flag.Int("port", 7777, "an int")
	flag.Parse()
	NodePort = ":" + strconv.Itoa(*portIntpointer)

	DNode.connectLocal(NodePort)
	fmt.Println(DNode)
	router := NewRouter()
	if DNode.Id != 0 {
		isNodeConnected(DNode.MasterIp)
		go getNumber(DNode.MasterIp)
	}
	//	go func() {
	log.Fatal(http.ListenAndServe(NodePort, router))
	//	}()

}
