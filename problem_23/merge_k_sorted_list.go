package problem_23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var ans, cur *ListNode

	for {
		min := minNode(lists)
		if min == nil {
			break
		}

		if cur == nil {
			ans = *min
			cur = ans
		} else {
			cur.Next = *min
			cur = cur.Next
		}

		(*min) = (*min).Next
	}
	return ans
}

func minNode(lists []*ListNode) (min **ListNode) {
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		if min == nil {
			min = &lists[i]
		} else if (*min).Val > lists[i].Val {
			min = &lists[i]
		}
	}
	return
}
