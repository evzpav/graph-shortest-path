package main

import (
	"fmt"
	"graph-shortest-path/graph"
	"log"
)

func main() {

	a := graph.NewNode("A").WithColor("#1984c7")
	b := graph.NewNode("B").WithColor("#c76919").WithSize(30)
	c := graph.NewNode("C").WithColor("#8419c7")
	d := graph.NewNode("D").WithColor("#c71969")
	e := graph.NewNode("E").WithColor("#c79f19").WithSize(40)
	f := graph.NewNode("F").WithColor("#199fc7").WithSize(60)

	inputGraph := graph.InputGraph{
		Name: "graph_example",
		InputEdges: []graph.InputEdges{
			{From: a, To: b, Weight: 2},
			{From: a, To: d, Weight: 1},
			{From: a, To: c, Weight: 5},
			{From: b, To: c, Weight: 3},
			{From: b, To: d, Weight: 2},
			{From: d, To: e, Weight: 1},
			{From: d, To: c, Weight: 3},
			{From: e, To: c, Weight: 1},
			{From: e, To: f, Weight: 2},
			{From: c, To: f, Weight: 5},
		},
	}

	graphInstance := graph.New(inputGraph)
	graphInstance.Print()

	if err := graphInstance.CreateChart(); err != nil {
		log.Fatal(err)
	}

	resp := graphInstance.GetShortestPath(a, f)
	fmt.Printf("\nSHORTEST PATH: %s\n", resp.Path)

}
