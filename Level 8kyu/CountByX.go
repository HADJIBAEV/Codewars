package main

import "fmt"

func CountBy(x, n int) []int {
	var slice []int
	for i := 1; i <= n; i++ {
		slice = append(slice, i*x)
	}
	return slice
}

//	func CountBy(x, n int) []int {
//		slice:=make([]int, n)
//		for i := 0; i < n; i++ {
//			slice[i] = (i+1)*x
//		}
//		return slice
//	}
func main() {
	fmt.Println(CountBy(1, 10)) //should return  {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(CountBy(2, 5))  //should return {2,4,6,8,10}
}
