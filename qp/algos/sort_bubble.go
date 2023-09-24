package algos

import "fmt"

func Main_SortBubble() {
	/*
        complexity = n^2
        sort largest ele 1st

        at each time in inner for loop it is swapping the values i.e bubble
    */
	arr := []int{10, 5, 7, 2, 20, 5, 1}
	fmt.Println("before : ", arr)
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr) - 1; j++ {
				if arr[j] > arr[j+1] {
					temp := arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
				}
			}
		}

	fmt.Println("after : ", arr)
}