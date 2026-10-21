package main

type Foo struct {
	firstDone, secondDone chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		firstDone:  make(chan struct{}),
		secondDone: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.firstDone <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	<-f.firstDone
	/// Do not change this line
	printSecond()
	f.secondDone <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	<-f.secondDone
	// Do not change this line
	printThird()
}
