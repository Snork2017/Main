package main

import (
	"fmt"
	"errors"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	binarySearch(numbers, 7)
}

func binarySearch(array []int, target int) (int, error) {
	var startPos = 0
	var endPos = len(array)

	for startPos != endPos {
		var midPos = (startPos + endPos) / 2
		var midValue = array[midPos]

		if target == midValue {
			fmt.Println(target)
			break

		} else if target < midValue {
			endPos = midPos + 1
			fmt.Println(target)
			break
		} else {
			startPos = midPos - 1
			fmt.Println(target)
			break
		}
	}
	return -1, errors.New("no target")
}
