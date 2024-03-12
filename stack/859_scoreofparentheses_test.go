package stack

import "testing"

func scoreOfParentheses(s string) int {
	var score = 0
	var multiplier = 1
	var stack []rune

	for _, c := range s {

		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		top := stack[len(stack)-1]
		if top == '(' && c == ')' {
			stack = stack[:len(stack)-1]
			score += multiplier
		}

		if top == '(' && c == '(' {
			multiplier *= 2
		}

		if top == ')' && c == ')' {
			multiplier /= 2
		}
		stack = append(stack, c)
	}
	return score
}

func TestScoreOfParentheses(t *testing.T) {
	println(scoreOfParentheses("()()()()()"))
	println(scoreOfParentheses("((()))()(())"))
	println(scoreOfParentheses("(()(()))"))
}
