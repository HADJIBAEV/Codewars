package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	initials := ""
	ConvertRune := []rune(name)
	for i, v := range ConvertRune {
		if string(v) == " " {
			initials = string(ConvertRune[0]) + "." + string(ConvertRune[i+1])
		}
	}
	return strings.ToUpper(initials)
}

func main() {
	fmt.Println(AbbrevName("Sam harris"))  //S.H
	fmt.Println(AbbrevName("Иван Петров")) // И.П
}
