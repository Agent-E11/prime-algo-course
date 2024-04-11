package graph

import (
	"fmt"
	"slices"
)

// Weighted Adjacency Matrix
type WAM = [][]int

func NewWAM(n int) WAM {
	g := make([][]int, n)
	for i := range n {
		g[i] = make([]int, n)
	}
	return g
}

// Breadth First Search
func BFSearch(graph WAM, source int, needle int) []int {
	seen := make([]bool, len(graph))

	prev := make([]int, len(graph))

	for i := range prev { prev[i] = -1 }

	seen[source] = true
	q := []int{source}

	for {
		fmt.Printf("seen: %v\n", seen)
		fmt.Printf("prev: %v\n", prev)
		fmt.Printf("q: %v\n", q)

		curr := q[0]
		q = q[1:]

		if curr == needle {
			break
		}

		cons := graph[curr]
		for other, con := range cons {
			if con == 0 { continue }
			if seen[other] { continue }

			seen[other] = true
			prev[other] = curr

			q = append(q, other)
		}

		if len(q) <= 0 { break }
	}

	// Build the path

	fmt.Println("starting at needle:", needle)
	curr := needle
	path := []int{}
	
	for prev[curr] != -1 {
		fmt.Printf("curr: %v\n", curr)
		path = append(path, curr)
		curr = prev[curr]
	}

	fmt.Printf("path: %v\n", path)

	if len(path) > 0 {
		fmt.Println("Reversing and returning path")
		path = append(path, source)

		slices.Reverse(path)
		return path
	}
	return []int{}
}

func PrintWAM(graph WAM) {
	for self, cons := range graph {
		for other, con := range cons {
			if con != 0 {
				fmt.Printf("%d --%d-> %d\n", self, con, other)
			}
		}
	}
}
