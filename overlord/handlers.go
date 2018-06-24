package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Available(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Overlord is Alive and ready for connections")
}

var masterIp string

func ConnectNewNode(w http.ResponseWriter, r *http.Request) {
	var newNode Node

	defer r.Body.Close()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		panic(err)
	} else {

		if err = json.Unmarshal(body, &newNode); err != nil {
			panic(err)
		} else {

			if len(Leader.Nodes) < 1 {
				masterIp = newNode.Ip
				newNode.Id = -1
				Leader.AddNode(newNode)
				Leader.MasterNode = newNode
			} else {
				newNode.Id = -1
				Leader.AddNode(newNode)
			}
			newNode.MasterIp = Leader.MasterNode.Ip
			if err := json.NewEncoder(w).Encode(newNode); err != nil {
				panic(err)
			}
		}
	}
}

func MasterPortal(w http.ResponseWriter, r *http.Request) {

	var payload MasterPayload

	defer r.Body.Close()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		panic(err)
	} else {
		if err = json.Unmarshal(body, &payload); err != nil {
			panic(err)
		} else {
			//This is where you succesfully decoded json
			Leader.AppendPrimes(payload.FoundPrimes)
			Leader.CurrentCount = payload.CurrentValue
			Leader.LastMasterPing = time.Now()
		}
	}
}

func GetCount(w http.ResponseWriter, r *http.Request) {
	if len(Leader.Primes) > 0 {
		fmt.Fprintf(w, strconv.Itoa(Leader.Primes[len(Leader.Primes)-1]+1))
	} else {
		fmt.Fprint(w, "0")
	}
}
