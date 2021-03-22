package main

import "fmt"

func permutation(data []int, i int, length int) {
	if length == i {
		fmt.Println(data)

	} else {
		for j := i; j < length; j++ {
			swap(data, i, j)
			permutation(data, i+1, length)
			swap(data, i, j)
		}

	}

}
func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]

}

func main() {

	var data = []int{0, 2, 3}
	permutation(data[:], 0, len(data))
}
