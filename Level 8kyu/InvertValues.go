package main

import "fmt"

func Invert(arr []int) []int {
	Inverted := make([]int, len(arr))
	for i, v := range arr {
		Inverted[i] = -v
	}

	return Inverted
}

func main() {
	fmt.Println(Invert([]int{1, 2, 3, 4, 5}))   //  [-1, -2, -3, -4, -5]
	fmt.Println(Invert([]int{1, -2, 3, -4, 5})) //  [-1, 2, -3, 4, -5]
}
