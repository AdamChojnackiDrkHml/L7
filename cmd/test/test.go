package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 16000000; i++ {
		fmt.Println(i)
	}
	fmt.Println(a)
}
