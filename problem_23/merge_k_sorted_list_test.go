package problem_23

import (
	"strconv"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	head1, backup1 := buildListNode(1, 4, 5), buildListNode(1, 4, 5)
	head2, backup2 := buildListNode(1, 3, 4), buildListNode(1, 3, 4)
	head3, backup3 := buildListNode(2, 6), buildListNode(2, 6)
	expected := buildListNode(1, 1, 2, 3, 4, 4, 5, 6)

	head := mergeKLists([]*ListNode{head1, head2, head3})
	if !checkListEqual(head, expected) {
		t.Error(
			"\n three sorted list: ",
			stringifyListNode(backup1, backup1),
			stringifyListNode(backup2, backup2),
			stringifyListNode(backup3, backup3),
			"\n the expected result is ",
			stringifyListNode(expected, expected),
			"\n but the actual result is ",
			stringifyListNode(head, head),
		)
	}
}

func buildListNode(lists ...int) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return &ListNode{
		Val:  lists[0],
		Next: buildListNode(lists[1:]...),
	}
}

func stringifyListNode(head *ListNode, cur *ListNode) (str string) {
	if cur == nil {
		return
	}
	if cur != head {
		str += "->"
	}
	str += strconv.Itoa(cur.Val)
	str += stringifyListNode(head, cur.Next)
	return
}

func checkListEqual(a *ListNode, b *ListNode) bool {
	for {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		a, b = a.Next, b.Next
	}
}
