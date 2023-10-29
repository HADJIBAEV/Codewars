package main

import (
	"fmt"
	"strconv"
)

func PrinterError(s string) string {
	errors := 0
	total := len(s)
	for _, v := range s {
		if v < 'a' || v > 'm' {
			errors++
		}
	}

	return (strconv.Itoa(errors)) + "/" + (strconv.Itoa(total))
}

func main() {
	fmt.Println(PrinterError("aaabbbbhaijjjm"))         // "0/14"
	fmt.Println(PrinterError("aaaxbbbbyyhwawiwjjjwwm")) // "8/22"
}
