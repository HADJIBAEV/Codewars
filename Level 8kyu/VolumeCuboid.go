package main

import "fmt"

func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}
func main() {
	fmt.Println(GetVolumeOfCuboid(1.0, 2.0, 2.0)) // 4
}
