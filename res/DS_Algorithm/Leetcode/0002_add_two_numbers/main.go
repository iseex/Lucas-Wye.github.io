package main

import "fmt"


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
 

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {    
    head := &ListNode{
        Val:  0,
        Next: nil,
	}   
    pos := head
    
    carry := 0 
    sum := 0
    
    for l1 != nil || l2 != nil {
        sum = carry
        if l1 != nil {
            sum = sum + l1.Val
        }
        if l2 != nil {
            sum = sum + l2.Val
        }
        
        carry = sum / 10
        sum = sum % 10
        
        pos.Next = &ListNode {
			Val:  sum,
			Next: nil,
		}
        pos = pos.Next
        
        if l1 != nil {
            l1 = l1.Next 
        }
        if l2 != nil {
            l2 = l2.Next
        }                             
        
    }
    
    if carry > 0 {
        pos.Next = &ListNode {
			Val:  carry,
			Next: nil,
		} 
    }
    
    return head.Next
}


func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	sumNode := addTwoNumbers(l1, l2)
	
}
