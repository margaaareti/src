//страница 392

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("task2") // получает сегмент с элементами, представляющими содержимое "task2"
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files { // для каждого файла в сегменте.
		if file.IsDir() { // если файл является каталогом ...
			fmt.Println("Directory:", file.Name()) // ... то выводится имя каталога
		} else {
			fmt.Println("File:", file.Name()) // ... а если нет, то выводится имя файла
		}
	}
}
