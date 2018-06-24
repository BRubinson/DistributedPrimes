package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Node struct {
	Id       int64
	Ip       string
	MasterIp string
}

func (node *Node) String() string {
	return "ID: " + string(node.Id) + " Val: " + node.Ip + " Master IP: " + node.MasterIp
}
func (node *Node) SetValues(id int64, ip string) {
	node.Id = id
	node.Ip = ip
}

type connectLocalJson struct {
	Ip string
}

func (node *Node) connectLocal(port string) {
	nodeIp := connectLocalJson{"http://localhost" + port}
	jstring, _ := json.Marshal(nodeIp)
	if r, err := http.Post(OverlordAddr+"/api/connect", "json", bytes.NewBuffer(jstring)); err != nil {
		panic(err)
	} else {
		defer r.Body.Close()
		if body, err := ioutil.ReadAll(r.Body); err != nil {
			panic(err)
		} else {

			if err = json.Unmarshal(body, node); err != nil {
				panic(err)
			} else {

			}
		}

	}

}

func (node *Node) run() {
	if node.Ip != node.MasterIp {
		r, err := http.Get(node.MasterIp + "/api/getNumbers")
		if err != nil {
			// response FAIL
			time.Sleep(time.Second * 4)
			node.run()
		} else {
			var meisterNumber NumbersLoad
			defer r.Body.Close()
			if body, err := ioutil.ReadAll(r.Body); err != nil {
				//body read fail
				panic(err)
			} else {
				if err = json.Unmarshal(body, &meisterNumber); err != nil {
					//decode failure
					panic(err)
				} else {
					//This is where you succesfully decoded json
					fmt.Println(meisterNumber.Numbers)
					meisterNumber.Numbers = IsPrimeSlice(meisterNumber.Numbers)
					fmt.Println(meisterNumber.Numbers)
					jstring, _ := json.Marshal(meisterNumber)

					http.Post(node.MasterIp+"/api/foundPrimes", "json", bytes.NewBuffer(jstring))

					time.Sleep(time.Second / 2)

					node.run()
				}
			}
		}
	} else {
		node.IsMaster()
	}
}

type MasterPayload struct {
	FoundPrimes  []int
	CurrentValue int
}

var LastLen = 0

func (node *Node) IsMaster() {
	Init()
	MasterDisplay()

	//pings overlord with new info
	go func() {
		for {
			waitTime := time.Second * 2
			payload := MasterPayload{[]int{}, NextNumber}
			curLength := len(Primes)
			payload.FoundPrimes = Primes[LastLen:curLength]

			jstring, _ := json.Marshal(payload)

			http.Post(OverlordAddr+"/api/master", "json", bytes.NewBuffer(jstring))
			LastLen = curLength
			time.Sleep(waitTime)

		}
	}()
}
