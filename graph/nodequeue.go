package graph

import "fmt"

type Node struct {
	Key   string
	Size  int
	Color string
}

func NewFullNode(key, color string, size int) *Node {
	return &Node{
		Key:   key,
		Color: color,
		Size:  size,
	}
}

func NewNode(key string) *Node {
	return &Node{
		Key:   key,
		Color: "black",
		Size:  10,
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Key)
}

func (n *Node) WithColor(color string) *Node {
	n.Color = color
	return n
}

func (n *Node) WithSize(size int) *Node {
	n.Size = size
	return n
}

type Edge struct {
	Node   *Node
	Weight int
}

func (e *Edge) String() string {
	return fmt.Sprintf("%v[%d]", e.Node.Key, e.Weight)
}

type Vertex struct {
	Node     *Node
	Distance int
}

type PriorityQueue []*Vertex

type NodeQueue struct {
	Items []Vertex
}

func (nq *NodeQueue) Enqueue(v Vertex) {
	if len(nq.Items) == 0 {
		nq.Items = append(nq.Items, v)
		return
	}
	var insertFlag bool
	for k, vertex := range nq.Items {
		// add vertex distance less than travers's vertex distance
		if v.Distance < vertex.Distance {
			nq.Items = append(nq.Items[:k+1], nq.Items[k:]...)
			nq.Items[k] = v
			insertFlag = true
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		nq.Items = append(nq.Items, v)
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
