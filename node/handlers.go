package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func NodeOnline(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Node is online :)")
}

var count int = 0

func NewNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, strconv.Itoa(count))
	count++
}
