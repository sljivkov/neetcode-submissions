func reorderList(head *ListNode) {
	fast, slow := head.Next, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil
	var prev *ListNode

	for second != nil {
		next := second.Next
		second.Next = prev
		prev = second
		second = next
	}

	first := head
	second = prev

	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next

		first.Next = second
		second.Next = tmp1

		first, second = tmp1, tmp2

	}
}