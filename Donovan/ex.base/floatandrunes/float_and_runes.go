package main

import (
	"fmt"
	"math"
)

func main() {
	medals := []string{"золото", "серебро", "бронза"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	// Глаголы форматирования чисел int в Printf
	myNumber := 0555
	fmt.Printf("%o %[1]d %[1]b %[1]x\n", myNumber) // %o - восьмеричная СС, %d- десятичная, %b - двоичная, %x - шестнадцатиречная

	// Глаголы форматирования символов
	ascii := 'a'
	unicode := 2719
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // %d - код символа в unicode, %c - сам символ, %q - символ в кавычках
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)

	//Глагол форматирования числа с плавающей точкой
	// Экспонента является показательной функцией ƒ(x) = ex, где число e (число Эйлера, e ≈ 2,718) – основание степени, а х – степень в которую нужно его возвести.
	for x := 1; x < 8; x++ {
		fmt.Printf("x = %d  е= %8.3f \n", x, math.Exp(float64(x))) // приводим тип x к float64
	}

	s := "Hello,世 界"
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Printf("%x", string(s))

}
