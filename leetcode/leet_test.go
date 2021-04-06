package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var nums []int

func init() {
	nums = []int{1,1,1,1,2,2,2,2,2,3,3}
}

func TestRemoveDuplicates(t *testing.T) {
	l := RemoveDuplicates(nums)
	t.Logf("nums:%v l:%v\n", nums, l)
}


func TestRemoveRepeatK1(t *testing.T) {
	l := RemoveRepeatK(nums, 1)
	t.Logf("nums:%v l:%v\n", nums, l)
	assert.Equal(t, nums[0:l], []int{1, 2, 3})
}

func TestRemoveRepeatK2(t *testing.T) {
	l := RemoveRepeatK(nums, 2)
	t.Logf("nums:%v l:%v\n", nums, l)
	assert.Equal(t, nums[0:l], []int{1, 1, 2, 2, 3, 3})
}

func TestRemoveRepeatK3(t *testing.T) {
	l := RemoveRepeatK(nums, 3)
	t.Logf("nums:%v l:%v\n", nums, l)
	assert.Equal(t, nums[0:l], []int{1, 1, 1, 2, 2, 2, 3, 3})
}
func TestRemoveRepeatK4(t *testing.T) {
	l := RemoveRepeatK(nums, 4)
	t.Logf("nums:%v l:%v\n", nums, l)
	assert.Equal(t, nums[0:l], []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3})
}
