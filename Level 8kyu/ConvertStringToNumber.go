package main

import (
	"fmt"
	"log"
	"strconv"
)

func StringToNumber(str string) int {
	convert, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return convert
}
func main() {
	fmt.Println(StringToNumber("1234")) // --> 1234
}
