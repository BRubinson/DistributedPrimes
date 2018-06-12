package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

		}
	}
	newNode.MasterIp = masterIp
	if Leader.NodeCount == 0 {
		masterIp = newNode.Ip
		newNode.Id = Leader.NodeCount
		Leader.AddNode(newNode)
	} else {
		newNode.Id = Leader.NodeCount
		Leader.AddNode(newNode)
	}

	if err := json.NewEncoder(w).Encode(newNode); err != nil {
		panic(err)
	}
}
