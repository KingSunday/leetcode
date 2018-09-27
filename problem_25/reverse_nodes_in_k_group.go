package problem_25


func reverseKGroup(head *ListNode, k int) *ListNode {

	first, cur := head, head
	count := 1
	for {
		if count == k {
			count = 1
		}
		cur = cur.Next
		count ++
	}
}

func reverseListNode(head, tail *ListNode) {
	start, end := head, tail
	for ; head == tail; head = head.Next {
		start = head
		if end == tail {
			start.Next = end.Next
		} else {
			start.Next = end
		}
		end = start
	}
}