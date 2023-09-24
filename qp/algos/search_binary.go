package algos

import "fmt"

func Main_SearchBinary() {
	// NOTE : sorted input list needed
	arr := []int{1, 4, 6, 8, 9}
	key := 9

	found := false
	index := -1

	// binary sort uses divide and conquer approach
	low := 0
	high := len(arr) - 1
	mid := 0

	for low < high {
		mid = high / 2

		if arr[mid] < key {
			low = mid + 1
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			found = true
			index = mid
			break
		}
	}

	
	if found {
		fmt.Println("key found at index : ", index)
	} else {
		fmt.Println("key not found")
	}
}