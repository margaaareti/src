package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

func main() {
	myArrive := []int{1, 2, 10, 50, 5}
	complete := SortNumbers(myArrive)
	fmt.Println(complete)
	arriveEmpty := []int{}
	complete = SortNumbers(arriveEmpty)
	fmt.Println(complete)

}
