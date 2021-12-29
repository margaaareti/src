// Преобразования между строками, рунами и байтами, и числами

package main

import "fmt"

//определяем размер среза
func main() {

	numbers := make([]int, 4)
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i
	}
	fmt.Println(numbers)

	//Многомерный срез

	twoSlices := [][]string{{"a", "b", "c", "d", "e", "f", "g", "h", "i"}, {"g", "k", "l", "m", "n", "o", "p"}}
	for i := 0; i <= len(twoSlices[0])-1; i++ {
		fmt.Println(twoSlices[0][i])
		if i == len(twoSlices[0])-1 {
			for j := 0; j <= len(twoSlices[1])-1; j++ {
				fmt.Println(twoSlices[1][j])
			}
		}

	}

	//Вырезаем второй элемент в 0 массиве
	twoSlices[0] = append(twoSlices[0][:1], twoSlices[0][2:]...)
	fmt.Println(twoSlices[0])
}
