package main

import "fmt"

func Opposite(value int) int {
	if value > 0 {
		return value * (-1)
	} else {
		return value * (-1)
	}
}
func main() {
	fmt.Println(Opposite(1)) // -1
}
