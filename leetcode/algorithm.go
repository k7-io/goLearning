package leetcode

/*
3的幂
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
*/
func isPowerOfThree(n int) bool {
	/**循环
	if n < 1 {
	    return false
	}
	for n % 3 == 0 {
	    n = n / 3
	}
	return n == 1**/
	//3**19 % 3的倍数 == 0
	if n < 1 {
		return false
	}
	return 1162261467%n == 0
}

/*
位1的个数
编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
*/
func hammingWeight(num uint32) int {
	ans := 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		mask = 1 << i
		if num&mask != 0 {
			ans++
		}
	}
	return ans
}

/*
有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
*/
func isValid(s string) bool {
	stack := []uint8{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			if len(stack) != 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if s[i] == ']' {
			if len(stack) != 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			if len(stack) != 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

/*
螺旋矩阵 II
*/
func generateMatrix(n int) [][]int {
	var (
		left, right, top, bottom = 0, n - 1, 0, n - 1
		ans                      [][]int
		index                    = 1
	)
	for i := 0; i < n; i++ {
		inline := make([]int, n)
		ans = append(ans, inline)
	}
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			ans[top][column] = index
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			ans[row][right] = index
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				ans[bottom][column] = index
				index++
			}
			for row := bottom; row > left; row-- {
				ans[row][left] = index
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return ans
}
