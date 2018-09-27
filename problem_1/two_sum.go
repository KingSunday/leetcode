package problem_1

import "sort"

func twoSum(numList []int, target int) []int {
	length := len(numList)
	copNumList := make([]int, length)
	copy(copNumList, numList)
	sort.Ints(copNumList)

	sub := 0
	for i, vi := range numList {
		sub = target - vi
		index := sort.SearchInts(copNumList, sub)
		if index == length || copNumList[index] != sub {
			continue
		}
		for j, vj := range numList {
			if vj == copNumList[index] && i != j {
				return []int{i, j}
			}
		}
	}
	panic("illegal input")
}