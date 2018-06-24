package main

type Node struct {
	Id       int64
	Ip       string
	MasterIp string
}

func (node *Node) SetValues(id int64, ip string) {
	node.Id = id
	node.Ip = ip
}

type MasterPayload struct {
	FoundPrimes  []int
	CurrentValue int
}
