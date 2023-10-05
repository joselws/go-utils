package sorts

import "github.com/joselws/go-utils/mytypes"

// Sort a slice with insertion sort.
func InsertionSort[T mytypes.Number](elements []T) {
	for i := range elements {
		for j := i; j > 0 && elements[j] < elements[j-1]; j-- {
			elements[j], elements[j-1] = elements[j-1], elements[j]
		}
	}
}

// Sort a slice with selection sort.
func SelectionSort[T mytypes.Number](elements []T) {
	for i := range elements {
		var minIndex int = i
		for j := i; j != len(elements); j++ {
			if elements[j] < elements[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			elements[i], elements[minIndex] = elements[minIndex], elements[i]
		}
	}
}