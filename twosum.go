func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for index, num := range nums {
		difference := target - num
		if res, ok := record[difference]; ok {
			return []int{index, res}
		}
		record[num] = index
	}

	return []int{}
}
