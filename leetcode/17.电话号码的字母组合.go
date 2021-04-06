package leetcode

var combinations []string
var mp map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	index := 0
	combination := ""
	combinations = []string{}
	backtrace(digits, index, combination)
	return combinations
}

func backtrace(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		letter := digits[index]
		for i := 0; i < len(mp[letter]); i++ {
			backtrace(digits, index+1, combination+string(mp[letter][i]))
		}
	}
}
