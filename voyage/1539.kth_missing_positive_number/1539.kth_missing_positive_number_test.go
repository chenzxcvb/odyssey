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
		{[]int{3, 4, 5}, 1, 1},
		{[]int{3, 4, 5, 9}, 3, 6},
		{[]int{1, 2, 3}, 2, 5},
	}
	for _, ca := range cases {
		actual := FindKthPositive(ca.arr, ca.k)
		if actual != ca.expected {
			t.Errorf("FindKthPositive(%v, %d) == %d, expected %d", ca.arr, ca.k, actual, ca.expected)
		}
	}

}
