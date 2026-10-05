package main

type stack struct {
	values []int
}

func (this *stack) push(x int) {
	this.values = append(this.values, x)
}

func (this *stack) peek() int {
	return this.values[len(this.values)-1]
}

func (this *stack) pop() int {
	top := this.values[len(this.values)-1]
	this.values = this.values[:len(this.values)-1]
	return top
}

func (this *stack) size() int {
	return len(this.values)
}

func (this *stack) isEmpty() bool {
	return len(this.values) == 0
}

type MyQueue struct {
	st stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.st.push(x)
}

func (this *MyQueue) Pop() int {
	var cache stack
	for !this.st.isEmpty() {
		cache.push(this.st.pop())
	}
	peek := cache.pop()
	for !cache.isEmpty() {
		this.st.push(cache.pop())
	}
	return peek
}

func (this *MyQueue) Peek() int {
	var cache stack
	for !this.st.isEmpty() {
		cache.push(this.st.pop())
	}
	peek := cache.peek()
	for !cache.isEmpty() {
		this.st.push(cache.pop())
	}
	return peek
}

func (this *MyQueue) Empty() bool {
	return this.st.isEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
