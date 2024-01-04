package main

import (
	"fmt"
)

// solve calculates the minimum possible maximum weight of any component after k-1 cuts.
func solve(n int, k int, weights []int, edges [][]int) int64 {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var maxW int64
	var sumW int64
	for _, w := range weights {
		if int64(w) > maxW {
			maxW = int64(w)
		}
		sumW += int64(w)
	}

	// Binary search for the optimal maximum weight
	low, high := maxW, sumW
	ans := sumW

	for low <= high {
		mid := low + (high-low)/2
		cuts := 0

		// Greedy DFS to count necessary cuts for a given limit 'mid'
		var dfs func(u, p int) int64
		dfs = func(u, p int) int64 {
			currentSum := int64(weights[u])
			for _, v := range adj[u] {
				if v == p {
					continue
				}
				childSum := dfs(v, u)
				// If adding the child subtree exceeds mid, we must cut the edge to this child
				if currentSum+childSum > mid {
					cuts++
				} else {
					currentSum += childSum
				}
			}
			return currentSum
		}

		dfs(0, -1)

		// If the number of cuts is within k-1, we can potentially lower the max weight
		if cuts <= k-1 {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func main() {
	n := 4
	k := 3
	weights := []int{10, 20, 30, 40}
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}}
	fmt.Printf("Result: %d\n", solve(n, k, weights, edges))
}