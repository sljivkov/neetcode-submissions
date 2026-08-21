func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := 0
	dummy := &ListNode{}
	prev := dummy
    
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		 
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + res

		digit := sum % 10
		res = sum / 10

		curr := &ListNode{Val: digit}
		prev.Next = curr
		prev = curr

	}

	if res != 0 {
		prev.Next = &ListNode{Val: res}
	}

	return dummy.Next
}