package voyage

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
	missed := make([]int, 0)
	set := NewSetFromList(arr)
	for i, v := range arr {
		idx := i + 1
		for v < arr[idx+1] {
			if !set.Contains(idx) {
				missed = append(missed, v)
			}
			v++
		}
	}

	return missed[k-1]
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
