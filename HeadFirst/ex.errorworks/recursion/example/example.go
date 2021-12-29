// стр. 394

package main

import "fmt"

func count(start int, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	fmt.Println(start) //вывести текущее начальное число
	if start < end {   // если конечное число еще не достигнуто...
		count(start+1, end) // ... то функция count вызывает сама себя с начальным числом на 1 больше текущего
	}
	fmt.Printf("Returning from count(%d, %d) call\n", start, end)
}

func main() {

	count(1, 3)

}
