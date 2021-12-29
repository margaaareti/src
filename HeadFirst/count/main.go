//подсчитывает вхождение каждой строки в файле

package main

import (
	"HeadFirst/datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetString("votes.txt") //Читает указанный файл и возвращает сегмент, содержащий все строки из файла
	if err != nil {
		log.Fatal(err)
	}
	var names []string // переменная для хранения сегмента с именами кандидатов
	var counts []int   // сегмент с кол. вхождений каждого имени

	for _, line := range lines {
		matched := false
		for i, name := range names { //перебор всез значений из сегмента names
			if name == line { // если эта строка совпадает с текущим именем
				counts[i]++    //увеличивает соответствующий счетчик
				matched = true // устанавливает признак обнаруженного совпадения
			}
		}
		if matched == false { // если совпадение не найдено
			names = append(names, line) // добавить его, как новое имя в сегмент
			counts = append(counts, 1)  // добавить счетчик
		}

	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}

}
