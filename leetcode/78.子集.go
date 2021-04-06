package leetcode

func subsets(nums []int) [][]int {
	ans := [][]int{}
	setSlice := []int{}
	var backtrace func(cur int)
	backtrace = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int{}, setSlice...))
			return
		}
		setSlice = append(setSlice, nums[cur])
		backtrace(cur + 1)
		setSlice = setSlice[:len(setSlice)-1]
		backtrace(cur + 1)
	}
	backtrace(0)
	return ans
}
