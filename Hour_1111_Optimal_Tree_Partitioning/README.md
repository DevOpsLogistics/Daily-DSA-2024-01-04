# Optimal Tree Partitioning

**Language:** Go | **Date:** 2024-01-04 11:11

## Description
You are given an undirected tree with n nodes, indexed from 0 to n-1. Each node i has an associated weight weights[i]. Your task is to partition the tree into exactly k connected components by removing k-1 edges from the tree. The objective is to minimize the maximum weight among all resulting components. The weight of a component is the sum of the weights of all nodes it contains.

Input:
- n: The number of nodes in the tree.
- k: The number of components required.
- weights: An array of n integers representing the weight of each node.
- edges: A 2D array where edges[i] = [u, v] representing an edge between nodes u and v.

Output: The smallest possible maximum weight of any component.

Example:
Input: n = 4, k = 3, weights = [10, 20, 30, 40], edges = [[0, 1], [1, 2], [1, 3]]
Output: 40
Explanation: By removing edges (1, 2) and (1, 3), we obtain three components: {0, 1} with weight 30, {2} with weight 30, and {3} with weight 40. The maximum weight is 40. No other partition yields a smaller maximum weight.

## Explanation
The problem is solved using a combination of Binary Search on the Answer and a Greedy DFS strategy on the tree. 

1. Binary Search: The range of possible answers lies between the maximum weight of a single node (the smallest possible max weight) and the sum of all node weights (the max weight when k=1). 
2. Greedy Strategy: For a fixed maximum weight limit 'mid', we use DFS to traverse the tree from the leaves upwards. For each node, we accumulate the weight of its subtree. If the sum of the current node's weight and a child's subtree weight exceeds 'mid', we greedily cut the edge connecting to that child. This is optimal because cutting an edge higher up in the tree (closer to the root) never performs worse than cutting it lower down for minimizing component sums.
3. Feasibility: If the number of cuts performed is less than or equal to k-1, then 'mid' is a feasible maximum weight, and we attempt to find a smaller value in the lower half of our binary search range.

Complexity:
- Time Complexity: O(N log(TotalSum)), where N is the number of nodes and TotalSum is the sum of all node weights.
- Space Complexity: O(N) to store the adjacency list and the recursion stack for DFS.
