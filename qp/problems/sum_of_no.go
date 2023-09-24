package problems

import "fmt"

func Main_SumOfNo() {
	arr := []int{1, 2, 3, 4, 5}

	way1(arr)
	way2(arr...)
}

func way1(arr []int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}

func way2(arr ...int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}
