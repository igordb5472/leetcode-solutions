package main

type FooBar struct {
	n          int
	printFooCh chan struct{}
	printBarCh chan struct{}
}

func NewFooBar(n int) *FooBar {
	printFooCh, printBarCh := make(chan struct{}, 1), make(chan struct{}, 1)
	printFooCh <- struct{}{}
	return &FooBar{
		n:          n,
		printFooCh: printFooCh,
		printBarCh: printBarCh,
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.printFooCh
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.printBarCh <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.printBarCh
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.printFooCh <- struct{}{}
	}
}
