package stack

import "testing"

/*
https://leetcode.com/problems/valid-parentheses/description/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false

Constraints:

    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.

*/

func isValid(s string) bool {
	if len(s) == 1 || len(s)%2 != 0 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []rune

	for _, c := range s {
		if _, ok := pairs[c]; ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 || pairs[stack[len(stack)-1]] != c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func TestValidParentheses(t *testing.T) {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("(])"))
	println(isValid("([{}])"))
	println(isValid("]](("))
}
