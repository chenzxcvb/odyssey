package voyage

/**
NOTE:
first try time complexity O(n), a solution found on youtube by "Fisher Coder"
another more efficient solution O(logn) is leverage binary tree search. TODO
*/

/**
Array -- EASY
https://leetcode.com/problems/kth-missing-positive-number/

Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.

Find the kth positive integer that is missing from this array.



Example 1:

Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
Example 2:

Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.


Constraints:

1 <= arr.length <= 1000
1 <= arr[i] <= 1000
1 <= k <= 1000
arr[i] < arr[j] for 1 <= i < j <= arr.length
*/

func FindKthPositive(arr []int, k int) int {
	// case 1: arr: [3, 4, 5], k: 1, result:1
	// case 2: arr: [3, 4, 5, 9], k: 3, result: 6
	// case 3: arr: [1, 2, 3], k: 2, result: 5
	missed := 0
	for i, v := range arr {
		if i == 0 {
			missed += v - 1
			if missed >= k { // case 1
				return k
			}
		} else {
			missed += arr[i] - arr[i-1] - 1
			if missed >= k { // case 2
				missed -= arr[i] - arr[i-1] - 1
				result := arr[i-1]
				for missed < k {
					missed++
					result++
				}
				return result
			}
		}
	}

	result := arr[len(arr)-1]
	for missed < k { // case3
		missed++
		result++
	}
	return result
}

type IntSet map[int]int

func NewSetFromList(ss []int) IntSet {
	result := make(IntSet)
	for idx, s := range ss {
		result[s] = idx
	}
	return result
}

func (s IntSet) Contains(v int) bool {
	_, ok := s[v]
	return ok
}
