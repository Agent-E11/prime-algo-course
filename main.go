package main

import (
	"fmt"

	"github.com/agent-e11/prime-algo-course/day1"
)

func main() {
	maze := []string{
		"XXXXXXXXXX$X",
		"X        X X",
		"X        X X",
		"X XXXXXXXX X",
		"X          X",
		"X^XXXXXXXXXX",
	}
	start := day1.Point{
		Y: 5,
		X: 1,
	}
	end := day1.Point{
		Y: 0,
		X: 10,
	}

	fmt.Println("Path:", day1.Solve(maze, "X", start, end))
}
