package main

import "fmt"

func main() {

	num := []int{9, 7, 8, 5, 6, 4, 52, 2}
	x := 54
	LinearSerch(num, x)

}
func LinearSerch(num []int, x int) {
	var i int
	for i = 0; i < len(num); i++ {
		if num[i] == x {
			fmt.Println("number is present at :- ", i, "th place")
			break
		}
	}
	if i == len(num) {
		fmt.Println("number not present....")
	}

}
