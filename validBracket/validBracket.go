package main

func judgeValidBracket(str string) bool {
	top := 0
	stack := make([]string, len(str))

	for s := range str {
		switch str[s] {
		case '(', '{', '[':
			stack[top] = string(str[s])
			top++
		case ')', '}', ']':
			if top == 0 {
				return false
			}
			top--
			if (str[s] == ')' && stack[top] != "(") ||
				(str[s] == '}' && stack[top] != "{") ||
				(str[s] == ']' && stack[top] != "[") {
				return false
			}
		}
	}

	return top == 0
}

func main() {
	str := "{[(fsdf)(ttt)b]c}"
	if judgeValidBracket(str) {
		println("The string is a valid bracket sequence.")
	} else {
		println("The string is not a valid bracket sequence.")
	}
}
