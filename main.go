package main

import (
	"fmt"

	"github.com/agent-e11/prime-algo-course/graph"
)

func main() {
	g := graph.NewWAM(5)

	g[3][1] = 1
	g[1][2] = 2
	g[2][4] = 3
	g[0][4] = 4

	graph.PrintWAM(g)

	path := graph.BFSearch(g, 3, 4)

	fmt.Printf("path: %v\n", path)

	fmt.Println("Total weight")

	curr := path[0]
	weight := 0
	for i := 1; i < len(path); i++ {
		weight += g[curr][path[i]]
		curr = path[i]
	}

	fmt.Printf("weight: %v\n", weight)
}
