# Graph shortest path

Creates a graph with nodes and weighted edges.
Able to calculate shortest path from one node to another.
Also create a grpah chart in html.


See example:

```bash
    go run cmd/main.go 
```
Prints to terminal:
```
GRAPH: graph_example

NODE[A]: B[2]  D[1]  C[5] 
NODE[B]: C[3]  D[2] 
NODE[D]: E[1]  C[3] 
NODE[C]: F[5] 
NODE[E]: C[1]  F[2] 
NODE[F]:

SHORTEST PATH: [A D E F]
```


```go
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

	// instantiate the graph
	graphInstance := graph.New(inputGraph)

	// prints graph to terminal
	graphInstance.Print()

	// creates an html file with a graph chart
	if err := graphInstance.CreateChart(); err != nil {
		log.Fatal(err)
	}

	// gets shortest path
	resp := graphInstance.GetShortestPath(a, f)
	fmt.Printf("\nSHORTEST PATH: %s\n", resp.Path)
```

