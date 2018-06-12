package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
