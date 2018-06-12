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

type Overlord struct {
	MasterId  int64
	NodeCount int64
	Nodes     map[int64]Node
}

func (lord *Overlord) AddNode(newNode Node) {
	lord.Nodes[newNode.Id] = newNode
	lord.NodeCount++
}

func (lord *Overlord) SetMasterNodeId(id int64) {
	lord.MasterId = id
}
func (lord *Overlord) SetMasterNode(node Node) {
	lord.MasterId = node.Id
}

func (lord *Overlord) getMasterNode() Node {
	return lord.Nodes[lord.MasterId]
}
