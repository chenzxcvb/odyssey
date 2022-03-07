package voyage

import (
	"testing"
)

func TestFindKthPositive(t *testing.T) {
	var cases = []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{7, 11}, 2, 5},
		{[]int{1, 2, 3, 4}, 2, 6},
		{[]int{2, 3, 4, 7, 11}, 5, 9},
	}
	for _, ca := range cases {
		actual := FindKthPositive(ca.arr, ca.k)
		if actual != ca.expected {
			t.Errorf("FindKthPositive(%v, %d) == %d, expected %d", ca.arr, ca.k, actual, ca.expected)
		}
	}

}
