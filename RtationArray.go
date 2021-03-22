package main

import "fmt"

func RotateArray(num []int, k int, n int) {
	temp := num[0]
	for i := 0; i <= k; i++ {
		for j := 0; j < n; j++ {
			num[j] = num[j+1]
			num[n-1] = temp

		}

	}
	for i := 0; i < n; i++ {
		fmt.Print(num[i], " ")
	}
}
func main() {
	numb := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	n := len(numb)
	k := 4
	RotateArray(numb, k, n)
}
