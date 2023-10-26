package main

import "fmt"

func combat(health, damage float64) float64 {
	newHealth := health - damage
	if newHealth < 0 {
		return 0
	}
	return newHealth
}

func main() {
	fmt.Println(combat(100.0, 50.0)) // 50.0
}
