// пакет datafile предназначен для чтение данных из файлов
package datafile

import (
	"bufio"
	"os"
)

//GetStrings читает каждое значение float64 из каждой строки файла

func GetString(fileName string) ([]string, error) {
	var lines []string             //в переменной хранится массив строк
	file, err := os.Open(fileName) // отрывает файл с переданным именем
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file) // Возвращает значения данных из файла
	for scanner.Scan() {              //читает , последовательно, каждую строку из файла
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close() // возвращает ошибку при закрытии файла, если она есть
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
