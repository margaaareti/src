package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
	"unicode/utf8"
)

func saveString(fileName string, str string) error {
	err := ioutil.WriteFile(fileName, []byte(str), 0600)
	return err
}

func main() {
	rand.Seed(time.Now().Unix())

	if err := saveString("hindi.txt", "Namaste"); err != nil {
		log.Fatal(err)
	}
	if err := saveString("english.txt", "Hello"); err != nil {
		log.Fatal(err)
	}
	/////////////////////////////////////////////////

	asciiString := "ABCDE"
	utf8String := "БГДЖИ"

	// выводит количество рун (символов в строке)
	fmt.Println(utf8.RuneCountInString(utf8String))

	// Строки преобразуются в сегменты рун
	asciiBytes := []rune(asciiString)
	utf8Rune := []rune(utf8String)
	//Из каждого сегмента исключаются первые три руны
	asciiBytesPartially := asciiBytes[3:]
	utf8RunesPartially := utf8Rune[3:]
	// Сегмент рун преобразуется в строку
	fmt.Println(string(asciiBytesPartially))
	fmt.Println(string(utf8RunesPartially))

	for index, currentRune := range asciiBytes { // Обрабатывает каждый байт в сегменте
		fmt.Printf("%d: %s\n", index, string(currentRune)) // преобразует байт встроку и выводит рузультат
	}
	for index, currentRune := range utf8Rune {
		fmt.Printf("%d: %s\n", index, string(currentRune))

	}

}
