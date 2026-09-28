package main

import "fmt"

func main() {
	bytes := []byte("a")
	fmt.Println(bytes)
	fmt.Println(compress(bytes))
	fmt.Println(bytes)
}
