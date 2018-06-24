package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var NextNumber = 0
var nextNumberChan = make(chan int)

func Init() {
	go func() {
		r, _ := http.Get(OverlordAddr + "/api/getnum")
		if body, err := ioutil.ReadAll(r.Body); err != nil {
			panic(err)
		} else {
			NextNumber, _ = strconv.Atoi(string(body))
		}
		for {
			if len(nextNumberChan) < 50 {
				for len(nextNumberChan) < 100 {
					nextNumberChan <- NextNumber
					NextNumber++
				}
			}

		}
		time.Sleep(2 * time.Second)
	}()
}

func NodeOnline(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Node is online :)")
}

func UpdateNode(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		panic(err)
	} else {
		if err = json.Unmarshal(body, &DNode); err != nil {
			panic(err)
		} else {
			fmt.Println(DNode)
			//This is where you succesfully decoded json
		}
	}
}

type NumbersLoad struct {
	Numbers []int
}

func GetNumbers(w http.ResponseWriter, r *http.Request) {

	var number = make([]int, 10, 10)
	coutn := 0
	for coutn < 10 {
		number[coutn] = <-nextNumberChan
		coutn++
	}
	payload := NumbersLoad{number}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		panic(err)
	}
}

var PrimeCount int = 0

func FoundPrime(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meisterNumber struct {
		Numbers []int
	}
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		panic(err)
	} else {
		if err = json.Unmarshal(body, &meisterNumber); err != nil {
			panic(err)
		} else {
			//This is where you succesfully decoded json
			PrimeCount += len(meisterNumber.Numbers)
			Primes = append(Primes, meisterNumber.Numbers...)
		}
	}
}
