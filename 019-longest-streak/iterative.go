package longeststreak

func Iterative(head *Node) int {
	var (
		prev     *Node
		count    int
		maxCount int
	)

	for current := head; current != nil; current = current.Next {
		if (prev != nil) && (prev.Val == current.Val) {
			count++
		} else {
			count = 1
		}

		prev = current

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
