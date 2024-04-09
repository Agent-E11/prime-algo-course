package binarytree

func BreadthFirstSearch(head *Node[int], needle int) bool {
	q := []*Node[int]{head}

	for len(q) > 0 {
		// Get first element, and remove it from the queue
		curr := q[0]
		q = q[1:]

		if curr == nil {
			continue
		}

		if curr.Value == needle {
			return true
		}

		q = append(q, curr.Left, curr.Right)
	}

	return false
}
