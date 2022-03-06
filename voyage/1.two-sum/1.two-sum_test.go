package voyage

import "testing"

func TestTwoSum(t *testing.T) {
	var cases = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 18, []int{1, 2}},
		{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
	}
	for _, test := range cases {
		if actual := TwoSum(test.nums, test.target); !equal(actual, test.expected) {
			t.Errorf("TwoSum(%v, %v) == %v, expected %v", test.nums, test.target, actual, test.expected)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
