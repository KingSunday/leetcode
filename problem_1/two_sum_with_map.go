package problem_1

func twoSumWithMap(numList []int, target int) []int {
	numMap := make(map[int]int)

	for i, v := range numList{
		if j, ok := numMap[target - v]; ok {
			return []int{j, i}
		}
		numMap[v] = i
	}
	panic("illegal input")
}