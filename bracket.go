package bracket

import (
	"strings"
)

func Match2(input string) bool {
	for (strings.Index(input, "{}") >= 0) || (strings.Index(input, "[]") >= 0) || (strings.Index(input, "()") >= 0) {
		input = strings.Replace(input, "[]", "", -1)
		input = strings.Replace(input, "{}", "", -1)
		input = strings.Replace(input, "()", "", -1)
	}

	return len(input) == 0
}

func Match(input string) bool {
	stack := []rune{}
	bracketPair := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	for _, c := range input {
		if c == '[' || c == '{' || c == '(' {
			stack = append(stack, c)
		} else {
			if c == ']' || c == '}' || c == ')' {
				if len(stack) > 0 {
					top := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if bracketPair[c] != top {
						return false
					}
				} else {
					return false
				}
			}
		}
	}

	return (len(stack) == 0)
}
