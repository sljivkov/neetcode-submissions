func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var curr *ListNode
	dummy := &ListNode{}

	if list1 == nil {return list2}
	if list2 == nil {return list1}

	if list1.Val < list2.Val {
		curr = list1
		list1 = list1.Next
	} else {
		curr = list2
		list2 = list2.Next
	}

	dummy.Next = curr

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
			curr = curr.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
			curr = curr.Next
		}
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next
}