package main

import (
	"fmt"
	"net/http"
)

func Tester(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Api is online :)")
}
