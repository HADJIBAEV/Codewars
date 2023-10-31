package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int) []int {
	if numbers == nil || len(numbers) == 0 {
		return []int{}
	}
	sort.Ints(numbers)
	return numbers
}

//	func SortNumbers(numbers []int) []int {
//		if numbers == nil || len(numbers) == 0 {
//			return []int{}
//		}
//		for i := 0; i < len(numbers)-1; i++ {
//			for j := 0; j < len(numbers)-i-1; j++ {
//				if numbers[j] > numbers[j+1] {
//					numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
//				}
//			}
//		}
//		return numbers
//	}
func main() {
	fmt.Println(SortNumbers([]int{1, 2, 3, 10, 5}))

}
