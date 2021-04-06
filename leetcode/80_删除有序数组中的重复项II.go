package leetcode

func removeDuplicatesV1(nums []int) int {
	// l := 0
	// for i:=1; i<len(nums); i++ {
	//     if nums[i] != nums[l] {
	//         l += 1
	//         nums[l] = nums[i]
	//     }
	// }
	// // fmt.Printf("%v:%v", l+1, nums)
	// return l + 1
	l, f := 0, 1
	n := len(nums)
	for f < n {
		if nums[f] != nums[l] {
			l ++
			nums[l] = nums[f]
		}
		f ++
	}
	return l+1
}

/*
* nums: 有序数组
* 移除重复出现的元素，最多保留2个
*/
func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return 2
	}

	l, f := 2, 2
	for f < n {
		// nums[f]: 当前需要检测的值
		// nums[l]: 需要被填充的位置
		// nums[l-1]: 第二次出现
		// nums[l-2]: 第一次出现
		if nums[f] != nums[l-2] {
			nums[l] = nums[f]
			l++
		}
		f++
	}
	return l
}

/*
为了让解法更具有一般性，我们将原问题的「保留 2 位」修改为「保留 k 位」。
对于此类问题，我们应该进行如下考虑：
由于是保留 k 个相同数字，对于前 k 个数字，我们可以直接保留
对于后面的任意数字，能够保留的前提是：与当前写入的位置前面的第 k 个元素进行比较，不相同则保留
作者：AC_OIer
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/solution/gong-shui-san-xie-guan-yu-shan-chu-you-x-glnq/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
[1,1,1,2,2,3]
   l,f
*/
func RemoveRepeatK(nums []int, k int) int {
	n := len(nums)
	if n <= k {
		return n
	}

	l, f := k, k
	for f < n {
		// nums[f]: 当前需要检测的值
		// nums[l]: 需要被填充的位置
		// nums[l-k] 上一个保留的字符
		if nums[f] != nums[l-k] {
			nums[l] = nums[f]
			l++
		}
		f++
	}
	return l
}