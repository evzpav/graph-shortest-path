package main

import (
	"fmt"
	"graph-shortest-path/graph"
	"log"
)

func main() {
	inputGraph := graph.InputGraph{
		InputData: []graph.InputData{
			{From: "A", To: "B", Weight: 2},
			{From: "A", To: "D", Weight: 1},
			{From: "A", To: "C", Weight: 5},
			{From: "B", To: "C", Weight: 3},
			{From: "B", To: "D", Weight: 2},
			{From: "D", To: "E", Weight: 1},
			{From: "D", To: "C", Weight: 3},
			{From: "E", To: "C", Weight: 1},
			{From: "E", To: "F", Weight: 2},
			{From: "C", To: "F", Weight: 5},
		},
	}
	if err := inputGraph.Validate(); err != nil {
		log.Fatal(err)
	}

	graphInstance := graph.New(inputGraph)
	graphInstance.Print()

	resp := graphInstance.GetShortestPath("A", "F")

	fmt.Println("\nSHORTEST PATH")
	s := ""
	for i, p := range resp.Path {
		prefix := " --> "
		if i == 0 {
			prefix = ""
		}
		s += prefix + p
	}

	fmt.Println(s)
}
