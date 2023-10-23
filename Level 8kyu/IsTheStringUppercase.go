package main

import (
	"fmt"
	"unicode"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	hasLower := false
	for _, v := range s {
		if unicode.IsLower(v) {
			hasLower = true
			break
		}
	}
	return !hasLower
}

//func (s MyString) IsUpperCase() bool {
//	for _, v := range s {
//		if unicode.IsLower(v) {
//			return false
//		}
//	}
//	return true
//}

func main() {
	fmt.Println(MyString("c").IsUpperCase())                      // -> False
	fmt.Println(MyString("C").IsUpperCase())                      // -> True
	fmt.Println(MyString("hello I AM DONALD").IsUpperCase())      // -> False
	fmt.Println(MyString("HELLO I AM DONALD").IsUpperCase())      // -> True
	fmt.Println(MyString("ACSKLDFJSgSKLDFJSKLDFJ").IsUpperCase()) // -> False
	fmt.Println(MyString("ACSKLDFJSGSKLDFJSKLDFJ").IsUpperCase()) // -> True
}
