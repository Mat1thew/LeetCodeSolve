package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func totalHammingDistance(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			result += hammingDistance(nums[i], nums[j])
		}
	}
	return result
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	tempList := &ListNode{Next: head}
	for tmpNode := tempList.Next; tmpNode.Next != nil; {
		if tmpNode.Val == val {
			tempList.Next = tempList.Next.Next
		} else {
			tempList = tempList.Next
		}
	}
	return head
}
func main() {
	a := totalHammingDistance([]int{1, 2, 3})
	fmt.Println(a)
}
