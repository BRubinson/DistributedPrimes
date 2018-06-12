package main

import (
	"bytes"
	"encoding/json"
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
var Primes []int

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

func getNumbers(masterIp string) {
	r, err := http.Get(masterIp + "/api/getNumbers")
	if err != nil {
		panic(err)
	}
	var meisterNumber NumbersLoad
	defer r.Body.Close()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		panic(err)
	} else {
		if err = json.Unmarshal(body, &meisterNumber); err != nil {
			panic(err)
		} else {
			//This is where you succesfully decoded json
			fmt.Println(meisterNumber.Numbers)
			meisterNumber.Numbers = IsPrimeSlice(meisterNumber.Numbers)
			fmt.Println(meisterNumber.Numbers)
			jstring, _ := json.Marshal(meisterNumber)

			http.Post(masterIp+"/api/foundPrimes", "json", bytes.NewBuffer(jstring))

			time.Sleep(time.Second / 2)

			getNumbers(masterIp)
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
		go getNumbers(DNode.MasterIp)
	} else {
		MasterDisplay()
	}
	//	go func() {
	Init()
	log.Fatal(http.ListenAndServe(NodePort, router))
	//	}()

}
