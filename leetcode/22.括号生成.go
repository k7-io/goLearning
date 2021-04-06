package leetcode

var (
	parenthesis string
	ans         []string
)

func generateParenthesis(n int) []string {
	parenthesis = ""
	ans = []string{}
	backstrace(parenthesis, n, n)
	return ans
}

func backstrace(parenthesis string, left int, right int) {
	if left == 0 && right == 0 {
		ans = append(ans, parenthesis)
	}
	if left > 0 {
		parenthesis += "("
		backstrace(parenthesis, left-1, right)
		parenthesis = parenthesis[:len(parenthesis)-1]
	}
	if right > left {
		parenthesis += ")"
		backstrace(parenthesis, left, right-1)
	}
}
