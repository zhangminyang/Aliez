package main

import (
	"errors"
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	n := search(nums, 0)
	//mid := size / 2
	fmt.Printf("%v", n)
}

//插入排序
func InsertSort(A []int) {
	var tmp, j int
	N := len(A)
	for i := 0; i < N; i++ {
		tmp = A[i]
		for j = i; j > 0 && A[j-1] > tmp; j-- {
			A[j] = A[j-1]
		}
		A[j] = tmp
	}
}

//归并排序
func MergeSort(A []int) {
	//分配额外空间
	N := len(A)
	tmp := make([]int, N)
	if tmp != nil {
		MSort(A, tmp, 0, N-1)
	} else {
		errors.New("No space for tmp array!")
	}

}

func MSort(A []int, tmp []int, Left int, Right int) {
	var Center int
	if Left < Right {
		//拆分左右数组 递归进行排序
		Center = (Left + Right) / 2
		MSort(A, tmp, Left, Center)
		MSort(A, tmp, Center+1, Right)
		//合并左右数组
		Merge(A, tmp, Left, Center+1, Right)
	}
}

func Merge(A []int, TmpArray []int, Lptr int, Rptr int, REnd int) {
	LEnd := Rptr - 1
	TmpPos := Lptr
	NumElements := REnd - Lptr + 1

	//
	for Lptr <= LEnd && Rptr <= REnd {
		if A[Lptr] <= A[Rptr] {
			TmpArray[TmpPos] = A[Lptr]
			TmpPos++
			Lptr++
		} else {
			TmpArray[TmpPos] = A[Rptr]
			TmpPos++
			Rptr++
		}
	}
	//
	for Lptr <= LEnd {
		TmpArray[TmpPos] = A[Lptr]
		TmpPos++
		Lptr++
	}
	//
	for Rptr <= REnd {
		TmpArray[TmpPos] = A[Rptr]
		TmpPos++
		Rptr++
	}

	//重新填入A数组
	for i := 0; i < NumElements; i, REnd = i+1, REnd-1 {
		A[REnd] = TmpArray[REnd]
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	r := len(lists)
	return msortList(lists, 0, r-1)
}

func msortList(lists []*ListNode, l int, r int) *ListNode {
	if l > r {
		return nil
	}
	if l == r {
		return lists[l]
	}
	mid := (l + r) >> 1
	return mergeTwoLists(msortList(lists, l, mid), msortList(lists, mid+1, r))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prevhead := new(ListNode)
	prevhead.Val = -1
	prev := prevhead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil && l2 != nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return prevhead.Next
}

//leetcode HOT 33
func search(nums []int, target int) int {
	size := len(nums)
	if size == 1 && nums[0] == target {
		return 0
	} else if size == 1 && nums[0] != target {
		return -1
	}
	var mid int
	l, r := 0, size-1
	for l <= r {
		mid = (l + r) / 2
		fmt.Printf("%v", mid)
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target < nums[mid] && target >= nums[0] {
				//在前半段查找
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[size-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

//leetcode HOT 139
func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	dp := make([]bool, size+1)
	dp[0] = true
	word := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		word[v] = true
	}
	for i := 1; i <= size; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && word[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[size]
}

//leetcode HOT 102
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[i]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, ans)
		queue = queue[l:]
	}
	return ret
}

func levelOrderDFS(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	var SerchRoot func(root *TreeNode, index int)
	SerchRoot = func(root *TreeNode, index int) {
		if root == nil {
			return
		}
		if len(ret) == index {
			ret = append(ret, []int{})
		}
		ret[index] = append(ret[index], root.Val)
		SerchRoot(root.Left, index+1)
		SerchRoot(root.Right, index+1)
	}
	SerchRoot(root, 0)
	return ret
}

// leetcode 257
func binaryTreePaths(root *TreeNode) []string {
	path := make([]string, 0)
	if root == nil {
		return path
	}
	var DFS func(root *TreeNode, str string)
	DFS = func(root *TreeNode, str string) {
		if root == nil {
			return
		}
		str = fmt.Sprintf("%s%d->", str, root.Val)
		if root.Left == nil && root.Right == nil {
			path = append(path, str[:len(str)-2])
		}
		DFS(root.Left, str)
		DFS(root.Right, str)
	}
	DFS(root, "")
	return path
}

//leetcode 107
func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[i]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, ans)
		queue = queue[l:]
	}
	//反转数组
	l := len(ret)
	for i := 0; i < l/2; i++ {
		tmp := ret[i]
		ret[i] = ret[l-i-1]
		ret[l-i-1] = tmp
	}
	return ret
}

//leetcode 144
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	var PT func(root *TreeNode)
	PT = func(root *TreeNode) {
		ret = append(ret, root.Val)
		if root.Left != nil {
			PT(root.Left)
		}
		if root.Right != nil {
			PT(root.Right)
		}
	}
	PT(root)
	return ret
}

func preorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ret
}

//leetcode 114
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var tmp *TreeNode
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmp != nil {
			tmp.Right, tmp.Left = nil, node
		}
		left, right := tmp.Left, tmp.Right
		if right != nil {
			stack = append(stack, node.Right)
		}
		if left != nil {
			stack = append(stack, node.Left)
		}
		tmp = node
	}
}

func flatten3(root *TreeNode) {
	if root == nil {
		return
	}
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			per := next
			for per.Right != nil {
				per = per.Right
			}
			per.Right = curr.Right
			curr.Left = nil
			curr.Right = next
		}
		curr = curr.Right
	}
}
