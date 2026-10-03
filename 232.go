package main

type stack[T any] struct {
	values []T
}

func (this *stack[T]) push(x T) {
	this.values = append(this.values, x)
}

func (this *stack[T]) peek() T {
	return this.values[len(this.values)-1]
}

func (this *stack[T]) pop() T {
	top := this.values[len(this.values)-1]
	this.values = this.values[:len(this.values)-1]
	return top
}

func (this *stack[T]) size() int {
	return len(this.values)
}

func (this *stack[T]) isEmpty() bool {
	return len(this.values) == 0
}

type MyQueue struct {
}

func Constructor() MyQueue {

}

func (this *MyQueue) Push(x int) {

}

func (this *MyQueue) Pop() int {

}

func (this *MyQueue) Peek() int {

}

func (this *MyQueue) Empty() bool {

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
