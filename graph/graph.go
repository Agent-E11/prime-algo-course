package graph

import (
	"fmt"
	"math"
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

// Weighted Adjacency List
type WAL = [][]GraphEdge

type GraphEdge struct {
	To int
	Weight int
}

func walkDFS(graph WAL, curr int, needle int, seen *[]bool, path *[]int) bool {

	if curr >= len(*seen) {
		diff := curr - len(*seen) + 1
		*seen = append(*seen, make([]bool, diff)...)
	}

	if (*seen)[curr] {
		return false
	}

	(*seen)[curr] = true

	*path = append(*path, curr)

	if curr == needle {
		return true
	}

	list := graph[curr]

	for i := 0; i < len(list); i++ {
		edge := list[i]

		if walkDFS(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func DFSearch(graph WAL, source int, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}

	walkDFS(graph, source, needle, &seen, &path)

	return path
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i, s := range seen {
		if !s && dists[i] < math.MaxInt {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDistance := math.MaxInt

	for i := 0; i < len(seen); i++ {
		if seen[i] { continue }

		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			idx = i
		}
	}

	return idx
}

func DijkstraList(source int, sink int, arr WAL) []int {
	seen := make([]bool, len(arr))
	prev := make([]int, len(arr))
	for i := range len(prev) {
		prev[i] = -1
	}
	dists := make([]int, len(arr))
	for i := range len(dists) {
		dists[i] = math.MaxInt
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight

			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	path := []int{}
	curr := sink

	for prev[curr] != -1 {
		path = append(path, curr)

		curr = prev[curr]
	}

	path = append(path, source)
	slices.Reverse(path)

	return path
}

func PrintWAL(graph WAL) {
	for self, edges := range graph {
		for _, edge := range edges {
			fmt.Printf("%d --%d-> %d\n", self, edge.Weight, edge.To)
		}
	}
}
