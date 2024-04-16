package day1

type Point struct {
	Y int
	X int
}

var dirs [4][2]int = [4][2]int{
	{-1,  0},
	{ 0,  1},
	{ 1,  0},
	{ 0, -1},
}

func walk(maze []string, wall rune, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	// Base cases

	// Current point is off the map
	if curr.X < 0 || curr.X >= len(maze[0]) ||
		curr.Y < 0 || curr.Y >= len(maze) {
		
		return false
	}

	// Current point is a wall
	if rune(maze[curr.Y][curr.X]) == wall {
		return false
	}

	// Current point is the end
	if curr == end {
		*path = append(*path, curr)
		return true
	}

	// Current point has been visited already
	if seen[curr.Y][curr.X] {
		return false
	}

	// Recursive case
	seen[curr.Y][curr.X] = true
	*path = append(*path, curr)


	for _, dir := range dirs {
		y, x := dir[0], dir[1]
		if walk(
			maze,
			wall,
			Point{ Y: curr.Y + y, X: curr.X + x },
			end,
			seen,
			path,
		) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func Solve(maze []string, wall rune, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[0]))
	}

	path := []Point{}

	walk(maze, wall, start, end, seen, &path)

	return path
}
