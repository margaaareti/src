package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//GetFloat читает каждое значение float64 из каждой строки файла

func GetFloats(fileName string) ([]float64, error) { //Имя файла с данными передается в аргументе. Функция возвращает массив чисел и ошибку
	var numbers []float64 // Объявление возвращаемого массива

	file, err := os.Open(fileName) // открывает файл с переданным именем
	if err != nil {
		return nil, err
	}

	//i := 0 //переменная для хранения индекса, по которому должно выполнятся присваивание

	scanner := bufio.NewScanner(file) // значение os.File передается функции bufio.NewScanner. Функция возвращает значение bufio.Scanner для чтения данных из файла
	for scanner.Scan() {              // читает строку из файла
		number, err := strconv.ParseFloat(scanner.Text(), 64) //строка, прочитанная из файла, преобразуется в float64
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number) //Новое число присоединяется к сегменту
		//i++ // переход к следующему индексу массива
	}

	err = file.Close() // Возвращает ошибку, если при закрытии файла таковая обнаружена
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil { // Возвращает ошибку, если при сканировании файла таковая обнаружена
		return nil, scanner.Err()
	}
	return numbers, nil
}
