package main

import "sync"

type DiningPhilosophers struct {
	once  sync.Once
	forks [5]chan struct{}
	limit chan struct{}
}

func (this *DiningPhilosophers) wantsToEat(
	philosopher int,
	pickLeftFork func(),
	pickRightFork func(),
	eat func(),
	putLeftFork func(),
	putRightFork func(),
) {
	// TODO: implement your solution here
	this.once.Do(func() {
		for i := range 5 {
			fork := make(chan struct{}, 1)
			fork <- struct{}{}
			this.forks[i] = fork
		}
		this.limit = make(chan struct{}, 4)
	})
	this.limit <- struct{}{}
	l, r := this.forks[(philosopher+1)%5], this.forks[philosopher]
	<-l
	pickLeftFork()
	<-r
	pickRightFork()
	eat()
	putLeftFork()
	l <- struct{}{}
	putRightFork()
	r <- struct{}{}
	<-this.limit
}
