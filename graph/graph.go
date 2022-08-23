package graph

import (
	"fmt"
	"math"
	"sync"
)

type InputGraph struct {
	InputData []InputData
}

func (ig *InputGraph) Validate() error {
	for _, v := range ig.InputData {
		if v.From == "" {
			return fmt.Errorf("from is required")
		}

		if v.To == "" {
			return fmt.Errorf("to is required")
		}

		if v.Weight < 0 {
			return fmt.Errorf("weight cannot be negative")
		}
	}

	return nil
}

type InputData struct {
	From   string
	To     string
	Weight int
}

type Response struct {
	Path     []string `json:"path"`
	Distance int      `json:"distance"`
}

type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	lock  sync.RWMutex
}

func New(data InputGraph) *Graph {
	var g Graph
	nodes := make(map[string]*Node)
	for _, v := range data.InputData {

		if _, found := nodes[v.From]; !found {
			nA := Node{v.From}
			nodes[v.From] = &nA
			g.AddNode(&nA)
		}
		if _, found := nodes[v.To]; !found {
			nA := Node{v.To}
			nodes[v.To] = &nA
			g.AddNode(&nA)
		}
		g.AddEdge(nodes[v.From], nodes[v.To], v.Weight)
	}
	return &g
}

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	g.Nodes = append(g.Nodes, n)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		Node:   n2,
		Weight: weight,
	}

	g.Edges[*n1] = append(g.Edges[*n1], &ed1)

	g.lock.Unlock()
}

func (g *Graph) Print() {
	for _, node := range g.Nodes {
		fmt.Printf("\nNODE: %v: ", node.String())
		for _, edge := range g.Edges[*node] {
			fmt.Printf(" %v ", edge.String())
		}
	}
	fmt.Println()
}

func (g *Graph) getShortestPath(startNode *Node, endNode *Node) ([]string, int) {
	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)

	q := NodeQueue{}
	pq := q.NewQueue()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	for _, nval := range g.Nodes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startNode.Value] = start.Distance
	pq.Enqueue(start)

	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true
		near := g.Edges[*v.Node]

		for _, val := range near {
			if !visited[val.Node.Value] {
				if dist[v.Node.Value]+val.Weight < dist[val.Node.Value] {
					store := Vertex{
						Node:     val.Node,
						Distance: dist[v.Node.Value] + val.Weight,
					}
					dist[val.Node.Value] = dist[v.Node.Value] + val.Weight

					prev[val.Node.Value] = v.Node.Value
					pq.Enqueue(store)
				}
			}
		}
	}

	pathval := prev[endNode.Value]
	var finalArr []string
	finalArr = append(finalArr, endNode.Value)
	for pathval != startNode.Value {
		finalArr = append(finalArr, pathval)
		pathval = prev[pathval]
	}
	finalArr = append(finalArr, pathval)

	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	return finalArr, dist[endNode.Value]

}

func (g *Graph) GetShortestPath(from, to string) *Response {
	path, distance := g.getShortestPath(&Node{from}, &Node{to})
	return &Response{
		Path:     path,
		Distance: distance,
	}
}
