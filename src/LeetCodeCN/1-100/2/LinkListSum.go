package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var l = ListNode{}
	var n = &l
	frontAdd := 0

	for l1 != nil && l2 != nil {
		//fmt.Println("for 1")
		sum := l1.Val + l2.Val+ frontAdd
		//fmt.Println("sum", sum)
		currentNum := sum%10
		frontAdd = sum/10
		n.Val = currentNum
		if l1.Next == nil || l2.Next == nil{
			break
		}
		n.Next = &ListNode{}
		//fmt.Println(l)
		//fmt.Println(n)
		n = n.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	l1 = l1.Next
	l2 = l2.Next


	if l1 != nil {
		n.Next = &ListNode{}
		n = n.Next
		for l1 != nil {
			fmt.Println("for l1")
			sum := l1.Val + frontAdd
			//fmt.Println("sum", sum)
			currentNum := sum%10
			frontAdd = sum/10
			n.Val = currentNum
			if l1.Next == nil{
				break
			}
			n.Next = &ListNode{}
			n = n.Next
			l1 = l1.Next
		}

	} else if l2 != nil{
		n.Next = &ListNode{}
		n = n.Next
		for l2 != nil {
			//fmt.Println("for l2")
			sum := l2.Val + frontAdd
			//fmt.Println("sum", sum)
			currentNum := sum%10
			frontAdd = sum/10
			n.Val = currentNum
			if l2.Next == nil{
				break
			}
			n.Next = &ListNode{}
			n = n.Next
			l2 = l2.Next
		}

	}

	if frontAdd > 0 {
		n.Next = &ListNode{}
		n = n.Next
		n.Val = frontAdd
	}



	return &l
}

func main() {

	//l1 := ListNode{
	//	Val: 	2,
	//	Next: 	&ListNode{
	//		Val: 	4,
	//		Next: 	&ListNode{
	//			Val: 	3,
	//			Next: 	nil,
	//		},
	//	},
	//}
	//
	//l2 := ListNode{
	//	Val: 	5,
	//	Next: 	&ListNode{
	//		Val: 	6,
	//		Next: 	&ListNode{
	//			Val: 	4,
	//			Next: 	nil,
	//		},
	//	},
	//}

	l1 := ListNode{
		Val: 	0,
		Next:	nil,
	}

	l2 := ListNode{
		Val: 	2,
		Next: 	&ListNode{
			Val: 	7,
			Next: 	&ListNode{
				Val: 	8,
				Next: 	nil,
			},
		},
	}

	l:=addTwoNumbers(&l1, &l2)
	fmt.Println("result", l)
	fmt.Println("result")
	for l!=nil {
		fmt.Println(l.Val)
		l = l.Next
	}



}
