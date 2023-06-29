package leetcode

// The worst case of the time complexity is 2N,
// and the space complexity of the worst is 2N.
func TwoSum(nums []int, target int) []int {
	// to add up by two numbers, we can think of it reversely from the target.
	// so when get first one number, using target minus first one number, then check if the outcome number is existing.
	// in order to speed up costing time, we will redefind data structure as map tpye.
	m := make(map[int]int)
	for _, v := range nums {
		m[v] += 1
	}
	// now on, we have the map which can make the cost close to 1 when retrieving data.
	answer := []int{}
	for index, v := range nums {
		m[v] -= 1
		num := target - v
		if b, ok := m[num]; ok && b > 0 {
			answer = append(answer, index)
		}
		m[v] += 1
	}
	return answer
}
