package algos

import "fmt"

func Main_SortInsertion() {
	arr := []int{10, 5, 7, 2, 20, 5, 1}
	fmt.Println("before : ", arr)

    for i:= 1; i< len(arr); i++ {
        key:= arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }

	fmt.Println("after : ", arr)
}