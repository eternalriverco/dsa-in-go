package stack

import (
	"testing"
)

/*
https://leetcode.com/problems/min-stack/description/

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

    MinStack() initializes the stack object.
    void push(int val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.

Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2



Constraints:

    -231 <= val <= 231 - 1
    Methods pop, top and getMin operations will always be called on non-empty stacks.
    At most 3 * 104 calls will be made to push, pop, top, and getMin.

*/

type Value struct {
	min int
	v   int
}

type MinStack struct {
	s []Value
}

func Constructor() MinStack {
	return MinStack{
		s: []Value{},
	}
}

func (this *MinStack) Push(val int) {
	v := Value{
		min: val,
		v:   val,
	}

	if len(this.s) == 0 {
		this.s = append(this.s, v)
		return
	}

	top := this.s[len(this.s)-1]
	if val < top.min {
		v.min = val
	} else {
		v.min = top.min
	}
	this.s = append(this.s, v)
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1].v
}

func (this *MinStack) GetMin() int {
	return this.s[len(this.s)-1].min
}

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	println(ms.GetMin())
	ms.Pop()
	println(ms.Top())
	println(ms.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
