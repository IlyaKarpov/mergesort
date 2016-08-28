package main

import (
	"fmt"
)

func Sort(array []int) {
	stopLength := len(array)
	split(array, stopLength)
}

func split(array []int, stopLength int) {
	length := len(array)
	if length <= 1 {
		return
	} else {
		middle := length / 2
		split(array[:middle], stopLength)
		split(array[middle:], stopLength)
		merge(array[:middle], array[middle:])
		if stopLength == len(array[:middle])+len(array[middle:]) {
			copy(array, array[:middle])
			array = append(array, array[middle:]...)
		}
	}
}

func merge(first []int, second []int) {
	firstLen := len(first)
	secondLen := len(second)

	newFirst := make([]int, firstLen)
	newSecond := make([]int, secondLen)

	fIndex := 0
	sIndex := 0
	newElement := 0
	countOfConcederedElements := 0

	for countOfConcederedElements != len(first)+len(second) {
		if fIndex >= firstLen {
			newElement = second[sIndex]
			sIndex++
		} else if sIndex >= secondLen {
			newElement = first[fIndex]
			fIndex++
		} else if first[fIndex] > second[sIndex] {
			newElement = second[sIndex]
			sIndex++
		} else {
			newElement = first[fIndex]
			fIndex++
		}
		countOfConcederedElements++

		if countOfConcederedElements <= firstLen {
			newFirst[countOfConcederedElements-1] = newElement
		} else {
			newSecond[countOfConcederedElements-firstLen-1] = newElement
		}
	}

	copy(first, newFirst)
	copy(second, newSecond)
}

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	Sort(array)
	fmt.Println(array)
}
