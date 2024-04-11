package main

import (
	"fmt"

	"github.com/agent-e11/prime-algo-course/graph"
)

func main() {
	g := graph.WAL{
		[]graph.GraphEdge{ // 0
			{ To: 2, Weight: 5 },
			{ To: 1, Weight: 1 },
		},
		[]graph.GraphEdge{ // 1
			{ To: 2, Weight: 7 },
			{ To: 3, Weight: 6 },
		},
		[]graph.GraphEdge{ // 2
			{ To: 4, Weight: 1 },
		},
		[]graph.GraphEdge{ // 3
			{ To: 2, Weight: 1 },
		},
	}

	path := graph.DFSearch(g, 0, 4)
	fmt.Printf("path: %v\n", path)
}
