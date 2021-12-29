package main

import (
	"HeadFirst/ex.testing/prose"
	"fmt"
)

func main() {

	phrases := []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"apple", "orange", "pear"}
	fmt.Println(prose.JoinWithCommas(phrases))

}
