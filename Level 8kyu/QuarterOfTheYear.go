package main

import "fmt"

func QuarterOf(month int) int {
	var quarter int
	if month >= 1 && month <= 12 {
		switch {
		case month >= 1 && month <= 3:
			quarter = 1
		case month >= 4 && month <= 6:
			quarter = 2
		case month >= 7 && month <= 9:
			quarter = 3
		case month >= 10 && month <= 12:
			quarter = 4
		}
	} else {
		return 0
	}
	return quarter
}
func main() {
	fmt.Println(QuarterOf(3)) // Quarter 1
}
