package main

import (
	"fmt"
)

var yourSlice = []int{1, 222, 66, 44, 22, 2, 6, 4, 5, 10, 9, 23, 19}

func main() {

	maxFigureOfArray(yourSlice)
	minFigureOfArray(yourSlice)
	sortArrayFromMin(yourSlice)
	sortArrayFromMax(yourSlice)
}

func maxFigureOfArray(array []int) {
	max := 0
	for _, value := range array {
		if value > max {
			max = value
		}
		fmt.Println("max:", max)
	}
}

func minFigureOfArray(array []int) {
	min := 0
	for _, value := range array {
		if value < min {
			min = value
		}
	}
	fmt.Println("min:", min)
}

func sortArrayFromMin(array []int) {

	for j := range array {
		fmt.Println("busting array:", j)
		for i := 0; i < len(array); i++ {
			for j = i + 1; j < len(array); j++ {
				if array[i] > array[j] {

					array[i], array[j] = array[j], array[i]
				}
			}
			fmt.Println(array[i])
		}
	}
}

func sortArrayFromMax(array []int){
	for j := range array {
		fmt.Println("busting array:", j)
		for i := 0; i < len(array); i++ {
			for j = i + 1; j < len(array); j++ {
				if array[i] < array[j] {

					array[i], array[j] = array[j], array[i]
				}
			}
			fmt.Println(array[i])
		}
	}
}
