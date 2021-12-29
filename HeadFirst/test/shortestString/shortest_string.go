// Возвращает самое короткое слово в строке
package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	var t int
	for i, s := range strings.Fields(s) { // strings.Fields разбивает строку на слова
		fmt.Println(i, s)
		if len(s) < t || i == 0 {
			t = len(s)
		}
	}
	return t
}

func main() {
	myString := " assss ASSSSSSSSSSS fdfsdfsdfsdfsd"
	complete := FindShort(myString)
	fmt.Println(complete)
}
