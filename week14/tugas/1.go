package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	count := 0
	for i := 1; 1 <= n; i++ {
		if i%2 != 0 {
			count++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil\n", count)
}