package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Aaslema, baba!"
	count := strlen(str)
	fmt.Printf("String '%s' feha %d runes.", str, count)
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}
