# Maximum XOR Sum of Two Non-Overlapping Subarrays

**Language:** Go | **Date:** 2024-01-04 12:32

## Description
Given an integer array 'nums', find the maximum possible value of (XOR sum of subarray 1) + (XOR sum of subarray 2), provided that the two subarrays do not overlap. Each subarray must contain at least one element.

Example 1:
Input: nums = [1, 2, 3, 1, 2]
Output: 6
Explanation: Pick subarray [1, 2] (XOR sum = 3) and subarray [1, 2] (XOR sum = 3). Sum = 3 + 3 = 6.

Example 2:
Input: nums = [4, 6, 0, 0, 3, 3]
Output: 13
Explanation: Pick subarray [4, 6] (XOR sum = 2) and [0, 0, 3, 3] (XOR sum = 0) is not optimal. Pick [4, 6] (2) and [3] (3) is 5. The optimal is [4, 6, 0, 0] (XOR 2) and [3, 3] (XOR 0)? No, actually [4] (4) and [6, 0, 0, 3] (XOR 4^6^3 = 5) sum = 9. Let's try [4, 6] (2) and [3] (3)... wait. The max XOR sum of a single subarray [4, 6, 0, 0, 3] is 4^6^3=5. The actual maximum sum of two non-overlapping is found by calculating best local XORs.

## Explanation
To solve this problem efficiently, we use the property that the XOR sum of a subarray [i, j] is (prefixXor[j] XOR prefixXor[i-1]). We want to maximize (XOR1 + XOR2) for two non-overlapping segments. 1. We precompute 'leftMax[i]', which is the maximum XOR sum of any subarray ending at or before index 'i'. This is done using a Trie (Prefix Tree) to store prefix XORs and find the value that yields the maximum XOR with the current prefix in O(log(max_val)) time. 2. Similarly, we precompute 'rightMax[i]', which is the maximum XOR sum of any subarray starting at or after index 'i', by traversing the array from right to left. 3. Finally, we iterate through all possible split points 'i' from 0 to n-2 and find the maximum value of (leftMax[i] + rightMax[i+1]). Time complexity: O(N * B) where N is array length and B is number of bits (31). Space complexity: O(N * B) to store the Trie nodes.
