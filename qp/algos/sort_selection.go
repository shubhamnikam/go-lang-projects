package algos

import "fmt"

func Main_SortSelection() {
	/*
        complexity = n^2
        sort smallest ele 1st

        in the end per loop only one time swap
    */
	arr := []int{10, 5, 7, 2, 20, 5, 1}
	fmt.Println("before : ", arr)

	for i := 0; i < len(arr); i++ {
		// in each iteration we are trying to find smallest elements index
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < minIndex {
				minIndex = j
			}
		}
		// swap min index with i
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}

	fmt.Println("after : ", arr)
}