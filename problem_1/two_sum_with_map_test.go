package problem_1

import "testing"

func TestTwoSumWithMap(t *testing.T) {
	result := []int{}

	result = twoSum([]int{2, 7, 11, 15}, 9)
	if !checkSliceEqual(result, []int{0, 1}) {
		t.Error("error")
	}

	result = twoSum([]int{2, 5, 5, 10}, 10)
	if !checkSliceEqual(result, []int{1, 2}) {
		t.Error("error")
	}
}