package algos

import "fmt"

func Main_SearchLinear() {
	arr := []int{10, 5, 7, 2, 20, 5, 1}
	key := 11

	found := false
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			found = true
			index = i
			break
		}
	}

	if found {
		fmt.Println("key found at index : ", index)
	} else {
		fmt.Println("key not found")
	}
}