package main

import "fmt"

func main() {
	var i = GCD(32, 16)
	fmt.Println(i)
}
func GCD(m int, n int) int {
	if n > m {
		GCD(n, m)
	}
	if m%n == 0 {
		return n
	}
	return GCD(n, m%n)
}
