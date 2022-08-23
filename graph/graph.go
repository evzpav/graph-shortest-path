package graph

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type InputGraph struct {
	Name       string
	InputData  []InputData
	InputEdges []InputEdges
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

type InputEdges struct {
	From   *Node
	To     *Node
	Weight int
}

type Response struct {
	Path     []string `json:"path"`
	Distance int      `json:"distance"`
}

type Graph struct {
	Name  string
	Nodes []*Node
	Edges map[Node][]*Edge
}

func New(inputGraph InputGraph) *Graph {
	g := Graph{Name: inputGraph.Name}
	nodes := make(map[string]*Node)
	for _, n := range inputGraph.InputEdges {
		from := n.From
		if _, found := nodes[from.Key]; !found {
			nA := NewFullNode(from.Key, from.Color, from.Size)
			nodes[from.Key] = nA
			g.AddNode(nA)
		}
		to := n.To

		if _, found := nodes[to.Key]; !found {
			nA := NewFullNode(to.Key, to.Color, to.Size)
			nodes[to.Key] = nA
			g.AddNode(nA)
		}

		g.AddEdge(nodes[from.Key], nodes[to.Key], n.Weight)
	}
	return &g
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		Node:   n2,
		Weight: weight,
	}

	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
}

func (g *Graph) Print() {
	fmt.Printf("GRAPH: %s\n", g.Name)

	for _, node := range g.Nodes {
		fmt.Printf("\nNODE[%s]:", node.String())
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
		dist[nval.Key] = math.MaxInt64
	}
	dist[startNode.Key] = start.Distance
	pq.Enqueue(start)

	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.Node.Key] {
			continue
		}
		visited[v.Node.Key] = true
		near := g.Edges[*v.Node]

		for _, val := range near {
			if !visited[val.Node.Key] {
				if dist[v.Node.Key]+val.Weight < dist[val.Node.Key] {
					store := Vertex{
						Node:     val.Node,
						Distance: dist[v.Node.Key] + val.Weight,
					}
					dist[val.Node.Key] = dist[v.Node.Key] + val.Weight

					prev[val.Node.Key] = v.Node.Key
					pq.Enqueue(store)
				}
			}
		}
	}

	pathval := prev[endNode.Key]
	var finalArr []string
	finalArr = append(finalArr, endNode.Key)
	for pathval != startNode.Key {
		finalArr = append(finalArr, pathval)
		pathval = prev[pathval]
	}
	finalArr = append(finalArr, pathval)

	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	return finalArr, dist[endNode.Key]

}

func (g *Graph) GetShortestPath(from, to *Node) *Response {
	path, distance := g.getShortestPath(from, to)
	return &Response{
		Path:     path,
		Distance: distance,
	}
}

func (g *Graph) CreateChart() error {
	page := components.NewPage()

	page.AddCharts(graphChart(g))

	file, err := os.Create(g.Name + ".html")
	if err != nil {
		return err
	}

	if err := page.Render(io.MultiWriter(file)); err != nil {
		return err
	}

	return nil
}

func graphChart(g *Graph) *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: g.Name}),
	)

	links := make([]opts.GraphLink, 0)
	for n, edges := range g.Edges {
		for _, e := range edges {
			links = append(links, opts.GraphLink{
				Source: n.Key,
				Target: e.Node.Key,
				Value:  float32(e.Weight * 100),
			})
		}
	}

	graphNodes := make([]opts.GraphNode, 0)
	for _, n := range g.Nodes {
		graphNodes = append(graphNodes, opts.GraphNode{
			Name:       n.Key,
			SymbolSize: n.Size,
			ItemStyle:  &opts.ItemStyle{Color: n.Color},
		})
	}

	graph.AddSeries("graph", graphNodes, links,
		charts.WithGraphChartOpts(
			opts.GraphChart{
				Layout: "force",
				Force: &opts.GraphForce{
					Repulsion:  1000,
					Gravity:    1,
					EdgeLength: 30,
				},
			},
		),
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "right",
		}),
	)

	return graph
}
