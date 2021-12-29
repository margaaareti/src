// стр.384
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) { // открывает файл и возвращает указатель на него, а также значение ошибки
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName) // Вместо прямого вызова os.Open вызывается функция OpenFile
	if err != nil {
		return nil, err
	}
	defer CloseFile(file) // Вместо прямого вызова file.Close вызывается CloseFil. Даже если будет получена ошибка, функция CloseFile будет вызвана
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {

	numbers, err := GetFloats(os.Args[1]) //Сохраняет сегмент чисел, прочитанных из файла
	if err != nil {
		log.Fatal(err)
	}

	var sum float64
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Sum: %0.2f \n", sum)

}
