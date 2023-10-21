package main

import "fmt"

func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
}
func main() {
	fmt.Println(SetAlarm(true, true))
	fmt.Println(SetAlarm(true, false))
	fmt.Println(SetAlarm(false, true))
	fmt.Println(SetAlarm(false, false))
}
