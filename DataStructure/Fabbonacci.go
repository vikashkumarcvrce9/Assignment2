//programme to print nth Fabbonaci
package main

import "fmt"

func main() {
	var i = fabbonaci(8)
	fmt.Println(i)
}
func fabbonaci(n int) int {

	for n <= 1 {
		return n
	}
	return fabbonaci(n-1) + fabbonaci(n-2)
}
