package main

import (
	"fmt"
	"sync"
)

func main() {
	zeo := NewZeroEvenOdd(5)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		zeo.Zero(func(n int) { fmt.Print(n) })
		wg.Done()
	}()
	go func() {
		zeo.Even(func(n int) { fmt.Print(n) })
		wg.Done()
	}()
	go func() {
		zeo.Odd(func(n int) { fmt.Print(n) })
		wg.Done()
	}()
	wg.Wait()
}
