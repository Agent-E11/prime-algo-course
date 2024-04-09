package main

import (
	"fmt"

	bt "github.com/agent-e11/prime-algo-course/binarytree"
)

func main() {
	n := bt.From(4,
		bt.From(
			2,
			bt.FromValue(1),
			bt.FromValue(3),
		),
		bt.From(
			6,
			bt.FromValue(5),
			bt.FromValue(7),
		),
	)

	n.Debug()

	n.BSTInsert(8)
	n.BSTInsert(3)

	n.Debug()

	fmt.Printf("n.InOrderSearch(): %v\n", n.InOrderSearch())
}
