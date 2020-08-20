package main

import (
	"fmt"
)

func main() {
	// userInputs := []int{}
	fmt.Println("Enter array size")
	const arraySize = 12

	var myArray [arraySize]int
	fmt.Println("Enter 12 integers")
	for i := 0; i < arraySize; i++ {
		fmt.Println("Enter an integer")
		fmt.Scan(&myArray[i])
	}

	part1 := arraySize / 4
	middle := arraySize / 2
	part3 := part1 + middle
	c := make(chan []int, 4)
	go sort(myArray[:part1], c)
	go sort(myArray[part1:middle], c)
	go sort(myArray[middle:part3], c)
	go sort(myArray[part3:], c)
	arr1 := <-c
	arr2 := <-c
	arr3 := <-c
	arr4 := <-c

	result := merge(merge(arr1, arr2), merge(arr3, arr4))
	fmt.Println(result)
}

func sort(arr []int, c chan []int) {
	for j := 0; j < len(arr)-1; j++ {
		for i := j; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	fmt.Println(arr)
	c <- arr
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
