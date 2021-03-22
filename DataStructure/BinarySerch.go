package main

import (
	"fmt"
	"sort"
)

func main() {

	num := []int{9, 7, 8, 5, 6, 4, 52, 2}
	sort.Ints(num)
	key := 11
	BinarySerch(num, key)

}

func BinarySerch(num []int, key int) {
	low := 0
	high := len(num)
	for low < high {

		mid := (low + high) / 2
		if key < num[mid] {
			high = mid - 1
		} else if key > num[mid] {
			low = mid + 1
		} else if num[mid] == key {
			fmt.Println("number is present at index:-", mid)
			break
		}

	}
	if high < low {
		fmt.Println("number not present")
	}

}
