package main

type ZeroEvenOdd struct {
	n               int
	zero, even, odd chan int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zero, even, odd := make(chan int, 1), make(chan int, 1), make(chan int, 1)
	zero <- 1
	zeo := &ZeroEvenOdd{
		n:    n,
		zero: zero,
		even: even,
		odd:  odd,
	}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for cur := range z.zero {
		printNumber(0)
		if cur%2 == 1 {
			z.odd <- cur
		} else {
			z.even <- cur
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for cur := range z.even {
		printNumber(cur)
		if cur == z.n {
			close(z.zero)
			close(z.even)
			close(z.odd)
			return
		}
		z.zero <- cur + 1
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for cur := range z.odd {
		printNumber(cur)
		if cur == z.n {
			close(z.zero)
			close(z.even)
			close(z.odd)
			return
		}
		z.zero <- cur + 1
	}
}
