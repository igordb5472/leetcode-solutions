package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(0))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(0))
}
