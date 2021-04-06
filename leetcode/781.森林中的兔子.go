package leetcode

import "sort"

func numRabbits(answers []int) int {
	size := len(answers)
	if size == 0 {
		return 0
	}
	sort.Ints(answers)
	ans := 0
	leastNum := answers[0] + 1
	count := 1
	answer := answers[0]
	for i := 1; i < size; i++ {
		if answers[i] == answer && leastNum > count {
			count++
		} else {
			ans += leastNum
			leastNum = answers[i] + 1
			count = 1
			answer = answers[i]
		}
	}
	ans += leastNum
	return ans
}
