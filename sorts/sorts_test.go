package sorts

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	numbers := []int{3, 10, 8, 4, 9, 6, 1, 2, 5, 7}
	SelectionSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestSelectionSortEmpty(t *testing.T) {
	numbers := []int{}
	SelectionSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestSelectionSortSingleNumber(t *testing.T) {
	numbers := []int{1}
	SelectionSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestSelectionSortAlreadySorted(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	SelectionSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	numbers := []int{3, 10, 8, 4, 9, 6, 1, 2, 5, 7}
	InsertionSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestInsertionSortEmpty(t *testing.T) {
	numbers := []int{}
	InsertionSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestInsertionSortSingleElement(t *testing.T) {
	numbers := []int{1}
	InsertionSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestInsertionSortAlreadySorted(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	InsertionSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	numbers := []int{3, 10, 8, 4, 9, 6, 1, 2, 7, 5}
	BubbleSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestBubbleSortEmpty(t *testing.T) {
	numbers := []int{}
	BubbleSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestBubbleSortSingleElement(t *testing.T) {
	numbers := []int{1}
	BubbleSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	BubbleSort(numbers)
	for index, number := range numbers {
		if index+1 != number {
			t.Error("Number should be", index+1, "not", number)
		}
	}
}
