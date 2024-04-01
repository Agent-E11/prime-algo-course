package day1

type Point struct {
	Y int
	X int
}

func walk(maze []string, wall string, curr Point, end Point, seen [][]bool) bool {
	// Base cases

	// Current point is off the map
	if curr.X < 0 || curr.X >= len(maze[0]) ||
		curr.Y < 0 || curr.Y >= len(maze) {
		
		return false
	}

	// Current point is a wall
	if maze[curr.Y][curr.X] == wall[0] {
		return false
	}

	// Current point is the end
	if curr == end {
		return true
	}

	// Current point has been visited already
	if seen[curr.Y][curr.X] {
		return false
	}

	return false
}

func Solve(maze []string, wall string, start Point, end Point) []Point {
	return []Point{}
}
