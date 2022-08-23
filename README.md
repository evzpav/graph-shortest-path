# Graph shortest path

Creates a graph with weighted edges and calculates shortest path from one node to another


See example:

```bash
    go run cmd/main.go 
```
Prints to terminal:
```
NODE: A:  B[2]  D[1]  C[5] 
NODE: B:  C[3]  D[2] 
NODE: D:  E[1]  C[3] 
NODE: C:  F[5] 
NODE: E:  C[1]  F[2] 
NODE: F: 

SHORTEST PATH
A --> D --> E --> F
```


Creating graph and printing it:

```go
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

	graphInstance := graph.New(inputGraph)
```

Getting shortest path:

```go
resp := graphInstance.GetShortestPath("A", "F")

```


