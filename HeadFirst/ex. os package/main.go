package main

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*
		//Читаем данные из файла
		file, err := os.OpenFile("testfile.txt", os.O_RDONLY, os.FileMode(600)) // os.O_RDONLY- файл открывается для чтения
		check(err)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	*/

	//Записываем данные
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("testfile.txt", options, os.FileMode(0700)) // файл открывается для записи
	check(err)
	_, err = file.Write([]byte("amazing!\n")) // записывает данные в файл
	check(err)
	err = file.Close()
	check(err)

}
