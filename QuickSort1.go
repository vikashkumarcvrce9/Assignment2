package main

import (
	"fmt"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		var Pivot = partition(arr, low, high)
		quickSort(arr, low, Pivot)
		quickSort(arr, Pivot+1, high)

	}

}
func partition(arr []int, low, high int) int {
	var Pivot = arr[low]
	var i = low
	var j = high
	for i < j {
		for arr[i] <= Pivot && i < high {
			i++

		}
		for arr[i] > Pivot && j > low {
			j--
		}
		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	arr[low] = arr[j]
	arr[j] = Pivot
	return j
}
func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")

	}
	fmt.Println()
}
func main() {
	var arr = []int{110, 16, 11, 14, 8, 9, 7}
	fmt.Print("Before Sorting: ")
	printArray(arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Print("After Sorting:")
	printArray(arr)

}
