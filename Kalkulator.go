package main

import "fmt"

func main() {
	var array = []int{1, 10, 45, 67, 33, 99, 2, 8}
	min, max, average := calculatorStatistik(array)
	fmt.Println("Terbesar : ", max)
	fmt.Println("Terkecil : ", min)
	fmt.Println("rata rata : ", average)

}

func calculatorStatistik(array []int) (min, max, average int) {
	min = array[0]
	max = array[0]
	var total int
	for _, value := range array {
		total += value
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		average = total / len(array)
	}
	return min, max, average
}
