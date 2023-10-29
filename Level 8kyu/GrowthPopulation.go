package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0
	population := p0
	for population < p {
		increase := population + int(float64(population)*(percent/100))
		population = increase + aug
		years++
	}
	return years
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000)) // 15
	fmt.Println(NbYear(1000, 2, 50, 1200))  // 3
}
