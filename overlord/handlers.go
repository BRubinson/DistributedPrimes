package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Available(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Overlord is Alive and ready for connections")
}

func ConnectNewNode(w http.ResponseWriter, r *http.Request) {
	var newNode Node
	if Leader.NodeCount == 0 {
		newNode.SetValues(Leader.NodeCount, "Master")
		Leader.AddNode(newNode)
	} else {
		newNode.SetValues(Leader.NodeCount, "Sloove")
		Leader.AddNode(newNode)
	}

	if err := json.NewEncoder(w).Encode(newNode); err != nil {
		panic(err)
	}
}
