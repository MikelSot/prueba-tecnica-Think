package main

func removeZeroSum(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}

	sum := 0

	sumMap := make(map[int]int)
	sumMap[0] = -1

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if j, ok := sumMap[sum]; ok {
			arr = append(arr[:j+1], arr[i+1:]...)
			i = j
		} else {
			sumMap[sum] = i
		}
	}

	return arr
}
