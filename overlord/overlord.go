package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Overlord struct {
	MasterNode     Node
	Nodes          []Node
	Primes         []int
	CurrentCount   int
	LastMasterPing time.Time
}

func (lord *Overlord) AddNode(newNode Node) {
	lord.Nodes = append(lord.Nodes, newNode)
}

func (lord *Overlord) AppendPrimes(primes []int) {
	lord.Primes = append(lord.Primes, primes...)
}

func (lord *Overlord) ManageMaster() {
	go func() {
		maxDelay := time.Second * 6
		waitTime := time.Second * 3
		for {
			if lord.MasterNode.Id != -2 {
				elapsed := time.Now().Sub(Leader.LastMasterPing)
				if elapsed > maxDelay {
					// assign new master Node
					lord.NewMaster()
					time.Sleep(30 * time.Second)
				} else {
					time.Sleep(waitTime)
				}
			}
		}
	}()
}

func (lord *Overlord) NewMaster() {
	if len(lord.Nodes) > 0 {
		lord.Nodes = lord.Nodes[1:]
		if len(lord.Nodes) > 1 {
			lord.MasterNode = lord.Nodes[0]
			for _, node := range lord.Nodes {
				node.MasterIp = Leader.MasterNode.Ip
				jstring, _ := json.Marshal(node)

				http.Post(node.Ip+"/api/update", "json", bytes.NewBuffer(jstring))
			}
		}
	}
}
