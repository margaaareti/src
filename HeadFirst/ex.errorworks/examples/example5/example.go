// вызов recover ghb панике
package main

import "fmt"

func calmDown() {
	recover()              //разместить вызов recover в другой функции
	fmt.Println(recover()) // выводит значение, переданное panic
}

func freakOut() {
	defer calmDown() // отложить вызов функции recover
	panic("oh,no")   // Если в программе после этого возникнет паника, отложенный вызов функции позволит провести восстановление (и вернет значение panic)

}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
