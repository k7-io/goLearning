package leetcode

func trap(height []int) int {
	size := len(height)
	if size <= 2 {
		return 0
	}
	level := 1
	left := 0
	right := size - 1
	totalArea := 0
	columnarArea := 0
	for left <= right {
		for left <= right && height[left] < level {
			left++
		}
		for left <= right && height[right] < level {
			right--
		}
		totalArea += (right - left + 1)
		level++
	}
	for _, v := range height {
		if v != 0 {
			columnarArea += v
		}
	}
	return totalArea - columnarArea
}
