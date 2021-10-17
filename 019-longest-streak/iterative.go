package longeststreak

func Iterative(head *Node) int {
	var (
		lastNode *Node
	    count int
	    maxCount int
	)

	for current := head; current != nil; current = current.Next {
		if (lastNode != nil) && (lastNode.Value == current.Value) {
			count++
		} else {
			count = 1
		}

		lastNode = current

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
