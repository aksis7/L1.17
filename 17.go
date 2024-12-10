package main

import "fmt"

// Функция бинарного поиска
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	// Пока есть элементы для поиска
	for low <= high {
		mid := low + (high-low)/2 // Находим середину массива

		// Если элемент найден
		if arr[mid] == target {
			return mid
		}

		// Если целевой элемент меньше среднего
		if arr[mid] > target {
			high = mid - 1
		} else { // Если целевой элемент больше среднего
			low = mid + 1
		}
	}

	// Если элемент не найден
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Элемент найден на индексе: %d\n", result)
	} else {
		fmt.Println("Элемент не найден")
	}
}
