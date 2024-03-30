package main

import (
	"github.com/agent-e11/prime-algo-course/day1"
	"fmt"
)

func main() {
	b0 := []bool{false, false, false, false, false, false, false, false, false}

	i := day1.TwoCrystalBalls(b0)


	fmt.Println("First break:", i)
}
