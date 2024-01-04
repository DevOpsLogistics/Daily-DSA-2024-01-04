package main

import (
	"fmt"
)

type TrieNode struct {
	children [2]*TrieNode
}

func insert(root *TrieNode, val int) {
	curr := root
	for i := 30; i >= 0; i-- {
		bit := (val >> i) & 1
		if curr.children[bit] == nil {
			curr.children[bit] = &TrieNode{}
		}
		curr = curr.children[bit]
	}
}

func query(root *TrieNode, val int) int {
	curr := root
	res := 0
	for i := 30; i >= 0; i-- {
		bit := (val >> i) & 1
		target := 1 - bit
		if curr.children[target] != nil {
			res |= (1 << i)
			curr = curr.children[target]
		} else if curr.children[bit] != nil {
			curr = curr.children[bit]
		} else {
			return 0
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximizeXorSum(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// Forward pass
	lRoot := &TrieNode{}
	insert(lRoot, 0)
	prefixXor := 0
	currentMax := 0
	for i := 0; i < n; i++ {
		prefixXor ^= nums[i]
		currentMax = max(currentMax, query(lRoot, prefixXor))
		insert(lRoot, prefixXor)
		leftMax[i] = currentMax
	}

	// Backward pass
	rRoot := &TrieNode{}
	insert(rRoot, 0)
	suffixXor := 0
	currentMax = 0
	for i := n - 1; i >= 0; i-- {
		suffixXor ^= nums[i]
		currentMax = max(currentMax, query(rRoot, suffixXor))
		insert(rRoot, suffixXor)
		rightMax[i] = currentMax
	}

	maxSum := 0
	for i := 0; i < n-1; i++ {
		maxSum = max(maxSum, leftMax[i]+rightMax[i+1])
	}

	return maxSum
}

func main() {
	nums := []int{1, 2, 3, 1, 2}
	fmt.Println(maximizeXorSum(nums)) // Output: 6
}