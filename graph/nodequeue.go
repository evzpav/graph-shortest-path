package graph

import "fmt"

type Node struct {
	Value string
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type Edge struct {
	Node   *Node
	Weight int
}

func (e *Edge) String() string {
	return fmt.Sprintf("%v[%d]", e.Node.Value, e.Weight)
}

type Vertex struct {
	Node     *Node
	Distance int
}

type PriorityQueue []*Vertex

type NodeQueue struct {
	Items []Vertex
}

func (nq *NodeQueue) Enqueue(t Vertex) {
	if len(nq.Items) == 0 {
		nq.Items = append(nq.Items, t)
		return
	}
	var insertFlag bool
	for k, v := range nq.Items {
		// add vertex distance less than travers's vertex distance
		if t.Distance < v.Distance {
			nq.Items = append(nq.Items[:k+1], nq.Items[k:]...)
			nq.Items[k] = t
			insertFlag = true
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		nq.Items = append(nq.Items, t)
	}
}

func (nq *NodeQueue) Dequeue() *Vertex {
	item := nq.Items[0]
	nq.Items = nq.Items[1:len(nq.Items)]
	return &item
}

func (nq *NodeQueue) NewQueue() *NodeQueue {
	nq.Items = []Vertex{}
	return nq
}

func (nq *NodeQueue) IsEmpty() bool {
	return len(nq.Items) == 0
}

func (nq *NodeQueue) Size() int {
	return len(nq.Items)
}
