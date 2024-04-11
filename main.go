package main

import (
	"fmt"

	"github.com/agent-e11/prime-algo-course/graph"
)

func main() {
	g := graph.WAL{
		[]graph.GraphEdge{ // 0
			{ To: 2, Weight: 20 },
			{ To: 1, Weight: 1 },
		},
		[]graph.GraphEdge{ // 1
			{ To: 2, Weight: 10 },
			{ To: 3, Weight: 6 },
		},
		[]graph.GraphEdge{ // 2
			{ To: 4, Weight: 1 },
		},
		[]graph.GraphEdge{ // 3
			{ To: 2, Weight: 1 },
		},
		[]graph.GraphEdge{ // 4
		},
	}

	graph.PrintWAL(g)

	path := graph.DFSearch(g, 0, 4)
	fmt.Printf("path: %v\n", path)

	fmt.Printf("graph.DijkstraList(0, 4, g): %v\n", graph.DijkstraList(0, 4, g))
}
