package main

import (
	"fmt"
	"strings"
)

func NoSpace(word string) string {
	s := strings.Replace(word, " ", "", -1)
	return s
}
func main() {
	fmt.Println(NoSpace("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd")) //"8j8mBliB8gimjB8B8jlB")
}
