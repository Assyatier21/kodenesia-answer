package service

func FindBrackets(s string) string {
	stack := make([]rune, 0)
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 {
				return "TIDAK"
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if mapping[char] != top {
				return "TIDAK"
			}
		}
	}

	if len(stack) == 0 {
		return "YA"
	} else {
		return "TIDAK"
	}
}
