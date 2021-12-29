package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	x := 123
	// преобразует целое число в строку с символом преобразования
	y := fmt.Sprintf("%d", x)
	// strcon.Itoa- выполняет аналогичную функцию
	fmt.Println(y, strconv.Itoa(x))

	//преобразует строку в число
	z, err := strconv.Atoi("123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d,%[1]T ", z)

	//в десятичной системе счисления до 64 битов
	g, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d,%[1]T\n", g)

}
