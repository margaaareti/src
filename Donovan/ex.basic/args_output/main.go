package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	s = ""
	sep = ""
	for _, word := range os.Args[1:] {
		s += sep + word
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))

}
