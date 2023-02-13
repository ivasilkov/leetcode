package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptrA, ptrB := headA, headB
	for ptrA != ptrB {
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}
		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}
	return ptrA
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := make(map[*ListNode]struct{})
	for {
		if headA != nil {
			if _, ok := m[headA]; ok {
				return headA
			}
			m[headA] = struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := m[headB]; ok {
				return headB
			}
			m[headB] = struct{}{}
			headB = headB.Next
		}
		if headA == nil && headB == nil {
			break
		}
	}

	return nil
}
