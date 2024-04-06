package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	//"time"

	"github.com/agent-e11/prime-algo-course/day1"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	l := day1.NewLinkedList[int]()

	l.Println()

	for {
		if l.Length == 0 {
			fmt.Println("length 0, adding item")
			l.Append(9)
			continue
		}
		idx := rand.Intn(l.Length)
		item := rand.Intn(10)
		choice := rand.Intn(6)

		switch choice {
		case 0:
			fmt.Printf("Prepend(item: %v)\n", item)
			l.Prepend(item)
		case 1:
			fmt.Printf("Append(item: %v)\n", item)
			l.Append(item)
		case 2:
			idx = rand.Intn(l.Length+1)
			fmt.Printf("InsertAt(item: %v, idx: %v)\n", item, idx)
			l.InsertAt(item, idx)
		case 3:
			fmt.Printf("Remove(item: %v) -> ", item)
			v, ok := l.Remove(item)
			fmt.Printf("(value: %v, ok: %v)\n", v, ok)
		case 4:
			fmt.Printf("RemoveAt(idx: %v) -> ", idx)
			v, ok := l.RemoveAt(idx)
			fmt.Printf("(value: %v, ok: %v)\n", v, ok)
		case 5:
			fmt.Printf("Get(idx: %v) -> ", idx)
			v, ok := l.Get(idx)
			fmt.Printf("(value: %v, ok: %v)\n", v, ok)
		default:
			fmt.Println("Impossible")
		}
		l.Debug()

		//time.Sleep(2 * time.Second)

		fmt.Print("Continue... ")
		if s, _ := reader.ReadString('\n'); s == "q\n" {
			break
		}
	}


//	l.Append(7)
//	l.Append(6)
//	l.Append(5)
//	l.Append(6)
//	l.Append(1)
//
//	l.Println()
//
//	l.Remove(4)
//
//	l.Println()
}
