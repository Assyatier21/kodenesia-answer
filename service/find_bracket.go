package service

import "reflect"

func FindBrackets(s string) string {
	var (
		bracketOne   = []bool{}
		bracketTwo   = []bool{}
		bracketThree = []bool{}
	)

	if len(s) < 1 || len(s) > 1000 {
		return "TIDAK"
	}

	for _, char := range s {
		switch char {
		case '{':
			bracketOne = append(bracketOne, true)
		case '}':
			bracketOne = append(bracketOne, true)
		case '[':
			bracketTwo = append(bracketTwo, true)
		case ']':
			bracketTwo = append(bracketTwo, true)
		case '(':
			bracketThree = append(bracketThree, true)
		case ')':
			bracketThree = append(bracketThree, true)
		}
	}

	checker := map[string][]bool{
		"{}": bracketOne,
		"[]": bracketTwo,
		"()": bracketThree,
	}

	if reflect.DeepEqual(checker["{}"], []bool{true, true}) || reflect.DeepEqual(checker["[]"], []bool{true, true}) || reflect.DeepEqual(checker["()"], []bool{true, true}) {
		return "YA"
	}

	return "TIDAK"
}
