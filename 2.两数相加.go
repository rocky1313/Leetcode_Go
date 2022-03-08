package main

type ListNode struct {
	value int
	next  *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) {
	head1 := &ListNode{0, nil}
	head2 := &ListNode{0, nil}
	res := &ListNode{1, nil}
	num1, num2, sum := 0, 0, 0

}
