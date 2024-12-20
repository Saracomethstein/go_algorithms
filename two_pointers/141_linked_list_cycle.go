package two_pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

// link to the task on leetcode
// https://leetcode.com/problems/linked-list-cycle/description/

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
