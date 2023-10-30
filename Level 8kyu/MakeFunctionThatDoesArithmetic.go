package main

import (
	"fmt"
	"strings"
)

func Arithmetic(a int, b int, operator string) int {
	switch {
	case strings.ToLower(operator) == "add":
		return a + b
	case strings.ToLower(operator) == "subtract":
		return a - b
	case strings.ToLower(operator) == "multiply":
		return a * b
	case b == 0:
		return 0
	case strings.ToLower(operator) == "divide":
		return a / b
	default:
		return 0
	}

}
func main() {
	fmt.Println(Arithmetic(5, 5, "divide"))
}

