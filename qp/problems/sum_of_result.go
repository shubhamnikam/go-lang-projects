package problems

import "fmt"

func Main_SumOfResult() {
	inputArr := []int{1, 2, 3, 4, 5, 6}
	resultSum := 5
	resultMap := calc(inputArr, resultSum)
	fmt.Println(resultMap)

}

func calc(inputArr []int, resultSum int) (map[int][2]int) {
	resultMap := map[int][2]int{}

	for i:=0; i<len(inputArr); i++ {
		for j:=i+1; j<len(inputArr); j++ {
			if inputArr[i] + inputArr[j] == resultSum {
				resultMap[len(resultMap)+1] = [2]int{inputArr[i], inputArr[j]}
			}
		}
	}
	return resultMap
}